package main

import "net/http"
import "log"
import "./handler"
import "./router"
import (
	"fmt"
)

const defaultPort = 9001


func Register(rt *router.Router) *router.Router {

	rt.Get("/hello", handler.GetHello)
	rt.Post("/hello", handler.PostHello)
	rt.Put("/hello", handler.PutHello)
	rt.Delete("/hello", handler.DeleteHello)

	rt.Get("/issues", handler.GetAllIssues)
	rt.Get("/issues/search", handler.SearchIssues)
	rt.Get("/issues/{id}", handler.GetIssue)
	rt.Post("/issues", handler.CreateIssue)
	rt.Put("/issues/{id}", handler.UpdateIssue)
	rt.Delete("/issues/{id}", handler.DeleteIssue)
	return rt
}

func main() {
	rt := router.NewRouter()
	rt = Register(rt)
	handler := http.Handler(rt)
	port := defaultPort
	log.Printf("starting server : http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",port), handler))

}
