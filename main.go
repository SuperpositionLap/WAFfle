package main

import (
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

// WAF middleware
func WAF() gin.HandlerFunc {
    return func(c *gin.Context) {
        // SQL Injection detection
        query := c.Request.URL.RawQuery
        if strings.Contains(query, "SELECT") || strings.Contains(query, "UNION") || strings.Contains(query, "INSERT") || strings.Contains(query, "DELETE") {
            c.String(http.StatusForbidden, "SQL Injection detected")
            c.Abort()
            return
        }

        // XSS detection
        if strings.Contains(query, "<script>") || strings.Contains(query, "javascript:") {
            c.String(http.StatusForbidden, "XSS detected")
            c.Abort()
            return
        }

        c.Next()
    }
}

func main() {
    r := gin.Default()

    // Apply WAF middleware
    r.Use(WAF())

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
}