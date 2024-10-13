package ratelimiter

import (
    "net/http"
    "time"
)

type RateLimiter struct {
    requests map[string]int
    limit    int
}

func NewRateLimiter(limit int) *RateLimiter {
    return &RateLimiter{requests: make(map[string]int), limit: limit}
}

func (rl *RateLimiter) CheckLimit(req *http.Request) bool {
    ip := req.RemoteAddr
    rl.requests[ip]++
    
    if rl.requests[ip] > rl.limit {
        return true
    }

    // Reset setelah 1 menit
    go func() {
        time.Sleep(1 * time.Minute)
        rl.requests[ip] = 0
    }()
    return false
}
