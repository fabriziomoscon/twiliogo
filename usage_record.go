package twiliogo

import (
	"encoding/json"
	"net/url"
)

// UsageRecords contains a list of requested UsageRecord's and metadata
type UsageRecords struct {
	FirstPageUri    string        `json:"first_page_uri"`
	End             int           `json:"end"`
	PreviousPageUri string        `json:"previous_page_uri"`
	Uri             string        `json:"uri"`
	PageSize        int           `json:"page_size"`
	Start           int           `json:"start"`
	UsageRecords    []UsageRecord `json:"usage_records"`
}

// UsageRecord contains all data for a Twilio Usage Record
type UsageRecord struct {
	Category    string `json:"category"`
	Description string `json:"description"`
	AccountSid  string `json:"account_sid"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Count       string `json:"count"`
	CountUnit   string `json:"count_unit"`
	Usage       string `json:"usage"`
	UsageUnit   string `json:"usage_unit"`
	Price       string `json:"price,omitempty"`
	PriceUnit   string `json:"price_unit"`
	ApiVersion  string `json:"api_version"`
	Uri         string `json:"uri"`
}

type UsageFilter struct {
	Category  string
	StartDate string
	EndDate   string
}

func GetUsageNoSubresources(client Client, filter *UsageFilter) (*UsageRecords, error) {
	return getUsage(client, "", filter)
}

func GetUsageDaily(client Client, filter *UsageFilter) (*UsageRecords, error) {
	return getUsage(client, "Daily", filter)
}

func GetUsageMonthly(client Client, filter *UsageFilter) (*UsageRecords, error) {
	return getUsage(client, "Monthly", filter)
}

func GetUsageYearly(client Client, filter *UsageFilter) (*UsageRecords, error) {
	return getUsage(client, "Yearly", filter)
}

func GetUsageAllTime(client Client, filter *UsageFilter) (*UsageRecords, error) {
	return getUsage(client, "AllTime", filter)
}

func GetUsageToday(client Client, filter *UsageFilter) (*UsageRecords, error) {
	return getUsage(client, "Today", filter)
}

func GetUsageYesterday(client Client, filter *UsageFilter) (*UsageRecords, error) {
	return getUsage(client, "Yesterday", filter)
}

func GetUsageThisMonth(client Client, filter *UsageFilter) (*UsageRecords, error) {
	return getUsage(client, "ThisMonth", filter)
}

func GetUsageLastMonth(client Client, filter *UsageFilter) (*UsageRecords, error) {
	return getUsage(client, "LastMonth", filter)
}

func getUsage(client Client, subresource string, filter *UsageFilter) (*UsageRecords, error) {
	usageRecordsUrl := "/Usage/Records"
	if subresource != "" {
		usageRecordsUrl = usageRecordsUrl + "/" + subresource
	}

	params := url.Values{}
	if filter != nil {
		if filter.Category != "" {
			params.Set("Category", filter.Category)
		}
		if filter.StartDate != "" {
			params.Set("StartDate", filter.StartDate)
		}
		if filter.EndDate != "" {
			params.Set("EndDate", filter.EndDate)
		}
	}

	res, err := client.get(params, usageRecordsUrl+".json")
	if err != nil {
		return nil, err
	}

	usageRecords := new(UsageRecords)
	err = json.Unmarshal(res, usageRecords)

	return usageRecords, err
}
