package response

import "net/http"
import "encoding/json"

func SendJson(w http.ResponseWriter, v interface{}) {
	b,err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}



