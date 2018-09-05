package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dinhoabreu/go-training/web"
)

type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/mshl", mshl)
	http.HandleFunc("/encd", encd)
	util.ServeHTTP()
}

func root(w http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
		<html lang="en">
		<head>
		<meta charset="UTF-8">
		<title>FOO</title>
		</head>
		<body>
			<a href="/mshl">Marshal</a>
			<a href="/encd">Encode</a>
		</body>
		</html>`
	w.Write([]byte(s))
}

func mshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	j, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(j)
}

func encd(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}
