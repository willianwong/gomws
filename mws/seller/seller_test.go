package seller

import (
	"fmt"
	"github.com/willianwong/gomws/mws"
	"testing"
)

func TestSeller_GetServiceStatus(t *testing.T) {
	config := mws.Config{
		SellerId:  "xx",
		AuthToken: "xx",
		Region:    "DE",

		// Optional if set in env variable
		AccessKey: "xx",
		SecretKey: "xx",
	}

	productsClient, err := NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Example 1
	fmt.Println("------GetServiceStatus------")
	statusResponse, err := productsClient.GetServiceStatus()
	// Check http client error.
	if err != nil {
		fmt.Println(err.Error())
	}
	defer statusResponse.Close()
	// Check whether or not the API return errors.
	if statusResponse.Error != nil {
		fmt.Println(statusResponse.Error.Error())
	} else {
		xmlNode, _ := statusResponse.ResultParser()
		xmlNode.PrintXML() // Print the xml response with indention.
	}
}
