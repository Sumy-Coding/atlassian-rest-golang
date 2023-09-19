package main

import (
	"atlas-rest-golang/confluence/serv"
	"encoding/base64"
	"fmt"
	"log"
	"testing"
)

const (
	confUrl = "http://localhost:8510"
	user    = "admin"
)

func TestCreatePage(t *testing.T) {
	pageService := serv.PageService{}
	tok := base64.StdEncoding.EncodeToString([]byte("admin:admin"))
	pageService.CreateContent(
		confUrl,
		tok,
		"page",
		"DEV",
		"98371",
		"GO page 1",
		"asdasdsad")
}

func TestGetPage(t *testing.T) {
	//
	pageService := serv.PageService{}
	tok := base64.StdEncoding.EncodeToString([]byte("admin:admin"))
	page := pageService.GetPage(confUrl, tok, "98371")

	log.Println(page.Body.Storage.Value)

}

func TestCreateNpages(t *testing.T) {
	pageService := serv.PageService{}
	tok := base64.StdEncoding.EncodeToString([]byte("admin:admin"))
	for i := 2; i < 20; i++ {
		pageService.CreateContent(
			confUrl,
			tok,
			"page",
			"DEV",
			"98371",
			fmt.Sprintf("GO page %d", i),
			"asdasdsad")
	}
}
