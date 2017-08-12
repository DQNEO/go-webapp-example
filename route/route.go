package route

import "net/http"
import "../handler"

func GET(path string , h func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(path, h)
}

func Register() {
	GET("/hello", handler.GetHello)
	GET("/hello.html", handler.GetHelloHTML)
	GET("/hello.json", handler.GetHelloJson)
	GET("/issues", handler.GetIssues)
	GET("/issues/1", handler.GetIssue1)
}

