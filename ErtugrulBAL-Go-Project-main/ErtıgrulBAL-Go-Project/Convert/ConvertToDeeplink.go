package Convert

import (
	"bytes"
	"reflect"

	Object "github.com/ertugrulbal/Object"
)

/*
In this section we created a Converting function which convert from WebUrl to Deeplink
*/

// Main Function
func ConvertToDeeplink(WebUrl *Object.WebUrl) bytes.Buffer {
	var builder bytes.Buffer
	var deeplink bytes.Buffer

	builder.WriteString("ec://?Page=")

	builder.WriteString(WebUrl.AddressType)

	if reflect.DeepEqual(WebUrl.AddressType, "Product") {
		input1 := PrepareSearchDeeplink(WebUrl)
		builder.WriteString(input1.String())
	}

	if reflect.DeepEqual(WebUrl.AddressType, "Search") {
		input2 := PrepareProductDeeplink(WebUrl)
		builder.WriteString(input2.String())
	}
	builder = deeplink

	return deeplink

}

// İf URL has a "Search" parameter, this function work and return the new builder
func PrepareSearchDeeplink(WebUrl *Object.WebUrl) bytes.Buffer {
	var builder bytes.Buffer
	if WebUrl.SearchQuery != "" {
		builder.WriteString("&Query=")
		builder.WriteString(WebUrl.SearchQuery)
	}
	return builder

}

// İf URL has a "Product" parameter, this function work and return the new builder
func PrepareProductDeeplink(WebUrl *Object.WebUrl) bytes.Buffer {
	var builder bytes.Buffer
	if WebUrl.ContentId != "" {
		builder.WriteString("&ContentId=")
		builder.WriteString(WebUrl.ContentId)
	}
	if WebUrl.BoutiqueId != "" {
		builder.WriteString("&CampaignId=")
		builder.WriteString(WebUrl.BoutiqueId)
	}
	if WebUrl.MerchantId != "" {
		builder.WriteString("&MerchantId=")
		builder.WriteString(WebUrl.MerchantId)
	}
	return builder
}
