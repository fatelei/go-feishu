package model

type AccessToken struct {
	Token 	string		`json:"tenant_access_token"`
	Expire 	int64		`json:"expire"`
}
