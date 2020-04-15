package auth

import (
	"encoding/json"
	"github.com/fatelei/go-feishu/pkg/model"
	transport2 "github.com/fatelei/go-feishu/pkg/transport"
)

type Auth struct {
	appID string
	appSecret string
	transport *transport2.Transport
}


func NewAuth(appID string, appSecret string, endPoint string) *Auth {
	transport := &transport2.Transport{Endpoint: endPoint}
	return &Auth{
		appID:       appID,
		appSecret:   appSecret,
		transport:	 transport,
	}
}


func (p *Auth) GetAccessToken() *model.AccessToken {
	resp, err := p.transport.PostJson("/open-apis/auth/v3/tenant_access_token/internal", p.generateParam())
	if err != nil {
			panic(err)
		} else {
	    tmp := &model.AccessToken{}
		if err := json.Unmarshal(resp, tmp); err != nil {
				panic(err)
			}
		return tmp
	}
	return nil
}

func (p *Auth) generateParam() map[string]interface{} {
	body := make(map[string]interface{}, 0)
	body["app_id"] = p.appID
	body["app_secret"] = p.appSecret
	return body
}
