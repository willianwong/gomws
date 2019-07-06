package feeds

import (
	"encoding/xml"
	"fmt"
	"github.com/willianwong/gomws/mws"
	"testing"
	"time"
)

func TestFeed_GetFeedSubmissionList(t *testing.T) {
	config := mws.Config{
		SellerId:  "xxx",
		AuthToken: "xxx",
		Region:    "DE",

		// Optional if set in env variable
		AccessKey: "xxx",
		SecretKey: "xxx",
	}

	feedsClient, err := NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Example 1
	fmt.Println("------GetFeedSubmissionList------")
	feedSubmissionList, err := feedsClient.GetFeedSubmissionList()
	// Check http client error.
	if err != nil {
		fmt.Println(err.Error())
	}
	defer feedSubmissionList.Close()
	// Check whether or not the API return errors.
	if feedSubmissionList.Error != nil {
		fmt.Println(feedSubmissionList.Error.Error())
	} else {
		xmlNode, _ := feedSubmissionList.ResultParser()
		xmlNode.PrintXML() // Print the xml response with indention.
	}
}

func TestFeed_GetFeedSubmissionResult(t *testing.T) {
	config := mws.Config{
		SellerId:  "xxx",
		AuthToken: "xxx",
		Region:    "UK",

		// Optional if set in env variable
		AccessKey: "xxx",
		SecretKey: "xxx",
	}

	feedsClient, err := NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Example 1
	fmt.Println("------GetFeedSubmissionResult------")
	feedSubmissionResult, err := feedsClient.GetFeedSubmissionResult("xxx")
	// Check http client error.
	if err != nil {
		fmt.Println(err.Error())
	}
	defer feedSubmissionResult.Close()
	// Check whether or not the API return errors.
	if feedSubmissionResult.Error != nil {
		fmt.Println(feedSubmissionResult.Error.Error())
	} else {
		xmlNode, _ := feedSubmissionResult.ResultParser()
		xmlNode.PrintXML() // Print the xml response with indention.
	}
}

func TestFeed_SubmitFeed(t *testing.T) {
	config := mws.Config{
		SellerId:  "xxx",
		AuthToken: "xxx",
		Region:    "UK",

		// Optional if set in env variable
		AccessKey: "xxx",
		SecretKey: "xxx",
	}

	feedsClient, err := NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Example 1
	fmt.Println("------SubmitFeed------")
	type OrderFulfillmentFeedItem struct {
		AmazonOrderItemCode string `xml:"AmazonOrderItemCode"`
		Quantity string `xml:"Quantity"`
	}
	type OrderFulfillmentFeedData struct {
		CarrierName string `xml:"CarrierName"`
		ShippingMethod string `xml:"ShippingMethod"`
		ShipperTrackingNumber string `xml:"ShipperTrackingNumber"`
	}
	type OrderFulfillmentFeedContent struct {
		AmazonOrderID string `xml:"AmazonOrderID"`
		FulfillmentDate string `xml:"FulfillmentDate"`
		FulfillmentData OrderFulfillmentFeedData `xml:"FulfillmentData"`
		Item []OrderFulfillmentFeedItem `xml:"Item"`
	}
	type OrderFulfillmentFeedMessage struct {
		MessageID string `xml:"MessageID"`
		OrderFulfillment OrderFulfillmentFeedContent `xml:"OrderFulfillment"`
	}
	type OrderFulfillmentFeedHeader struct {
		DocumentVersion string `xml:"DocumentVersion"`
		MerchantIdentifier string `xml:"MerchantIdentifier"`
	}
	type OrderFulfillmentFeedBase struct {
		XMLName xml.Name `xml:"AmazonEnvelope"`
		XMLNs string `xml:"xmlns:xsi,attr"`
		Xsi string `xml:"xsi:noNamespaceSchemaLocation,attr"`
		Header OrderFulfillmentFeedHeader `xml:"Header"`
		MessageType string `xml:"MessageType"`
		Message []OrderFulfillmentFeedMessage `xml:"Message"`
	}
	var content OrderFulfillmentFeedBase
	content.XMLNs = "http://www.w3.org/2001/XMLSchema-instance"
	content.Xsi = "amznenvelope.xsd"
	content.MessageType = "OrderFulfillment"
	content.Header.DocumentVersion = "1.02"
	content.Header.MerchantIdentifier = "MyStore"
	var contentMessage OrderFulfillmentFeedMessage
	contentMessage.MessageID = "1"
	contentMessage.OrderFulfillment.AmazonOrderID = "xxx"
	contentMessage.OrderFulfillment.FulfillmentDate = time.Unix(1561953872, 0).Format(time.RFC3339)
	contentMessage.OrderFulfillment.FulfillmentData.CarrierName = "Yun Express"
	contentMessage.OrderFulfillment.FulfillmentData.ShippingMethod = "Standard"
	contentMessage.OrderFulfillment.FulfillmentData.ShipperTrackingNumber = "xxx"
	var itemContent OrderFulfillmentFeedItem
	itemContent.AmazonOrderItemCode = "aaa"
	itemContent.Quantity = "1"
	contentMessage.OrderFulfillment.Item = append(contentMessage.OrderFulfillment.Item,itemContent)
	content.Message = append(content.Message,contentMessage)
	data, _ := xml.MarshalIndent(&content, "", "  ")
	submitFeedRes, err := feedsClient.SubmitFeed(string(data), "_POST_ORDER_FULFILLMENT_DATA_", []string{})
	// Check http client error.
	if err != nil {
		fmt.Println(err.Error())
	}
	defer submitFeedRes.Close()
	// Check whether or not the API return errors.
	if submitFeedRes.Error != nil {
		fmt.Println(submitFeedRes.Error.Error())
	} else {
		xmlNode, _ := submitFeedRes.ResultParser()
		xmlNode.PrintXML() // Print the xml response with indention.
	}
}

func TestMd5V(t *testing.T) {
	sourceStr := "1"
	str := mws.Base64EncodeStr(mws.Md5V(sourceStr))
	fmt.Println(str)
}
