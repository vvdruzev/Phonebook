package util
import (
	"encoding/json"
	"net/http"
)

func ResponseOk(w http.ResponseWriter, body interface{}) {
	b,_:= json.Marshal(body)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	body := map[string]string{
		"error": message,
	}
	b,_:= json.Marshal(body)

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}