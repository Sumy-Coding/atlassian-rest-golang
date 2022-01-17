package models

type GenericLinks struct {
	Web string
	UI  string
}

type ContentArray struct {
	Results []Content `json:"results"`
	start   int
	limit   int
	size    int
	Links   GenericLinks `json:"_links"`
}
