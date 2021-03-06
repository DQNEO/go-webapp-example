package router

import (
	"net/http"
	"regexp"
	"strings"
)

type Router struct {
	SimpleMaps map[string]map[string]http.HandlerFunc
	RegexMaps  map[string]map[*regexp.Regexp]http.HandlerFunc
}

var URLParam [][]string

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
			URLParam = matches
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
