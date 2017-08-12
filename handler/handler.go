package handler

import "net/http"
import "fmt"
import "html"
import "../model"
import "encoding/json"

func GetIssue1(w http.ResponseWriter, r *http.Request) {
	id := 1
	issue, err := model.FindIssue(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(404)
		return
	}

	b,err := json.Marshal(issue)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(b)
}

func GetIssues(w http.ResponseWriter, r *http.Request) {
	issues := model.GetIssues()
	b, err := json.Marshal(issues)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}

	w.Write(b)
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

