package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	// URL: scheme://host[:port#]/path/.../[?query-string][#anchor]
	r.ParseForm()       // parse args
	fmt.Println(r.Form) // push infto to webserver
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello world") // push info to web-client
}

func main() {
	http.HandleFunc("/", sayHelloName)       // set access url and get a handler
	err := http.ListenAndServe(":9090", nil) // set listening port
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

/*
test the web-server

http://localhost:9090
http://localhost:9090/?url_long=111&url_long=222

*/
