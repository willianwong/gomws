package orders

import (
	"fmt"
	"github.com/willianwong/gomws/mws"
	"testing"
	"time"
)

func TestOrders_ListOrders(t *testing.T) {
	config := mws.Config{
		SellerId:  "xxx",
		AuthToken: "xxx",
		Region:    "UK",

		// Optional if set in env variable
		AccessKey: "xxx",
		SecretKey: "xxx",
	}

	ordersClient, err := NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Example 1
	fmt.Println("------GetOrder------")
	maps := make(map[string]interface{})
	maps["LastUpdatedAfter"] = time.Unix(1551746540, 0).Format(time.RFC3339)
	//maps["OrderStatus"] = []string{"Unshipped","PartiallyShipped"}
	listOrders, err := ordersClient.ListOrders(maps)
	// Check http client error.
	if err != nil {
		fmt.Println(err.Error())
	}
	defer listOrders.Close()
	// Check whether or not the API return errors.
	if listOrders.Error != nil {
		fmt.Println(listOrders.Error.Error())
	} else {
		xmlNode, _ := listOrders.ResultParser()
		xmlNode.PrintXML() // Print the xml response with indention.
	}
}

func TestOrders_GetOrder(t *testing.T) {
	config := mws.Config{
		SellerId:  "xxx",
		AuthToken: "xxx",
		Region:    "UK",

		// Optional if set in env variable
		AccessKey: "xxx",
		SecretKey: "xxx",
	}

	ordersClient, err := NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Example 1
	fmt.Println("------GetOrder------")
	orderDetail, err := ordersClient.GetOrder([]string{})
	// Check http client error.
	if err != nil {
		fmt.Println(err.Error())
	}
	defer orderDetail.Close()
	// Check whether or not the API return errors.
	if orderDetail.Error != nil {
		fmt.Println(orderDetail.Error.Error())
	} else {
		xmlNode, _ := orderDetail.ResultParser()
		xmlNode.PrintXML() // Print the xml response with indention.
	}
}
