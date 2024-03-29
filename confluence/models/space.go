package models

type Homepage struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Title  string `json:"title"`
}

type Space struct {
	ID         int      `json:"id"`
	Key        string   `json:"key"`
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Status     string   `json:"status"`
	Homepage   Homepage `json:"homepage"`
	Expandable struct {
		Settings    string `json:"settings"`
		Metadata    string `json:"metadata"`
		Operations  string `json:"operations"`
		LookAndFeel string `json:"lookAndFeel"`
		Identifiers string `json:"identifiers"`
		Permissions string `json:"permissions"`
		Icon        string `json:"icon"`
		Description string `json:"description"`
		Theme       string `json:"theme"`
		History     string `json:"history"`
	} `json:"_expandable"`
	Links struct {
		Webui string `json:"webui"`
		Self  string `json:"self"`
	} `json:"_links"`
}
