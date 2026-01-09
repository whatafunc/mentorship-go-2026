package main

import (
	"sync"
	"time"
)

type Bucket struct {
	Tokens float64
	Last   time.Time
}

var (
	mu      sync.Mutex
	buckets = map[string]*Bucket{}
)

func CheckIp(ip string) bool {
	mu.Lock()
	defer mu.Unlock()

	b, ok := buckets[ip]
	if !ok {
		b = &Bucket{Tokens: 5, Last: time.Now()}
		buckets[ip] = b //
	}

	return Allow(b, 0.1, 5) // 1 token per 10s, max 5
	//10 попыток в минуту для данного логина
}

func Allow(b *Bucket, rate float64, capacity float64) bool {
	now := time.Now()
	elapsed := now.Sub(b.Last).Seconds()
	b.Tokens += elapsed * rate
	if b.Tokens > capacity {
		b.Tokens = capacity
	}
	b.Last = now
	return b.Tokens >= 1
}
