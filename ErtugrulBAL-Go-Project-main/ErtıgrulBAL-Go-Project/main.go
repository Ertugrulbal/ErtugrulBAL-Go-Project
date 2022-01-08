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
	//DB Connection Established
	db = driver.ConnectDB()
	// Controller declared
	Controller := Controller.Controller{}
	//Endpoints
	http.HandleFunc("/ConvertFromWebUrl", Controller.ConvertFromUrl(db))
	http.HandleFunc("/ConvertFromDeeplink", Controller.ConvertFromDeeplink(db))
	//Home Endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// message := "API Home"
		message := API{"Yemeksepeti Link Converter Backend Applicant Test Programming Started."}
		output, err := json.Marshal(message)
		checkError(err)
		fmt.Fprintf(w, string(output))

	})
	// WebServer Info
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
