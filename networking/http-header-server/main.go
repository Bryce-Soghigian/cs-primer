package main

import (
	"log"
	"io"
	"net/http"
)

/*
Requirements:

The program will be a simple http server that serves back the headers of the request in JSON format.

*/


func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	


	http.HandleFunc("/headers", headersHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}


func headersHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Headers:\n")
	for name, headers := range req.Header {
		for _, h := range headers {
			io.WriteString(w, name)
			io.WriteString(w, ": ")
			io.WriteString(w, h)
			io.WriteString(w, "\n")
		}
	}
}
