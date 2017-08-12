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

type e string

func (s e) Error() string {
	return string(s)
}

func FindIssue(id int) (Issue, error) {
	for id, issue := range(issues) {
		if issue.Id == id {
			return issue,nil
		}
	}
	var e e
	e = "error"
	return Issue{},e
}

func GetIssue1(w http.ResponseWriter, r *http.Request) {
	id := 1
	issue, err := FindIssue(id)
	if err != nil {
		w.Write([]byte("Record Not Found"))
		w.WriteHeader(404)
		return
	}
	fmt.Fprintf(w, `%v`, issue)
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

