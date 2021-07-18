package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	auth2 "github.com/fatelei/go-feishu/pkg/auth"
	"github.com/fatelei/go-feishu/pkg/image"
	"io/ioutil"
	"os"
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
	imageApi := image.NewImageAPI(endpoint)
	resp, _ := imageApi.UploadFromUri("http://i.imgur.com/Dz2r9lk.jpg", accessToken.Token)
	fmt.Printf("%v\n", resp)

	dir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s/example/image/test.jpg", dir)
	resp, _ = imageApi.UploadFromFile(filePath, accessToken.Token)
	fmt.Printf("%v\n", resp)

	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	src, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	b64data := base64.StdEncoding.EncodeToString(src)
	resp, _ = imageApi.UploadFromB64Encode(b64data, accessToken.Token)
	fmt.Printf("%v\n", resp)
}
