package models

type CreateSpace struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	//Type     string   `json:"type"`
}

type CreatePageSpace struct {
	Key string `json:"key"`
}
