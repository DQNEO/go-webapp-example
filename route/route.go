package route

import "net/http"
import "../handler"

func Get(path string, h func(http.ResponseWriter, *http.Request)) {
	http.DefaultServeMux.HandleFunc(path, h)
}

func Register() {
	Get("/hello", handler.GetHello)
	Get("/hello.html", handler.GetHelloHTML)
	Get("/hello.json", handler.GetHelloJson)
	Get("/issues", handler.GetIssues)
	Get("/issues/1", handler.GetIssue1)
}
