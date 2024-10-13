package filters

import (
    "net/http"
    "strings"
)

func CheckXSS(req *http.Request) bool {
    // Cek apakah ada tanda-tanda XSS di URL atau parameter
    xssKeywords := []string{"<script>", "javascript:", "onerror", "onload"}

    for _, keyword := range xssKeywords {
        if strings.Contains(strings.ToLower(req.URL.RawQuery), keyword) {
            return true
        }
    }
    return false
}
