package twiliogo

import (
	"encoding/json"
	"errors"
	"net/url"
)

type AvaliablePhoneNumberWrapper struct {
	URI  string          `json:"uri"`
	Data json.RawMessage `json:"available_phone_numbers"`
}

type AvaliablePhoneNumber struct {
	FriendlyName        string       `json:"friendly_name"`
	PhoneNumber         string       `json:"phone_number"`
	IsoCountry          string       `json:"iso_country"`
	Capabilities        Capabilities `json:"capabilities"`
	AddressRequirements string       `json:"address_requirements"`
	Beta                bool         `json:"beta"`
	Lata                string       `json:"lata"`
	RateCenter          string       `json:"rate_center"`
	Latitude            string       `json:"latitude"`
	Longitude           string       `json:"longitude"`
	Region              string       `json:"region"`
	PostalCode          string       `json:"postal_code"`
}

func GetAvailablePhoneNumbers(client Client, isoCountryCode string, numType string, optionals ...Optional) (*[]AvaliablePhoneNumber, error) {
	if numType != "local" && numType != "mobile" && numType != "tollfree" {
		return nil, errors.New("Invalid number type " + numType)
	}
	var avaliablePhoneNumbers *[]AvaliablePhoneNumber
	avaliablePhoneNumbers = new([]AvaliablePhoneNumber)

	params := url.Values{}
	for _, optional := range optionals {
		param, value := optional.GetParam()
		params.Set(param, value)
	}

	res, err := client.get(params, "/AvailablePhoneNumbers/"+isoCountryCode+"/"+numType+".json")
	if err != nil {
		return nil, err
	}

	var wrapper AvaliablePhoneNumberWrapper
	if err = json.Unmarshal(res, &wrapper); err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(wrapper.Data), avaliablePhoneNumbers)
	return avaliablePhoneNumbers, err
}
