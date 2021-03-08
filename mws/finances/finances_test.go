package finances

import (
	"fmt"
	"github.com/willianwong/gomws/mws"
	"testing"
	"time"
)

func TestFinances_ListFinancialEventGroups(t *testing.T) {
	config := mws.Config{
		SellerId:  "",
		AuthToken: "",
		Region:    "UK",

		// Optional if set in env variable
		AccessKey: "",
		SecretKey: "",
	}

	financesClient, err := NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Example 1
	fmt.Println("------ListFinancialEventGroups------")
	maps := make(map[string]interface{})
	maps["FinancialEventGroupStartedAfter"] = time.Unix(1609430400, 0).Format(time.RFC3339)
	res, err := financesClient.ListFinancialEventGroups(maps)
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

func TestFinances_ListFinancialEvents(t *testing.T) {
	config := mws.Config{
		SellerId:  "",
		AuthToken: "",
		Region:    "UK",

		// Optional if set in env variable
		AccessKey: "",
		SecretKey: "",
	}

	financesClient, err := NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Example 1
	fmt.Println("------ListFinancialEvents------")
	maps := make(map[string]interface{})
	maps["AmazonOrderId"] = "xxx"
	res, err := financesClient.ListFinancialEvents(maps)
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
