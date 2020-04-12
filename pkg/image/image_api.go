package image

import (
	"bytes"
	"encoding/json"
	"fmt"
	auth2 "github.com/fatelei/go-feishu/pkg/auth"
	"github.com/fatelei/go-feishu/pkg/model"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type ImageAPI struct {
	auth *auth2.Auth
	endpoint string
}

func NewImageAPI(appID string, appSecret string, endPoint string) *ImageAPI {
	auth := auth2.NewAuth(appID, appSecret, endPoint)
	return &ImageAPI{auth:auth, endpoint:endPoint}
}


func (p *ImageAPI) UploadFromFile(filePath string) (*model.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()
	return p.do(file)
}


func (p *ImageAPI) UploadFromUri(imageUri string) (*model.Image, error) {
	resp, err := http.Get(imageUri)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	return p.do(resp.Body)
}


func (p *ImageAPI) do(binary io.Reader) (*model.Image, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", "image")
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, binary)
	writer.WriteField("image_type", "message")
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/open-apis/image/v4/put", p.endpoint)
	request, err := http.NewRequest("POST", url, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.auth.GetAccessToken()))
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	image := &model.Image{}
	fmt.Printf("%s\n", string(bytes))
	if err := json.Unmarshal(bytes, image); err != nil {
		return nil, err
	}
	return image, nil
}