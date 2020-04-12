package model

type ImageData struct {
	ImageKey string `json:"image_key"`
}

type Image struct {
	CommonField
	Data *ImageData
}

