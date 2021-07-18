package image

import (
	"bytes"
	"encoding/base64"
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
}

func NewImageAPI(endPoint string) *ImageAPI {
	return &ImageAPI{endpoint:endPoint}
}


func (p *ImageAPI) UploadFromFile(filePath string, accessToken string) (*model.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()
	return p.do(file, accessToken)
}

func (p *ImageAPI) UploadFromB64Encode(data string, accessToken string) (*model.Image, error) {
	unbased, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	r := bytes.NewReader(unbased)
	return p.do(r, accessToken)
}


func (p *ImageAPI) UploadFromUri(imageUri string, accessToken string) (*model.Image, error) {
	resp, err := http.Get(imageUri)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	return p.do(resp.Body, accessToken)
}


func (p *ImageAPI) do(binary io.Reader, accessToken string) (*model.Image, error) {
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
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
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