package main

import "net/http"
import "log"
import "./handler"
import (
	"regexp"
	"strings"
)

type Mux struct {
	SimpleMaps map[string] map[string] http.HandlerFunc
	RegexMaps  map[string] map[*regexp.Regexp] http.HandlerFunc
}

func(mux *Mux) RegisterHandler(method string ,pattern string, h func(http.ResponseWriter, *http.Request)) {
	if strings.Contains(pattern, "{id}") {
		newPattern := strings.Replace(pattern,"{id}", "([a-zA-Z0-9]+)", -1)
		reg := regexp.MustCompile(newPattern)
		mux.RegexMaps[method][reg] = http.HandlerFunc(h)
	} else {
		mux.SimpleMaps[method][pattern] = http.HandlerFunc(h)
	}
}

func(mux *Mux) Get(pattern string, h func(http.ResponseWriter, *http.Request)) {
	mux.RegisterHandler("GET", pattern, h)
}

func(mux *Mux) Post(pattern string, h func(http.ResponseWriter, *http.Request)) {
	mux.RegisterHandler("POST", pattern, h)
}

func(mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// test complete match
	if k,ok := mux.SimpleMaps[r.Method][r.URL.Path]; ok {
		k(w, r)
		return
	}

	for reg,h := range(mux.RegexMaps[r.Method]) {
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

func Register() *Mux {
	mux := &Mux{}

	mux.SimpleMaps = make(map[string] map[string] http.HandlerFunc)
	mux.RegexMaps = make(map[string] map[*regexp.Regexp] http.HandlerFunc)
	mux.SimpleMaps["GET"] = make(map[string]http.HandlerFunc)
	mux.RegexMaps["GET"] = make(map[*regexp.Regexp] http.HandlerFunc)
	mux.SimpleMaps["POST"] = make(map[string]http.HandlerFunc)
	mux.RegexMaps["POST"] = make(map[*regexp.Regexp] http.HandlerFunc)

	mux.Get("/hello", handler.GetHello)
	mux.Get("/hello.html", handler.GetHelloHTML)
	mux.Get("/hello.json", handler.GetHelloJson)
	mux.Get("/issues", handler.GetIssues)
	mux.Get("/issues/{id}", handler.GetIssue1)

	mux.Post("/hello", handler.PostHello)

	return mux
}

func main() {

	mux := Register()
	handler := http.Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))

}
