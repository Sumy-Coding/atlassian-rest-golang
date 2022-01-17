package main

import (
	"confluence-rest-golang/serv"
	"fmt"
)

func main() {
	CONF_URL := "http://localhost:7150"
	pageServ := serv.PageService{}
	child := pageServ.GetSpacePages(CONF_URL, "DEV")
	fmt.Println(child)
}
