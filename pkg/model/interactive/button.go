package model

type ButtonModule struct {
	Tag      string            `json:"tag"`
	Text     *TextModule        `json:"text"`
	Url      string            `json:"url,omitempty"`
	MultiUrl *UriModule         `json:"multi_url,omitempty"`
	Type     string            `json:"type,omitempty"`
	Value    map[string]string `json:"value,omitempty"`
	Confirm  *ConfirmModule     `json:"confirm,omitempty"`
}


func (p ButtonModule) SetValue(key string, value string) {
	p.Value[key] = value
}
