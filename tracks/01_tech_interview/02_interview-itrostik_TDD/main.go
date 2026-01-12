package main

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
	"time"
)

type Source interface {
	Get(key string) (string, error)
}

type Request struct {
	Key string `json:"key"`
}

type RequestData struct {
	data      string
	createdAt int64
}

func clearCache(cache map[string]RequestData, timeClearCache int) {
	go func() {
		for {
			time.Sleep(time.Duration(timeClearCache) * time.Millisecond)
			now := time.Now().Unix()
			for key, data := range cache {
				if now-data.createdAt > 5000 {
					delete(cache, key)
				}
			}
		}
	}()
}

func newHandler(api Source) http.HandlerFunc {
	cache := make(map[string]RequestData)
	timeClearCache := 5000 // milliseconds
	clearCache(cache, timeClearCache)
	var mu sync.RWMutex
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		req, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var request Request
		json.Unmarshal(req, &request)

		mu.RLock()
		if cached, ok := cache[request.Key]; ok { // fix the map lookup
			mu.RUnlock()
			w.Write([]byte(cached.data))
			return
		}
		mu.RUnlock()

		res, err := api.Get(request.Key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		mu.Lock()
		defer mu.Unlock()
		cache[request.Key] = RequestData{
			data:      res,
			createdAt: time.Now().Unix(),
		}

		w.Write([]byte(res))
	})
}
