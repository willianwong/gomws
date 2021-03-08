package reports

import (
	"fmt"
	"github.com/willianwong/gomws/mws"
	"testing"
	"time"
)

func TestReports_RequestReport(t *testing.T) {
	config := mws.Config{
		SellerId:  "",
		AuthToken: "",
		Region:    "UK",

		// Optional if set in env variable
		AccessKey: "",
		SecretKey: "",
	}

	reportsClient, err := NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Example 1
	fmt.Println("------ReportRequest------")
	maps := make(map[string]interface{})
	maps["StartDate"] = time.Unix(1609430400, 0).Format(time.RFC3339)
	maps["EndDate"] = time.Unix(1612108800, 0).Format(time.RFC3339)
	maps["MarketplaceIdList"] = []string{"A1F83G8C2ARO7P"}
	res, err := reportsClient.RequestReport("_GET_FLAT_FILE_ACTIONABLE_ORDER_DATA_", maps)
	// Check http client error.
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Close()
	// Check whether or not the API return errors.
	if res.Error != nil {
		fmt.Println(res.Error.Error())
	} else {
		xmlNode, _ := res.ResultParser()
		xmlNode.PrintXML() // Print the xml response with indention.
	}
}

func TestReports_GetReportRequestList(t *testing.T) {
	config := mws.Config{
		SellerId:  "",
		AuthToken: "",
		Region:    "UK",

		// Optional if set in env variable
		AccessKey: "",
		SecretKey: "",
	}

	reportsClient, err := NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Example 1
	fmt.Println("------ReportRequest------")
	//maps := make(map[string]interface{})
	//maps["StartDate"] = time.Unix(1609430400, 0).Format(time.RFC3339)
	//maps["EndDate"] = time.Unix(1612108800, 0).Format(time.RFC3339)
	//maps["MarketplaceIdList"] = []string{"A1F83G8C2ARO7P"}
	res, err := reportsClient.GetReportRequestList()
	// Check http client error.
	if err != nil {
		fmt.Println(err.Error())
	}
	defer res.Close()
	// Check whether or not the API return errors.
	if res.Error != nil {
		fmt.Println(res.Error.Error())
	} else {
		xmlNode, _ := res.ResultParser()
		xmlNode.PrintXML() // Print the xml response with indention.
	}
}
