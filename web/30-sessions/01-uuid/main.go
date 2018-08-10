package main

import (
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"
)

// For this code to run, you will need this package:
// go get github.com/satori/go.uuid

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8088", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		id, err := uuid.NewV4()
		if err != nil {
			fmt.Println(err)
		}
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}
