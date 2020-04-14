package room

import (
	"encoding/json"
	"github.com/fatelei/go-feishu/pkg/model"
	transport2 "github.com/fatelei/go-feishu/pkg/transport"
)

type RoomAPI struct {
	transport *transport2.Transport
	accessToken *model.AccessToken
}

func NewRoomAPI(endPoint string, accessToken *model.AccessToken) *RoomAPI {
	transport := &transport2.Transport{Endpoint:endPoint}
	return &RoomAPI{accessToken: accessToken, transport:transport}
}


func (p *RoomAPI) ListChatRoom(pageToken string, pageSize string) (*model.ListRoomResponse, error){
	param := map[string]string{
		"page_size": pageSize,
	}
	if len(pageToken) > 0 {
		param["page_token"] = pageToken
	}
	body, err := p.transport.Get("/open-apis/chat/v4/list", &param, p.accessToken.Token)
	if err != nil {
		return nil, err
	}

	response := &model.ListRoomResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		return nil, err
	}
	return response, nil
}
