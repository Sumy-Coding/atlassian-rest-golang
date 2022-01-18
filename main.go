package main

import (
	"confluence-rest-golang/serv"
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	locUrl := "http://localhost:7150"
	pageServ := serv.PageService{}
	//child := pageServ.GetSpacePages(CONF_URL, "DEV")
	//fmt.Println(child)
	//for i := 3; i < 30; i++ {
	//	pageServ.CreatePage(locUrl, "DEV", 2523138, fmt.Sprintf("dev 34-%d", i), "lorem ipsum dolor ...")
	//}
	pageServ.AddFileAsAttach(locUrl, "mod.go", 2523138)
	fmt.Printf("Script took %d secs", time.Now().Sub(now).Seconds())
}
