package feeds

// Reference http://docs.developer.amazonservices.com/en_US/products/Products_Overview.html

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/willianwong/gomws/mws"
)

// Products is the client for the api
type Feed struct {
	*mws.Client
}

// NewClient generate a new product client
func NewClient(config mws.Config) (*Feed, error) {
	feed := new(Feed)
	base, err := mws.NewClient(config, feed.Version(), feed.Name())
	if err != nil {
		return nil, err
	}
	feed.Client = base
	return feed, nil
}

// Version return the current version of api
func (f Feed) Version() string {
	return "2009-01-01"
}

// Name return the name of the api
func (f Feed) Name() string {
	return "Feeds"
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// GetServiceStatus Returns the operational status of the Products API section.
// http://docs.developer.amazonservices.com/zh_CN/sellers/Sellers_GetServiceStatus.html
func (f Feed) SubmitFeed(feedContent, feedType string, marketplaceIds []string) (*mws.Response, error) {
	params := mws.Parameters{
		"Action":      "SubmitFeed",
		"FeedType":    feedType,
		"FeedContent": feedContent,
	}
	structuredParams := params.StructureKeys("MarketplaceIdList", "Id")

	return f.SendRequest(structuredParams)
}

func (f Feed) GetFeedSubmissionList() (*mws.Response, error) {
	params := mws.Parameters{
		"Action": "GetFeedSubmissionList",
	}
	return f.SendRequest(params)
}

func (f Feed) GetFeedSubmissionResult(feedSubmissionId string) (*mws.Response, error) {
	params := mws.Parameters{
		"Action":           "GetFeedSubmissionResult",
		"FeedSubmissionId": feedSubmissionId,
	}
	return f.SendRequest(params)
}
