package twiliogo

import (
	"encoding/json"
	"net/url"
)

type Address struct {
	Sid          string `json:"sid"`
	AccountSid   string `json:"account_sid"`
	FriendlyName string `json:"friendly_name"`
	CustomerName string `json:"vustomer_name"`
	Street       string `json:"street"`
	City         string `json:"city"`
	Region       string `json:"region"`
	PostalCode   string `json:"postal_code"`
	IsoCountry   string `json:"iso_country"`
	Uri          string `json:"uri"`
}

func GetAddress(client Client, accSid string, addressId string) (*Address, error) {
	res, err := client.get(url.Values{}, buildAddressesUri(accSid)+"/"+addressId+".json")

	if err != nil {
		return nil, err
	}

	addr := new(Address)
	err = json.Unmarshal(res, addr)

	return addr, err
}

func buildAddressesUri(accSid string) string {
	return ROOT + "/" + VERSION + "/Accounts/" + accSid + "/Addresses"
}

func NewAddress(client Client, accSid string, optionals ...Optional) (*Address, error) {
	var addr *Address

	params := url.Values{}
	for _, optional := range optionals {
		param, value := optional.GetParam()
		params.Set(param, value)
	}

	res, err := client.post(params, buildAddressesUri(accSid)+".json")
	if err != nil {
		return nil, err
	}

	addr = new(Address)
	err = json.Unmarshal(res, addr)

	return addr, err
}

func RemoveAddress(client Client, accSid, addSid string) error {
	err := client.delete(buildAddressesUri(accSid) + "/" + addSid + ".json")
	if err != nil {
		return err
	}

	return nil
}
