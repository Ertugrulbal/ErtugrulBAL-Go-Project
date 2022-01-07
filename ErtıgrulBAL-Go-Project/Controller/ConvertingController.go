package Controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ertugrulbal/Convert"
	"github.com/ertugrulbal/Error"
	"github.com/ertugrulbal/Object"
	DeeplinkRepository "github.com/ertugrulbal/Repository/DeepLink"
	WebURLRepository "github.com/ertugrulbal/Repository/WebURL"
	"github.com/ertugrulbal/utils"
)

type Controller struct{}

func (c Controller) ConvertFromUrl(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var WebURL Object.WebUrl
		var error Error.Error

		json.NewDecoder(r.Body).Decode(&WebURL)

		if WebURL.AddressType == "web" {
			WebURL.Parse(WebURL.Raw)
			Convert.ConvertToDeeplink(&Object.WebUrl{Raw: WebURL.Raw})
			fmt.Println("From WebURL to Deeplink conversion succeed.")
		}
		WebURLRepo := WebURLRepository.WebURLRepository{}
		rowsAffected, err := WebURLRepo.ConvertFromUrl(db, WebURL, WebURL.ID)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsAffected)

	}
}
func (c Controller) ConvertFromDeeplink(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Deeplink Object.Deeplink
		var error Error.Error

		json.NewDecoder(r.Body).Decode(&Deeplink)

		if r.RequestURI == "Deeplink" {
			Deeplink.Parse(Deeplink.Raw)
			Convert.ConvertToWebUrl(&Object.Deeplink{Raw: Deeplink.Raw})
			fmt.Println("From WebURL to Deeplink conversion succeed.")
		}
		DeeplinkRepo := DeeplinkRepository.DeeplinkRepository{}
		rowsAffected, err := DeeplinkRepo.ConvertFromDeeplink(db, Deeplink, Deeplink.ID)

		if err != nil {
			error.Message = "Server Error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsAffected)

	}
}
