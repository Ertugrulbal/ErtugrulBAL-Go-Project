package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ertugrulbal/Controller"
	"github.com/ertugrulbal/driver"
)

var db *sql.DB

func main() {
	db = driver.ConnectDB()

	Controller := Controller.Controller{}

	http.HandleFunc("/ConvertFromWebUrl", Controller.ConvertFromUrl(db))
	http.HandleFunc("/ConvertFromDeeplink", Controller.ConvertFromDeeplink(db))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// message := "API Home"
		message := API{"Yemeksepeti Link Converter Backend Applicant Test Programming Started."}
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))

	})

	http.ListenAndServe(":8080", nil)
}

type API struct {
	Message string `json:message`
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal Error :", err.Error())
		os.Exit(1)
	}
}
