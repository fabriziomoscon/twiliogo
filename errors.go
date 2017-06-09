package twiliogo

type Error struct {
	Description string
}

func (e Error) Error() string {
	return e.Description
}

type TwilioError struct {
	Status   int    `json:"status"`
	Message  string `json:"message"`
	Code     int    `json:"code"`
	MoreInfo string `json:"more_info"`
}

func (e TwilioError) Error() string {
	return e.Message
}
