package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)


var mappings = map[string]string{
	"boutiqueId":   "CampaignId",
	"merchantId":   "MerchantId",
	"sr&q":         "Search&Query",
	"Product":      "brand/",
	"ContentId":    "name-p-",
	"CampaignId":   "boutiqueId",
	"MerchantId":   "merchantId",
	"Search&Query": "sr&q",
}

func main() {
	router := gin.Default()

	router.POST("/ConvertEachOther", Converting)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {

}

type url struct {
	Url string `json:"url" binding:"required"`
}

func productDetailPage(urlStruct url) string {
	var url_ string
	split1 := strings.Split(urlStruct.Url, "-p-")
	url_ = "ec://?Page=Product&ContentId="

	split2 := strings.Split(split1[1], "?")
	productId := split2[0]
	url_ += productId + "&"

	queryParams := strings.Split(split2[1], "&")

	for i, param := range queryParams {
		pair := strings.Split(param, "=")

		var key string
		var val string

		if _, ok := mappings[pair[0]]; ok {
			key = mappings[pair[0]]
		} else {
			continue
		}

		val = pair[1]

		url_ += key + "=" + val

		if i != len(queryParams)-1 {
			url_ += "&"
		}
	}

	return url_
}

func productDetailPage_(urlStruct url) string {
	var url_ string
	split1 := strings.Split(urlStruct.Url, "?")
	url_ = "https://www.ecommerce.com/"

	queryParams := strings.Split(split1[1], "&")

	for i, param := range queryParams {
		pair := strings.Split(param, "=")

		if len(pair) != 2 {
			continue
		}

		key := pair[0]
		val := pair[1]

		if _, ok := mappings[key]; ok {
			key = mappings[key]
		} else {
			continue
		}

		if key == "name-p-" {
			url_ += key + val + "?"
		} else {
			url_ += key + "=" + val
			if i != len(queryParams)-1 {
				url_ += "&"
			}
		}
	}

	return url_
}

func searchAndQuery(urlStruct url) string {
	var url_ string
	split1 := strings.Split(urlStruct.Url, "sr?q=")
	url_ = "ec://?Page=Search&Query="

	url_ += split1[1]

	return url_
}

func searchAndQuery_(urlStruct url) string {
	var url_ string

	url_ = urlStruct.Url
	url_ = strings.Replace(url_, "ec://?Page=Search&Query", "https://www.ecommerce.com/sr?q", 1)
	return url_
}

func Converting(c *gin.Context) {

	var url_ string
	var instance url

	err := c.BindJSON(&instance)
	if err != nil {
		return
	}

	fmt.Printf("URL to be converted to Deeplink: %v\n", instance.Url)

	isDeepLink, _ := regexp.MatchString("ec://", instance.Url)
	isProductUrl_, _ := regexp.MatchString("Product&ContentId", instance.Url)
	isSearchAndQuery_, _ := regexp.MatchString("Search&Query", instance.Url)

	isWebUrl, _ := regexp.MatchString("ecommerce.com", instance.Url)
	isProductUrl, _ := regexp.MatchString("-p-", instance.Url)
	isSearchAndQuery, _ := regexp.MatchString("sr\\?q=", instance.Url)

	if !isWebUrl && !isDeepLink {
		c.IndentedJSON(http.StatusBadRequest, "Given URL is neither web url nor deep link!")
		return
	}

	if !isProductUrl && !isSearchAndQuery && isWebUrl {
		url_ = "ec://?Page=Home"
	}

	if isProductUrl && isWebUrl {
		url_ = productDetailPage(instance)
	}

	if isSearchAndQuery && isWebUrl {
		url_ = searchAndQuery(instance)
	}

	if !isProductUrl_ && !isSearchAndQuery_ && isDeepLink {
		url_ = "https://www.ecommerce.com"
	}

	if isProductUrl_ && isDeepLink {
		url_ = productDetailPage_(instance)
	}

	if isSearchAndQuery_ && isDeepLink {
		url_ = searchAndQuery_(instance)
	}

	c.IndentedJSON(http.StatusOK, url_)

}
