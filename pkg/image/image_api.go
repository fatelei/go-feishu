package image

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatelei/go-feishu/pkg/model"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type ImageAPI struct {
	endpoint string
	accessToken *model.AccessToken
}

func NewImageAPI(endPoint string, accessToken *model.AccessToken) *ImageAPI {
	return &ImageAPI{endpoint:endPoint, accessToken: accessToken}
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
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", p.accessToken.Token))
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