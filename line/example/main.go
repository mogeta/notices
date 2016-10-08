package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, World")
}

func delayNotice(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	//fmt.Fprintf(w, "Hello, World")
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/delay",delayNotice)
	http.Handle("/", http.FileServer(http.Dir("res")))
	http.ListenAndServe(":8080", nil)
}