package main

import (
	"io"
	"log"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

func main() {
	http.HandleFunc("/hello", hellohandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("listenandserve: ", err.Error())
	}
}
