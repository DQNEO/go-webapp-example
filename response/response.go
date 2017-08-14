package response

import "net/http"
import "encoding/json"

func SucceedWithNoContent(w http.ResponseWriter) {
	w.WriteHeader(204)
}

func Succeed(w http.ResponseWriter, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func Fail(w http.ResponseWriter, err error) {
	w.WriteHeader(500) // or can be 404
	w.Write([]byte(err.Error()))
}
