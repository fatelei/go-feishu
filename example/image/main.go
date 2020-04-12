package main

import (
	"flag"
	"fmt"
	"github.com/fatelei/go-feishu/pkg/image"
	"os"
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

	imageApi := image.NewImageAPI(appID, appSecret, "https://open.feishu.cn")
	resp, _ := imageApi.UploadFromUri("http://i.imgur.com/Dz2r9lk.jpg")
	fmt.Printf("%v\n", resp)

	dir, _ := os.Getwd()
	resp, _ = imageApi.UploadFromFile(fmt.Sprintf("%s/example/image/test.jpg", dir))
	fmt.Printf("%v\n", resp)
}
