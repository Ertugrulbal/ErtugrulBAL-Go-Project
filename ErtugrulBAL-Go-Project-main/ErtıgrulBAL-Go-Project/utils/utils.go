package utils

import (
	"encoding/json"
	"net/http"

	"github.com/ertugrulbal/Error"
)

func SendError(w http.ResponseWriter, status int, error Error.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	// fmt.Println(data)
	json.NewEncoder(w).Encode(data)
}
