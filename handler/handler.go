package handler

import "net/http"
import "fmt"
import "html"
import "../model"

func GetIssue1(w http.ResponseWriter, r *http.Request) {
	id := 1
	issue, err := model.FindIssue(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		w.Write([]byte("Record Not Found"))
		w.WriteHeader(404)
		return
	}
	fmt.Fprintf(w, `%v`, issue)
}

func GetIssues(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `%v`, model.GetIssues)
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

