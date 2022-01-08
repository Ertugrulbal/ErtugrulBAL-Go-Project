package Object

import (
	"fmt"
	"log"
	"reflect"
	"strings"
)

type WebUrl struct {
	ID              int    `json:id`
	AddressType     string `json:"addressType"`
	BrandOrCategory string `json:"brandOrCategory"`
	ProductName     string `json:"productName"`
	ContentId       string `json:"contentId"`
	BoutiqueId      string `json:"boutiqueId"`
	MerchantId      string `json:"merchantId"`
	SearchQuery     string `json:"searchQuery"`
	Raw             string `json:"raw"`
}

func (w *WebUrl) Parse(url string) {
	w.Raw = url
	strings.ReplaceAll(url, "https://www.ecommerce.com/", "")
	if strings.Contains(url, "-p-") {
		ParseProductAddress(url)
	} else if strings.Contains(url, "search?q=") {
		ParseSearchAddress(url)

	} else {
		w.AddressType = "Home"
	}
}

func ParseBrandOrCategory(BrandOrCategory string) *WebUrl {
	if len(BrandOrCategory) < 1 {
		fmt.Println("Brand or Category could not parsed!")
	}
	return &WebUrl{BrandOrCategory: BrandOrCategory}
}

func ParseProductNameAndContentId(ProductNameAndContentId string) *WebUrl {
	strings.ReplaceAll(ProductNameAndContentId, "//?", "?")
	ProductNameAndContentIdList := strings.Split(ProductNameAndContentId, "\\?")
	ProductNameList := strings.Split(ProductNameAndContentIdList[0], "-p-")
	ProductName := ProductNameList[0]

	if len(ProductName) < 1 {
		log.Fatal("ProductName could not parsed.")
	}

	ContentID := ProductNameList[1]
	if len(ContentID) < 1 {
		log.Fatal("ContentId could not parsed.")
	}
	return &WebUrl{ProductName: ProductName, ContentId: ContentID}
}

func ParseQueryParameters(queryParameters string) *WebUrl {
	BoutiqueId := ""
	MerchantId := ""

	keyAndValue := strings.Split(queryParameters, "&")
	for i := 0; i < len(keyAndValue); i++ {
		keyValuePairArray := strings.Split(keyAndValue[i], "=")
		if len(keyValuePairArray) != 2 {
			log.Fatal("Query Parameter Seperation Error.")
		}
		Value := keyValuePairArray[1]
		keyValuePairArray[0] = strings.TrimSpace(keyValuePairArray[0])
		isBoutiqueId := reflect.DeepEqual(keyValuePairArray[0], "boutiqueId")
		if isBoutiqueId {
			BoutiqueId = Value
			continue
		}
		isMerchantId := reflect.DeepEqual(keyValuePairArray[0], "isMerchantId")
		if isMerchantId {
			MerchantId = Value
		}

	}
	return &WebUrl{BoutiqueId: BoutiqueId, MerchantId: MerchantId}
}

func ParseProductAddress(url string) *WebUrl {
	parts := strings.Split(url, "/")
	ParseBrandOrCategory(parts[0])
	ParseProductNameAndContentId(parts[1])

	if !strings.Contains(parts[1], "?") {
		log.Fatal("err")

	}
	strings.Split(parts[1], "\\?")
	strings.ReplaceAll(parts[1], "?", "")
	ParseQueryParameters(parts[1])
	return &WebUrl{AddressType: "Product"}
}

func ParseSearchAddress(url string) *WebUrl {
	SearchQuery := strings.ReplaceAll(url, "search?q=", "")
	if len(SearchQuery) < 1 || strings.TrimSpace(SearchQuery) == "" {
		log.Fatal("searchQuery could not parsed.")
	}
	return &WebUrl{SearchQuery: SearchQuery, AddressType: "Search"}
}

func (w *WebUrl) GetAddressType() string {
	return w.AddressType
}

func (w *WebUrl) GetBrandOrCategory() string {
	return w.BrandOrCategory
}

func (w *WebUrl) GetProductName() string {
	return w.ProductName
}

func (w *WebUrl) GetContentId() string {
	return w.ContentId
}

func (w *WebUrl) GetBoutiqueId() string {
	return w.BoutiqueId
}

func (w *WebUrl) GetMerchantId() string {
	return w.MerchantId
}

func (w *WebUrl) GetSearchQuery() string {
	return w.SearchQuery
}

func (w *WebUrl) GetRaw() string {
	return w.Raw
}
