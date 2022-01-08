package WebURLRepository

import (
	"database/sql"
	"log"

	"github.com/ertugrulbal/Object"
)

type WebURLRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (b WebURLRepository) GetWebURLs(db *sql.DB, WebURL Object.WebUrl, WebURLs []Object.WebUrl) ([]Object.WebUrl, error) {
	// SELECT * FROM WebURLs;
	rows, err := db.Query("SELECT * FROM WebURLs")
	if err != nil {
		return []Object.WebUrl{}, err
	}

	for rows.Next() {
		err = rows.Scan(&WebURL.ID, &WebURL.AddressType, &WebURL.BrandOrCategory, &WebURL.ProductName, &WebURL.ContentId, &WebURL.BoutiqueId, &WebURL.MerchantId, &WebURL.SearchQuery, &WebURL.Raw)
		WebURLs = append(WebURLs, WebURL)
	}

	if err != nil {
		return []Object.WebUrl{}, err
	}
	return WebURLs, nil
}

func (b WebURLRepository) ConvertFromUrl(db *sql.DB, WebURL Object.WebUrl, id int) (Object.WebUrl, error) {
	err := db.QueryRow("INSERT INTO WebURLs(ID, addressType, brandOrCategory, productName, contentId, boutiqueId, merchantId, searchQuery, raw ) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9,) RETURNING id;",
		&WebURL.ID, &WebURL.AddressType, &WebURL.BrandOrCategory, &WebURL.ProductName, &WebURL.ContentId, &WebURL.BoutiqueId, &WebURL.MerchantId, &WebURL.SearchQuery, &WebURL.Raw)

	if err != nil {

		return Object.WebUrl{}, nil
	}

	return WebURL, nil
}
