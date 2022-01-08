package Convert

import (
	"bytes"
	"reflect"

	Object "github.com/ertugrulbal/Object"
)

func ConvertToWebUrl(Deeplink *Object.Deeplink) bytes.Buffer {
	var builder bytes.Buffer
	var webUrl bytes.Buffer

	builder.WriteString("https://www.ecommerce.com/")

	if reflect.DeepEqual(Deeplink.LinkType, "Product") {
		input1 := prepareProductWebURL(Deeplink)
		builder.WriteString(input1.String())
	}

	if reflect.DeepEqual(Deeplink.LinkType, "Search") {
		input2 := prepareSearchWebURL(Deeplink)
		builder.WriteString(input2.String())
	}
	builder = webUrl

	return webUrl

}
func prepareProductWebURL(Deeplink *Object.Deeplink) bytes.Buffer {
	var builder bytes.Buffer
	var isContainsMerchantId bool
	var isContainsBoutiqueId bool

	builder.WriteString("brand/name-p-")
	builder.WriteString(Deeplink.ContentId)

	isContainsMerchantId = (Deeplink.MerchantId != "")
	isContainsBoutiqueId = (Deeplink.CampaignId != "")
	if isContainsBoutiqueId || isContainsMerchantId {
		builder.WriteString("?")
	}
	if isContainsBoutiqueId {
		builder.WriteString("boutiqueId=")
		builder.WriteString(Deeplink.CampaignId)
	}
	if isContainsBoutiqueId && isContainsMerchantId {
		builder.WriteString("&")
	}
	if isContainsMerchantId {
		builder.WriteString("merchantId=")
		builder.WriteString(Deeplink.MerchantId)
	}

	if Deeplink.SearchQuery != "" {
		builder.WriteString("&Query=")
		builder.WriteString(Deeplink.SearchQuery)
	}
	return builder

}
func prepareSearchWebURL(Deeplink *Object.Deeplink) bytes.Buffer {
	var builder bytes.Buffer

	builder.WriteString("sr?q=")

	if Deeplink.SearchQuery != "" {
		builder.WriteString(Deeplink.SearchQuery)
	}
	return builder
}
