package handler

import "net/http"
import "fmt"
import "../model"
import (
	"../response"
	"strconv"
)

var URLParam [][]string

func GetIssue(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(URLParam[0][1]) // ugly!
	if err != nil {
		response.Fail(w, err) // should be 404
		return
	}

	issue, err := model.FindIssue(id)
	if err != nil {
		response.Fail(w, err) // can be 404 or 500
		return
	}

	response.Succeed(w, issue)
}

func GetIssues(w http.ResponseWriter, r *http.Request) {
	issues, err := model.GetIssues()
	if err != nil {
		response.Fail(w, err)
		return
	}
	response.Succeed(w, issues)
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"msg":"hello"}`)
}

func PostHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "post hello")
}

func GetHelloHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func GetHelloJson(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `{"msg":"hello"}`)
}
