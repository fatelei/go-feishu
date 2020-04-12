package main

import (
	"flag"
	"fmt"
	message2 "github.com/fatelei/go-feishu/pkg/message"
	"github.com/fatelei/go-feishu/pkg/model/interactive"
)

func main() {
	var appID string
	var appSecret string
	var chatID string
	var imgKey string
	flag.StringVar(&appID, "app_id", "", "app id")
	flag.StringVar(&appSecret, "app_secret", "", "app secrect")
	flag.StringVar(&chatID, "chat_id", "", "chat id")
	flag.StringVar(&imgKey, "img_key", "", "img key")
	flag.Parse()

	if len(appID) == 0 || len(appSecret) == 0 {
		fmt.Println("app_id & app_secret is required")
		return
	}

	message := message2.NewMessageAPI(appID, appSecret, "https://open.feishu.cn")
	button := model.ButtonModule{
		Tag:   "button",
		Text:  &model.TextModule{Tag: "plain_text", Content: "测试"},
		Value: make(map[string]string),
	}
	button.SetValue("message_id", "1")
	actionModule := &model.ActionModule{
		Tag:     "action",
		Actions: []model.Interactive{button},
	}
	message.SendImage(chatID, "测试", imgKey, actionModule)
}
