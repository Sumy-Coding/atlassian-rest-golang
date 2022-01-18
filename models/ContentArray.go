package models

type GenericLinks struct {
	Webui      string
	Download   string
	Collection string
	Base       string
	Context    string
	Self       string
}

type ContentArray struct {
	Results []Content `json:"results"`
	start   int
	limit   int
	size    int
	Links   GenericLinks `json:"_links"`
}
