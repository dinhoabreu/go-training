package util

import (
	"fmt"
	"net/http"
	"os"
)

// GetEnv environment variables with default value
func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}

// ServeHTTP http.ListenAndServe using port from PORT environment
func ServeHTTP() {
	port := GetEnv("PORT", "8080")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
