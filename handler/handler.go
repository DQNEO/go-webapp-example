package handler

import "net/http"
import "fmt"
import "html"
import "../model"
import "../response"

func fail(w http.ResponseWriter, err error) {
	w.WriteHeader(500) // or can be 404
	w.Write([]byte(err.Error()))
}

func GetIssue1(w http.ResponseWriter, r *http.Request) {
	id := 1
	issue, err := model.FindIssue(id)
	if err != nil {
		fail(w, err)
		return
	}

	response.SendJson(w, issue)
}

func GetIssues(w http.ResponseWriter, r *http.Request) {
	issues, err := model.GetIssues()
	if err != nil {
		fail(w, err)
		return
	}
	response.SendJson(w, issues)
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
