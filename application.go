package twiliogo

import (
	"encoding/json"
	"net/url"
)

type Application struct {
	Sid                   string `json:"sid"`
	DateCreated           string `json:"date_created"`
	DateUpdated           string `json:"date_updated"`
	FriendlyName          string `json:"friendly_name"`
	AccountSid            string `json:"account_sid"`
	ApiVersion            string `json:"api_version"`
	VoiceUrl              string `json:"voice_url"`
	VoiceMethod           string `json:"voice_method"`
	VoiceFallbackUrl      string `json:"voice_fallback_url"`
	VoiceFallbackMethod   string `json:"voice_fallback_method"`
	StatusCallback        string `json:"status_callback"`
	StatusCallbackMethod  string `json:"status_callback_method"`
	VoiceCallerIdLookup   bool   `json:"voice_caller_id_lookup"`
	SmsUrl                string `json:"sms_url"`
	SmsMethod             string `json:"sms_method"`
	SmsFallbackUrl        string `json:"sms_fallback_url"`
	SmsFallbackMethod     string `json:"sms_fallback_method"`
	SmsStatusCallback     string `json:"sms_status_callback"`
	MessageStatusCallback string `json:"message_status_callback"`
	Uri                   string `json:"uri"`
}

func buildApplicationsUri(accSid string) string {
	return ROOT + "/" + VERSION + "/Accounts/" + accSid + "/Applications"
}

func GetApplication(client Client, accSid string, appId string) (*Application, error) {
	res, err := client.get(url.Values{}, buildApplicationsUri(accSid)+"/"+appId+".json")

	if err != nil {
		return nil, err
	}

	app := new(Application)
	err = json.Unmarshal(res, app)

	return app, err
}

func NewApplication(client Client, accSid string, optionals ...Optional) (*Application, error) {
	var app *Application

	params := url.Values{}
	for _, optional := range optionals {
		param, value := optional.GetParam()
		params.Set(param, value)
	}

	res, err := client.post(params, buildApplicationsUri(accSid)+".json")
	if err != nil {
		return nil, err
	}

	app = new(Application)
	err = json.Unmarshal(res, app)

	return app, err
}
