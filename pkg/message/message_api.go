package message

import (
	"encoding/json"
	"fmt"
	"github.com/fatelei/go-feishu/pkg/model"
	model2 "github.com/fatelei/go-feishu/pkg/model/interactive"
	transport2 "github.com/fatelei/go-feishu/pkg/transport"
)

type MessageAPI struct {
	transport *transport2.Transport
}


func NewMessageAPI(endPoint string) *MessageAPI {
	transport := &transport2.Transport{Endpoint:endPoint}
	return &MessageAPI{transport:transport}
}


func (p *MessageAPI) SendImage(
	chatId string, title string, imgKey string, action *model2.ActionModule, accessToken string) (*model.MessageAPIResponse, error) {
	imageModule := model2.ImageModule{
		Tag:    "img",
		ImgKey: imgKey,
		Alt:    &model2.TextModule{
			Tag:     "plain_text",
			Content: title,
		},
		Title:  &model2.TextModule{
			Tag:     "plain_text",
			Content: title,
		},
	}
	elements := make([]interface{}, 0)
	elements = append(elements, imageModule)
	if action != nil {
		elements = append(elements, *action)
	}
	messageCard := model2.MessageCard{
		Elements: elements,
	}
	body := make(map[string]interface{})
	body["open_chat_id"] = chatId
	body["msg_type"] = "interactive"
	body["card"] = messageCard
	if byte, err := json.Marshal(&body); err == nil {
		fmt.Printf("%s\n", string(byte))
	}
	resp, err := p.transport.Post("/open-apis/message/v4/send", &body, accessToken)
	if err != nil {
		return nil, err
	}
	data := &model.MessageAPIResponse{}
	if err := json.Unmarshal(resp, data); err != nil {
		return nil, err
	}
	return data, nil
}

func (p *MessageAPI) SendInteractiveCard(
	chatId string, title string, elements []interface{}, accessToken string) (*model.MessageAPIResponse, error) {
	messageCard := model2.MessageCard{
		Header: &model2.CardHeader{
			Title: &model2.TextModule{
				Tag:     "plain_text",
				Content: title,
			},
		},
		Elements: elements,
	}
	body := make(map[string]interface{})
	body["open_chat_id"] = chatId
	body["msg_type"] = "interactive"
	body["card"] = messageCard
	if byte, err := json.Marshal(&body); err == nil {
		fmt.Printf("%s\n", string(byte))
	}
	resp, err := p.transport.Post("/open-apis/message/v4/send", &body, accessToken)
	if err != nil {
		return nil, err
	}
	data := &model.MessageAPIResponse{}
	if err := json.Unmarshal(resp, data); err != nil {
		return nil, err
	}
	return data, nil
}