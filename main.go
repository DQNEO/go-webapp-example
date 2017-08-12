package main

import "./route"
import "net/http"
import "log"

func main() {

	route.Register()

	log.Fatal(http.ListenAndServe(":8080", nil))

}
