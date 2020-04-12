package model

type Room struct {
	Avatar string `json:"avatar"`
	Description string `json:"description"`
	ChatID string `json:"chat_id"`
	Name string `json:"name"`
	OwnerOpenID string `json:"owner_open_id"`
	OwnerUserID string `json:"owner_user_id"`
}


type RoomData struct {
	Groups []*Room `json:"groups"`
	PageToken string `json:"page_token"`
	HasMore bool `json:"has_more"`
}


type ListRoomResponse struct {
	CommonField
	Data *RoomData
}
