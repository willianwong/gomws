package finances

import (
	"github.com/willianwong/gomws/mws"
)

// Reports is the client for the api
type Finances struct {
	*mws.Client
}

// NewClient generate a new product client
func NewClient(config mws.Config) (*Finances, error) {
	report := new(Finances)
	base, err := mws.NewClient(config, report.Version(), report.Name())
	if err != nil {
		return nil, err
	}
	report.Client = base
	return report, nil
}

// Version return the current version of api
func (f Finances) Version() string {
	return "2015-05-01"
}

// Name return the name of the api
func (f Finances) Name() string {
	return "Finances"
}

// RequestReport Creates a report request and submits the request to Amazon MWS.
// Optional Parameters:
// 	StartDate - string. The start of a date range used for selecting the data to report. Values in ISO 8601 date time format.
//  EndDate - string. The end of a date range used for selecting the data to report. Values in ISO 8601 date time format.
//  ReportOptions - string. Additional information to pass to the report.
//  MarketplaceIdList - []string. A list of one or more marketplace IDs for the marketplaces you are registered to sell in.
// http://docs.developer.amazonservices.com/en_US/reports/Reports_RequestReport.html
func (f Finances) ListFinancialEventGroups(optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{
		"MaxResultsPerPage", "FinancialEventGroupStartedAfter", "FinancialEventGroupStartedBefore",
	}, optional)
	params := mws.Parameters{
		"Action": "ListFinancialEventGroups",
	}.Merge(op)

	return f.SendRequest(params)
}

func (f Finances) ListFinancialEvents(optional ...mws.Parameters) (*mws.Response, error) {
	op := mws.OptionalParams([]string{
		"MaxResultsPerPage", "AmazonOrderId", "FinancialEventGroupId", "PostedAfter", "PostedBefore",
	}, optional)
	params := mws.Parameters{
		"Action": "ListFinancialEvents",
	}.Merge(op)

	return f.SendRequest(params)
}
