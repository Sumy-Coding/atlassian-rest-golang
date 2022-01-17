package models

type GenericLinks struct {
	Web string
	UI  string
}

type ContentArray struct {
	Results []Content
	start   int
	limit   int
	size    int
	_links  GenericLinks
}
