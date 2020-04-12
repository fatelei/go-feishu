package main

import (
	"flag"
	"fmt"
	auth2 "github.com/fatelei/go-feishu/pkg/auth"
)

func main() {
	var appID string
	var appSecret string
	flag.StringVar(&appID, "app_id", "", "app id")
	flag.StringVar(&appSecret, "app_secret", "", "app secrect")
	flag.Parse()

	if len(appID) == 0 || len(appSecret) == 0 {
		fmt.Println("app_id & app_secret is required")
		return
	}

	auth := auth2.NewAuth(appID, appSecret, "https://open.feishu.cn")
	fmt.Printf("%s\n", auth.GetAccessToken())
}
