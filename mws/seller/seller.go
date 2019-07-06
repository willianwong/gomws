package seller

// Reference http://docs.developer.amazonservices.com/en_US/products/Products_Overview.html

import (
	"github.com/willianwong/gomws/mws"
)

// Products is the client for the api
type Seller struct {
	*mws.Client
}

// NewClient generate a new product client
func NewClient(config mws.Config) (*Seller, error) {
	seller := new(Seller)
	base, err := mws.NewClient(config, seller.Version(), seller.Name())
	if err != nil {
		return nil, err
	}
	seller.Client = base
	return seller, nil
}

// Version return the current version of api
func (s Seller) Version() string {
	return "2011-07-01"
}

// Name return the name of the api
func (s Seller) Name() string {
	return "Sellers"
}

// GetServiceStatus Returns the operational status of the Products API section.
// http://docs.developer.amazonservices.com/zh_CN/sellers/Sellers_GetServiceStatus.html
func (s Seller) GetServiceStatus() (*mws.Response, error) {
	params := mws.Parameters{
		"Action": "GetServiceStatus",
	}
	return s.SendRequest(params)
}
