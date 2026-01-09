package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockAPI struct {
	data string
}

func (m *mockAPI) Get(key string) (string, error) {
	return m.data, nil
}

func TestHandler_CacheHit(t *testing.T) {
	api := &mockAPI{data: "testData"}
	handler := newHandler(api)

	reqBody := `{"key": "someKey"}`
	req, _ := http.NewRequest(http.MethodPost, "/smth", strings.NewReader(reqBody))
	w := httptest.NewRecorder()
	handler(w, req)

	// Second request should hit the cache
	req2, _ := http.NewRequest(http.MethodPost, "/smth", strings.NewReader(reqBody))
	w2 := httptest.NewRecorder()
	handler(w2, req2)

	if w2.Body.String() != "testData" {
		t.Errorf("expected cached data 'testData', got '%s'", w2.Body.String())
	}
}

func TestHandler_MethodNotAllowed(t *testing.T) {
	api := &mockAPI{data: "testData"}
	handler := newHandler(api)

	req, _ := http.NewRequest(http.MethodGet, "/smth", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

func TestHandler_InvalidBody(t *testing.T) {
	api := &mockAPI{data: "testData"}
	handler := newHandler(api)

	req, _ := http.NewRequest(http.MethodPost, "/smth", strings.NewReader("{invalid json"))
	w := httptest.NewRecorder()
	handler(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_RaceCondition(t *testing.T) {
	api := &mockAPI{data: "raceData"}
	handler := newHandler(api)

	n := 100
	done := make(chan struct{}, n)
	for i := 0; i < n; i++ {
		go func(i int) {
			req, _ := http.NewRequest(http.MethodPost, "/smth", strings.NewReader(`{"key": "raceKey"}`))
			w := httptest.NewRecorder()
			handler(w, req)
			if w.Code != http.StatusOK {
				t.Errorf("goroutine %d: expected status code %d, got %d", i, http.StatusOK, w.Code)
			}
			done <- struct{}{}
		}(i)
	}
	for i := 0; i < n; i++ {
		<-done
	}
}
