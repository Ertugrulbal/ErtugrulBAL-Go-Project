package utils

import (
	"encoding/json"
	"net/http"

	"github.com/ertugrulbal/Error"
)

/*
Web Server Error Messages Management Function
*/
func SendError(w http.ResponseWriter, status int, error Error.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(error)
}

/*
Web Server Send Success Messages Management Function
*/
func SendSuccess(w http.ResponseWriter, data interface{}) {
	// fmt.Println(data)
	json.NewEncoder(w).Encode(data)
}
