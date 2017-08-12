package main

import "net/http"
import "log"
import "fmt"
import "html"

type Issue struct {
	Id   int
	Name string
}

var issues = []Issue{
	{Id:1, Name:"I need a help"},
	{Id:2, Name:"I need another help"},
}

func getIssue(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `%v`, issues[0])
}

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


	http.HandleFunc("/issues", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `%v`, issues)
	})

	http.HandleFunc("/issues/1", getIssue)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
