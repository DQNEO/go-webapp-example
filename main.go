package main

import "net/http"
import "log"
import "./handler"
import (
	"regexp"
	"strings"
	"fmt"
)

const defaultPort = 9001

type Router struct {
	SimpleMaps map[string]map[string]http.HandlerFunc
	RegexMaps  map[string]map[*regexp.Regexp]http.HandlerFunc
}

func (rt *Router) RegisterHandler(method string, pattern string, h func(http.ResponseWriter, *http.Request)) {
	if strings.Contains(pattern, "{id}") {
		newPattern := strings.Replace(pattern, "{id}", "([a-zA-Z0-9]+)", -1)
		reg := regexp.MustCompile(newPattern)
		rt.RegexMaps[method][reg] = http.HandlerFunc(h)
	} else {
		rt.SimpleMaps[method][pattern] = http.HandlerFunc(h)
	}
}

func (rt *Router) Get(pattern string, h func(http.ResponseWriter, *http.Request)) {
	rt.RegisterHandler("GET", pattern, h)
}

func (rt *Router) Post(pattern string, h func(http.ResponseWriter, *http.Request)) {
	rt.RegisterHandler("POST", pattern, h)
}

func (rt *Router) Put(pattern string, h func(http.ResponseWriter, *http.Request)) {
	rt.RegisterHandler("PUT", pattern, h)
}

func (rt *Router) Delete(pattern string, h func(http.ResponseWriter, *http.Request)) {
	rt.RegisterHandler("DELETE", pattern, h)
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// test complete match
	if k, ok := rt.SimpleMaps[r.Method][r.URL.Path]; ok {
		k(w, r)
		return
	}

	for reg, h := range rt.RegexMaps[r.Method] {
		if matches := reg.FindAllStringSubmatch(r.URL.Path, -1); matches != nil {
			handler.URLParam = matches
			h(w, r)
			return
		}
	}
	w.WriteHeader(404)
	w.Write([]byte("Not Found"))
	return

}

func NewRouter() *Router {
	rt := &Router{}

	rt.SimpleMaps = make(map[string]map[string]http.HandlerFunc)
	rt.RegexMaps = make(map[string]map[*regexp.Regexp]http.HandlerFunc)
	rt.SimpleMaps["GET"] = make(map[string]http.HandlerFunc)
	rt.RegexMaps["GET"] = make(map[*regexp.Regexp]http.HandlerFunc)
	rt.SimpleMaps["POST"] = make(map[string]http.HandlerFunc)
	rt.RegexMaps["POST"] = make(map[*regexp.Regexp]http.HandlerFunc)
	rt.SimpleMaps["PUT"] = make(map[string]http.HandlerFunc)
	rt.RegexMaps["PUT"] = make(map[*regexp.Regexp]http.HandlerFunc)
	rt.SimpleMaps["DELETE"] = make(map[string]http.HandlerFunc)
	rt.RegexMaps["DELETE"] = make(map[*regexp.Regexp]http.HandlerFunc)
	return rt
}

func Register() *Router {
	rt := NewRouter()

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

	rt := Register()
	handler := http.Handler(rt)
	port := defaultPort
	log.Printf("starting server : http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",port), handler))

}
