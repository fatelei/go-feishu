package room

import (
	"encoding/json"
	auth2 "github.com/fatelei/go-feishu/pkg/auth"
	"github.com/fatelei/go-feishu/pkg/model"
	transport2 "github.com/fatelei/go-feishu/pkg/transport"
)

type RoomAPI struct {
	auth      *auth2.Auth
	transport *transport2.Transport
}

func NewRoomAPI(appID string, appSecret string, endPoint string) *RoomAPI {
	auth := auth2.NewAuth(appID, appSecret, endPoint)
	transport := &transport2.Transport{Endpoint:endPoint}
	return &RoomAPI{auth:auth, transport:transport}
}


func (p *RoomAPI) ListChatRoom(pageToken string, pageSize string) (*model.ListRoomResponse, error){
	param := map[string]string{
		"page_size": pageSize,
	}
	if len(pageToken) > 0 {
		param["page_token"] = pageToken
	}
	body, err := p.transport.Get("/open-apis/chat/v4/list", &param, p.auth.GetAccessToken())
	if err != nil {
		return nil, err
	}

	response := &model.ListRoomResponse{}
	if err := json.Unmarshal(body, response); err != nil {
		return nil, err
	}
	return response, nil
}
