package main

import (
    "fmt"
    "net/http"
    "pwnwaf/core"
    "pwnwaf/config"
)

func main() {
    cfg := config.LoadConfig() // Muat konfigurasi dari file
    waf := core.NewWAF(cfg)    // Buat instance WAF

    // Jalankan server
    http.HandleFunc("/", waf.HandleRequest)
    fmt.Println("PwnWAF running on port", cfg.Server.Port)
    http.ListenAndServe(fmt.Sprintf(":%d", cfg.Server.Port), nil)
}