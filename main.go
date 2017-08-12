package main
import "./handler"
import "net/http"
import "log"

func Register() {
	http.HandleFunc("/hello", handler.GetHello)
	http.HandleFunc("/hello.html", handler.GetHelloHTML)
	http.HandleFunc("/hello.json", handler.GetHelloJson)
	http.HandleFunc("/issues", handler.GetIssues)
	http.HandleFunc("/issues/1", handler.GetIssue1)
}

func main() {

	Register()

	log.Fatal(http.ListenAndServe(":8080", nil))

}
