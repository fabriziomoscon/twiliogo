package twiliogo

type Capabilities struct {
	Voice bool `json:"voice"`
	SMS   bool `json:"SMS"`
	MMS   bool `json:"MMS"`
	FAX   bool `json:"fax"`
}
