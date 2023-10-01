package serv

import (
	token2 "atlas-rest-golang/srv"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
)

const (
	host = "http://localhost:8510"
	user = "admin"
	pass = "admin"
)

func GetAuthData() (string, string, string) {
	confUrl := os.Getenv("ATLAS_URL")
	user := os.Getenv("ATLAS_USER")
	pass := os.Getenv("ATLAS_PASS")
	return confUrl, user, pass
}

func TestGetPage(t *testing.T) {
	confUrl, user, pass := GetAuthData()
	ps := PageService{}
	tokService := token2.TokenService{}
	token := tokService.GetToken(user, pass)
	page := ps.GetPage(confUrl, token, "519308225")
	log.Println(page)
}

func TestGetPages(t *testing.T) {
	confUrl, user, pass := GetAuthData()
	ps := PageService{}
	tokService := token2.TokenService{}
	token := tokService.GetToken(user, pass)
	pages := ps.GetChildren(confUrl, token, "98388")
	for _, page := range pages.Results {
		log.Println(page)
	}
}

func TestCreatePage(t *testing.T) {
	host, user, pass := GetAuthData()
	ps := PageService{}
	tokService := token2.TokenService{}
	token := tokService.GetToken(user, pass)
	createdPage := ps.CreateContent(
		host, token, "page",
		"TEST", "519276317", "Go test 1", "lorem ipsum...")
	log.Println(createdPage)
}

func TestCreatePageAsync(t *testing.T) {
	wg := sync.WaitGroup{}
	host, user, pass := GetAuthData()
	ps := PageService{}
	tokService := token2.TokenService{}
	token := tokService.GetToken(user, pass)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup, num int) {
			wg.Add(1)
			createdPage := ps.CreateContent(
				host, token, "page",
				"TEST", "519276317", fmt.Sprintf("Go test %d", i), "lorem ipsum...")
			wg.Done()
			log.Println(createdPage)
		}(&wg, i)
		wg.Wait()
	}
}

func TestCreatePages(t *testing.T) {
	host, user, pass := GetAuthData()
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
