package core

import (
    "fmt"
    "net/http"
    "pwnwaf/filters"
)

type WAF struct {
    config *Config
}

func NewWAF(cfg *Config) *WAF {
    return &WAF{config: cfg}
}

func (w *WAF) HandleRequest(wr http.ResponseWriter, req *http.Request) {
    // Cek apakah request mencurigakan
    if filters.CheckSQLInjection(req) || filters.CheckXSS(req) {
        http.Error(wr, "Request blocked by PwnWAF", http.StatusForbidden)
        return
    }

    fmt.Fprintln(wr, "Request allowed by PwnWAF")
}
