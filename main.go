package main

import "net/http"
import "log"
import "./handler"
import "fmt"
import "regexp"

type Mux struct {
	Matched string
	GetMap map[string] http.HandlerFunc
}

func(mux *Mux) Get(pattern string, h func(http.ResponseWriter, *http.Request)) {
	if mux.GetMap == nil {
		mux.GetMap = make(map[string]http.HandlerFunc)
	}
	mux.GetMap[pattern] = http.HandlerFunc(h)
}

func(mux *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// test complete match
		if k,ok := mux.GetMap[r.URL.Path]; ok {
			k(w, r)
			return
		}
		reg := regexp.MustCompile(`/issues/([a-zA-Z0-9]+)`)
		if matches := reg.FindAllStringSubmatch(r.URL.Path, -1); matches != nil {
			handler.URLParam = matches
			handler.GetIssue1(w, r)
			return
		} else {
			handler.GetHello(w, r)
			return
		}
	} else if r.Method == "POST" {
		fmt.Fprintf(w, "POST\n", r.URL.Path)
	}

}

func Register() *Mux {
	mux := &Mux{}

	mux.Get("/hello", handler.GetHello)
	mux.Get("/hello.html", handler.GetHelloHTML)
	mux.Get("/hello.json", handler.GetHelloJson)
	mux.Get("/issues", handler.GetIssues)
	mux.Get(`/issues/([a-zA-Z0-9]+)`, handler.GetIssue1)

	return mux
}

func main() {

	mux := Register()
	handler := http.Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))

}
