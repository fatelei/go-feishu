package model

type messageAPIData struct {
	MessageID	string	`json:"message_id"`
}

type MessageAPIResponse struct {
	CommonField,
	Data	*messageAPIData
}
