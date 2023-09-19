package serv

import (
	token2 "atlas-rest-golang/srv"
	"fmt"
	"log"
	"testing"
)

const (
	host = "http://localhost:8510"
	user = "admin"
	pass = "admin"
)

func TestGetPage(t *testing.T) {
	ps := PageService{}
	tokService := token2.TokenService{}
	token := tokService.GetToken(user, pass)
	page := ps.GetPage(host, token, "98390")
	log.Println(page)
}

func TestGetPages(t *testing.T) {
	ps := PageService{}
	tokService := token2.TokenService{}
	token := tokService.GetToken(user, pass)
	pages := ps.GetChildren(host, token, "98388")
	for _, page := range pages.Results {
		log.Println(page)
	}
}

func TestCreatePage(t *testing.T) {
	ps := PageService{}
	tokService := token2.TokenService{}
	token := tokService.GetToken(user, pass)
	createdPage := ps.CreateContent(host, token, "page", "test", "98390", "Go test 1", "lorem ipsum...")
	log.Println(createdPage)
}

func TestCreatePages(t *testing.T) {
	ps := PageService{}
	tokService := token2.TokenService{}
	token := tokService.GetToken(user, pass)
	rootId := "98388"

	for i := 0; i < 30; i++ {
		createdPage := ps.CreateContent(host,
			token,
			"page",
			"test",
			rootId,
			fmt.Sprintf("Go test %d", i),
			"lorem ipsum...")
		log.Println(createdPage)
	}
}
