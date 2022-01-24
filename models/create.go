package models

type CreatePage struct {
	Type            string          `json:"type"`
	Title           string          `json:"title"`
	Ancestors       []Ancestor      `json:"ancestors"`
	CreatePageSpace CreatePageSpace `json:"space"`
	Body            Body            `json:"body"`
}

type EditPage struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Type    string   `json:"type"`
	Space   Space    `json:"space"`
	Body    Body     `json:"body"`
	Version VersionE `json:"version"`
}

type CreateLabel struct {
	Prefix string `json:"prefix"`
	Name   string `json:"name"`
}
