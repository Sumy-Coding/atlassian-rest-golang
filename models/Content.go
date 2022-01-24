package models

type VersionE struct {
	Number int `json:"number"`
}

type Metadata struct {
	Labels LabelArray `json:"labels"`
}

type Space struct {
	Id       int64    `json:"id"`
	Key      string   `json:"key"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	Metadata Metadata `json:"metadata"`
}

type Icon struct {
	Path      string `json:"path"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	IsDefault bool   `json:"isDefault"`
}

type UserDetails struct {
	Business string
	Personal string
}

type User struct {
	Type           string      `json:"type"`
	Username       string      `json:"username"`
	UserKey        string      `json:"userKey"`
	AccountId      string      `json:"accountId"`
	AccountType    string      `json:"accountType"`
	EMail          string      `json:"email"`
	PublicName     string      `json:"publicName"`
	DisplayName    string      `json:"displayName"`
	ProfilePicture Icon        `json:"profilePicture"`
	TimeZone       string      `json:"timeZone"`
	Details        UserDetails `json:"details"`
}

type Version struct {
	By        User   `json:"by"`
	When      string `json:"when"`
	Number    int    `json:"number"`
	MinorEdit bool   `json:"minorEdit"`
}

type ContentHistory struct {
	Latest      bool
	CreatedBy   User
	CreatedDate string  `json:"createdDate"`
	LastUpdated Version `json:"lastUpdated"`
}

//type View struct {
//	Value          string `json:"value"`
//	Representation string `json:"representation"`
//}

type Storage struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}

type Body struct {
	//View    View    `json:"view"`
	Storage Storage `json:"storage"`
}

type Ancestor struct {
	Id string `json:"id"`
}

type Expandable struct {
	//Container Content `json:"container"`
	Operations string `json:"operations"`
}

type Label struct {
	Id     string `json:"id"`
	Prefix string `json:"prefix"`
	Name   string `json:"name"`
}

type LabelArray struct {
	Results []Label `json:"results"`
	start   int
	limit   int
	size    int
	Links   GenericLinks `json:"_links"`
}

type Content struct {
	Id         string         `json:"id"`
	Type       string         `json:"type"`
	Status     string         `json:"status"`
	Title      string         `json:"title"`
	Body       Body           `json:"body"`
	Version    Version        `json:"version"`
	Space      Space          `json:"space"`
	History    ContentHistory `json:"history"`
	Links      GenericLinks   `json:"_links"`
	Ancestors  []Ancestor     `json:"ancestors"`
	Expandable Expandable     `json:"_expandable"`
}
