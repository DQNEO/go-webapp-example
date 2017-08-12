package handler

import "net/http"
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

func GetIssue1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `%v`, issues[0])
}

func GetIssues(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `%v`, issues)
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "path is %q", html.EscapeString(r.URL.Path))
}

func GetHelloHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func GetHelloJson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"msg":"hello"}`)
}

