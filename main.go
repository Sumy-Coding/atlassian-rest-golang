package main

import (
	"confluence-rest-golang/serv"
	"fmt"
)

func main() {
	CONF_URL := "http://localhost:7150"
	pageServ := serv.PageService{}
	page := pageServ.GetPage(CONF_URL, 65603)
	fmt.Println(page)
}
