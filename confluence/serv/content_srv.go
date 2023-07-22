package serv

type ContentService interface {
	GetContent(host string, token string, id string)
	CreateContent(host string, token string, id string, body string)
}
