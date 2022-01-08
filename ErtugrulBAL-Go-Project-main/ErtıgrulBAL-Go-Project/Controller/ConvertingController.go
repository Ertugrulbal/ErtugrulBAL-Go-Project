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

/* In this section we created a function which provide to us for converting controlling processings. */

// ConvertFromUrl take a Web URL and convert to Deeplink

func (c Controller) ConvertFromUrl(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var WebURL Object.WebUrl
		var error Error.Error

		json.NewDecoder(r.Body).Decode(&WebURL)

		WebURL.Parse(WebURL.Raw)
		fmt.Println(Convert.ConvertToDeeplink(&Object.WebUrl{}))

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

// ConvertFromDeeplink take a Deeplink and convert to WebURL
// Also DB Connection installed at the end of Converting processings
func (c Controller) ConvertFromDeeplink(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var Deeplink Object.Deeplink
		var error Error.Error

		json.NewDecoder(r.Body).Decode(&Deeplink)
		Deeplink.Parse(Deeplink.Raw)
		fmt.Println(Convert.ConvertToWebUrl(&Object.Deeplink{}))

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
