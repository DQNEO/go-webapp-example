package main

import "net/http"
import "log"
import "./router"
import "./route"
import (
	"fmt"
)

const defaultPort = 9001

func main() {
	rt := router.NewRouter()
	rt = route.Register(rt)
	handler := http.Handler(rt)
	port := defaultPort
	log.Printf("starting server : http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))

}
