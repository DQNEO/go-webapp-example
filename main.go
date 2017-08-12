package main

import "net/http"
import "log"
import "fmt"
import "html"

func main() {

	//http.Handle("/foo", fooHandler)

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "path is %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hello.html", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})

	http.HandleFunc("/hello.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"msg":"hello"}`)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
