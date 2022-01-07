package DeeplinkRepository

import (
	"database/sql"
	"log"

	"github.com/ertugrulbal/Object"
)

type DeeplinkRepository struct{}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (b DeeplinkRepository) GetDeeplinks(db *sql.DB, Deeplink Object.Deeplink, Deeplinks []Object.Deeplink) ([]Object.Deeplink, error) {
	// SELECT * FROM Deelinks;
	rows, err := db.Query("SELECT * FROM Deeplinks")
	if err != nil {
		return []Object.Deeplink{}, err
	}

	for rows.Next() {
		err = rows.Scan(&Deeplink.ID, &Deeplink.LinkType, &Deeplink.ContentId, &Deeplink.CampaignId, &Deeplink.MerchantId, &Deeplink.Raw, &Deeplink.SearchQuery)
		Deeplinks = append(Deeplinks, Deeplink)
	}

	if err != nil {
		return []Object.Deeplink{}, err
	}
	return Deeplinks, nil
}

func (b DeeplinkRepository) ConvertFromDeeplink(db *sql.DB, Deeplink Object.Deeplink, id int) (Object.Deeplink, error) {
	err := db.QueryRow("INSERT INTO Deeplinks(ID, LinkType, ContentId, CampaignId, MerchantId, Raw, SearchQuery ) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id;",
		&Deeplink.ID, &Deeplink.LinkType, &Deeplink.ContentId, &Deeplink.CampaignId, &Deeplink.MerchantId, &Deeplink.Raw, &Deeplink.SearchQuery)

	if err != nil {

		return Object.Deeplink{}, nil
	}

	return Deeplink, nil
}
