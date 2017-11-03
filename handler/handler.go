package handler

import "net/http"
import "fmt"
import "../model"
import (
	"../response"
	"../router"
	"strconv"
)

func GetIssue(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(router.URLParam[0][1]) // ugly!
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

func GetAllIssues(w http.ResponseWriter, r *http.Request) {
	issues, err := model.GetIssues()
	if err != nil {
		response.Fail(w, err)
		return
	}
	response.Succeed(w, issues)
}

func SearchIssues(w http.ResponseWriter, r *http.Request) {
	issues, err := model.GetIssues()
	if err != nil {
		response.Fail(w, err)
		return
	}

	//name := r.URL.Query("name")

	response.Succeed(w, issues)
}

func CreateIssue(w http.ResponseWriter, r *http.Request) {
	response.SucceedWithNoContent(w)
}

func UpdateIssue(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"msg":"ok"}`)
}

func DeleteIssue(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"msg":"ok"}`)
}

func PutHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"msg":"put hello"}`)
}

func DeleteHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"msg":"delete hello"}`)
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"msg":"hello"}`)
}

func PostHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"msg":"post hello"}`)
}
