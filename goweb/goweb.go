package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func createToken() (token string) {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token = fmt.Sprintf("%x", h.Sum(nil))
	return token
}

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

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // get the method of request
	if r.Method == "GET" {
		token := createToken()

		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, token)
	} else {
		// "POST", then parse the from data
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			// check token is valid
		} else {
			// token doesn't exist
		}
		fmt.Println("username length", len(r.Form["username"][0]))
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		token := createToken()

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20) // max memory to save uploadfile
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close() // file must be losed
		fmt.Fprintf(w, "%v", handler.Header)
		// /test is the dir of uploadfile
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/", sayHelloName) // set access url and get a handler
	http.HandleFunc("/login", login)   // set the url of login
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil) // set listening port
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

/*
test the web-server

http://localhost:9090
http://localhost:9090/?url_long=111&url_long=222

#login
http://127.0.0.1:9090/login

*/
