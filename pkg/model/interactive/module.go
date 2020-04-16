package model

type TextModule struct {
	Tag 	string 	`json:"tag"`
	Content string 	`json:"content"`
	Lines 	int64 	`json:"lines,omitempty"`
}

type ContentModule struct {
	Tag	string `json:"tag"`
	Text *TextModule `json:"text"`
	Fields []*Field `json:"fields,omitempty"`
}

type Field struct {
	IsShort bool        `json:"is_short"`
	Text 	*TextModule `json:"text"`
}

type ImageModule struct {
	Tag		string      `json:"tag"`
	ImgKey	string       `json:"img_key"`
	Alt		*TextModule `json:"alt"`
	Title   *TextModule    `json:"title,omitempty"`
}

type HrModule struct {
	Tag		string 		`json:"tag"`
}

type UriModule struct {
	Url				string		`json:"url"`
	AndroidUrl		string 		`json:"android_url"`
	IosUrl			string		`json:"ios_url"`
	PcUrl			string		`json:"pc_url"`
}

type ConfirmModule struct {
	Title TextModule `json:"title"`
	Text  TextModule `json:"text"`
}

type CardConfig struct {
	WideScreenMode	bool `json:"wide_screen_mode"`
}

type MessageCard struct {
	Config *CardConfig     `json:"config,omitempty"`
	Header *TextModule     `json:"header,omitempty"`
	Elements []interface{} `json:"elements"`
}

type ActionModule struct {
	Tag		string						`json:"tag"`
	Actions	[]Interactive				`json:"actions"`
}
