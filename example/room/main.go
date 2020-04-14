package main

import (
	"flag"
	"fmt"
	auth2 "github.com/fatelei/go-feishu/pkg/auth"
	room2 "github.com/fatelei/go-feishu/pkg/room"
)

func main() {
	var appID string
	var appSecret string
	var endpoint string
	flag.StringVar(&appID, "app_id", "", "app id")
	flag.StringVar(&appSecret, "app_secret", "", "app secrect")
	flag.StringVar(&endpoint, "endpoint", "https://open.feishu.cn", "endpoint")
	flag.Parse()

	if len(appID) == 0 || len(appSecret) == 0 {
		fmt.Println("app_id & app_secret is required")
		return
	}

	auth := auth2.NewAuth(appID, appSecret, endpoint)
	accessToken := auth.GetAccessToken()
	room := room2.NewRoomAPI(endpoint, accessToken.Token)
	resp, _ := room.ListChatRoom("", "100")
	fmt.Printf("%+v\n", resp.Data)
	if resp.Data != nil {
		for _, group := range resp.Data.Groups {
			fmt.Printf("%+v\n", group)
		}
	}
}
