package auth

import (
	"encoding/json"
	"github.com/fatelei/go-feishu/pkg/model"
	transport2 "github.com/fatelei/go-feishu/pkg/transport"
	"time"
)

type Auth struct {
	accessToken *model.AccessToken
	appID string
	appSecret string
	transport *transport2.Transport
}


func NewAuth(appID string, appSecret string, endPoint string) *Auth {
	transport := &transport2.Transport{Endpoint: endPoint}
	return &Auth{
		accessToken: nil,
		appID:       appID,
		appSecret:   appSecret,
		transport:	 transport,
	}
}


func (p *Auth) GetAccessToken() string {
	flag := false
	if p.accessToken == nil {
		flag = true
	} else if time.Now().Unix() - p.accessToken.Expire <= 10 {
		flag = true
	}
	if flag {
		resp, err := p.transport.PostJson("/open-apis/auth/v3/tenant_access_token/internal", p.generateParam())
		if err != nil {
			panic(err)
		} else {
		    tmp := &model.AccessToken{}
			if err := json.Unmarshal(resp, tmp); err != nil {
				panic(err)
			}
			p.accessToken = tmp
		}
	}
	return p.accessToken.Token
}

func (p *Auth) generateParam() map[string]interface{} {
	body := make(map[string]interface{}, 0)
	body["app_id"] = p.appID
	body["app_secret"] = p.appSecret
	return body
}
