package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // You have to parse these yourself
	fmt.Println(r.Form) // print form information in server side
	fmt.Println(r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key", k)
		fmt.Println("value", v)
	}
	fmt.Fprint(w, "Hello saurabh-sikchi")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) // get request method
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("./src/github.com/saurabh-sikchi/simple_webserver/login.html")
		t.Execute(w, token)
	} else {
		r.ParseForm()

		token := r.Form.Get("token")
		if token != "" {
			// log in request
		} else {
			// give error if no token
		}
		// logic part of log in
		fmt.Println("username length:", len(r.Form["username"][0]))
		fmt.Println("username: ", template.HTMLEscapeString(r.FormValue("username")))
		fmt.Println("password: ", template.HTMLEscapeString(r.FormValue("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
	}
}

func main() {
	http.HandleFunc("/", sayHelloName) // set router
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
