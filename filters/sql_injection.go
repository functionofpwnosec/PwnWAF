package filters

import (
    "net/http"
    "strings"
)

func CheckSQLInjection(req *http.Request) bool {
    // Cek apakah ada tanda-tanda SQL Injection di URL atau parameter
    sqlKeywords := []string{"SELECT", "INSERT", "UPDATE", "DELETE", "--", "'"}

    for _, keyword := range sqlKeywords {
        if strings.Contains(strings.ToUpper(req.URL.RawQuery), keyword) {
            return true
        }
    }
    return false
}
