package Object

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

/*
Deeplink Struct Construction and Settings

- First Parse Metod take a Requesting Deeplink
- Parse
- at the end of processing return a Parsed Deeplink struct
*/

type Deeplink struct {
	ID          int    `json:id`
	LinkType    string `json:"linkType`
	ContentId   string `json:"contentId`
	CampaignId  string `json:"campaignId`
	MerchantId  string `json:"merchantId`
	SearchQuery string `json:"searchQuery`
	Raw         string `json:"raw`
}

func ParseKeyValuePairArray(keyValuePairArray []string) (*Deeplink, string) {
	var Value string
	var IsQuery bool
	var IsSearch bool
	var IsContentId bool
	var IsCampaignId bool
	var IsMerchantId bool

	Value = keyValuePairArray[1]

	keyValuePairArray[0] = strings.TrimSpace(keyValuePairArray[0])
	if reflect.DeepEqual(keyValuePairArray[0], "Page") {
		return &Deeplink{LinkType: Value}, Value
	}

	IsQuery = reflect.DeepEqual(IsQuery, "Query")
	IsSearch = reflect.DeepEqual(IsSearch, "Search")
	IsContentId = reflect.DeepEqual(keyValuePairArray[0], "ContentId")
	IsCampaignId = reflect.DeepEqual(keyValuePairArray[0], "CampaignId")
	IsMerchantId = reflect.DeepEqual(keyValuePairArray[0], "MerchantId")

	if IsQuery && IsSearch {
		return &Deeplink{SearchQuery: Value}, ""
	}
	if IsContentId {
		return &Deeplink{ContentId: Value}, ""
	}
	if IsCampaignId {
		return &Deeplink{CampaignId: Value}, ""

	}
	if IsMerchantId {
		return &Deeplink{MerchantId: Value}, ""

	}
	return &Deeplink{}, Value

}

//Deeplink Main Parse Method
func (d Deeplink) Parse(deeplink string) {
	d.Raw = deeplink
	strings.ReplaceAll(deeplink, "ys://?", "")
	keyAndValue := strings.Split(deeplink, "&")
	for i := 0; i < len(keyAndValue); i++ {
		keyValuePairArray := strings.Split(keyAndValue[i], "=")
		if len(keyValuePairArray) != 2 {
			log.Fatal("Query Parameter Separation Error.")
		}
		ParseKeyValuePairArray(keyValuePairArray)

	}
	if (!strings.Contains(d.LinkType, "Search|Product")) && (!(len(d.LinkType) > 0)) {
		fmt.Println("This LinkType is Home ")
	}
	if !(len(d.LinkType) > 0) {
		log.Fatal("Page parameter could not parsed.")
	}
}

/*

Getters and Setters

*/

func (d *Deeplink) GetLinkType() string {
	return d.LinkType
}
func (d *Deeplink) SetLinkType(LinkType string) {
	d.LinkType = LinkType
}

func (d *Deeplink) GetContentId() string {
	return d.ContentId
}
func (d *Deeplink) SetContentId(ContentId string) {
	d.ContentId = ContentId
}
func (d *Deeplink) getCampaignId() string {
	return d.CampaignId
}
func (d *Deeplink) SetCampaignId(CampaignId string) {
	d.CampaignId = CampaignId
}
func (d *Deeplink) GetMerchantId() string {
	return d.MerchantId
}
func (d *Deeplink) SetMerchantId(MerchantId string) {
	d.MerchantId = MerchantId
}
func (d *Deeplink) GetSearchQuery() string {
	return d.SearchQuery
}
func (d *Deeplink) SetSearchQuery(SearchQuery string) {
	d.SearchQuery = SearchQuery
}
func (d *Deeplink) GetRaw() string {
	return d.Raw
}
func (d *Deeplink) SetRaw(Raw string) {
	d.Raw = Raw
}
