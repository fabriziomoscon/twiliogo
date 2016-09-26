package twiliogo

import (
	"encoding/json"
	"net/url"
)

type Account struct {
	Sid             string   `json:"sid"`
	DateCreated     string   `json:"date_created"`
	DateUpdated     string   `json:"date_updated"`
	FriendlyName    string   `json:"friendly_name"`
	Type            string   `json:"type"`
	Status          string   `json:"status"`
	AuthToken       string   `json:"auth_token"`
	Uri             string   `json:"uri"`
	SubresourceUris []string `json:"subresources_uris"`
	OwnerAccountSid string   `json:"owner_account_sid"`
}

func buildAccountsUri() string {
	return ROOT + "/" + VERSION + "/Accounts/"
}

func GetAccount(client Client, sid string) (*Account, error) {
	res, err := client.get(url.Values{}, buildAccountsUri()+sid+".json")

	if err != nil {
		return nil, err
	}

	account := new(Account)
	err = json.Unmarshal(res, account)

	return account, err
}

func NewAccount(client Client, friendlyName string) (*Account, error) {
	var account *Account

	params := url.Values{}
	params.Set("FriendlyName", friendlyName)

	res, err := client.post(params, buildAccountsUri()+".json")
	if err != nil {
		return nil, err
	}

	account = new(Account)
	err = json.Unmarshal(res, account)

	return account, err
}
