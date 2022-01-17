package main

import (
	"confluence-rest-golang/serv"
)

func main() {
	CONF_URL := "http://localhost:7150"
	pageServ := serv.PageService{}
	//child := pageServ.GetSpacePages(CONF_URL, "DEV")
	//fmt.Println(child)
	pageServ.CreatePage(CONF_URL, "DEV", 2523138, "dev 34-2", "asdasdasd")
}
