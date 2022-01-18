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

	// === Children
	//child := pageServ.GetSpacePages(CONF_URL, "DEV")
	//fmt.Println(child)

	// === Create Pages
	//for i := 3; i < 30; i++ {
	//	pageServ.CreatePage(locUrl, "DEV", 2523138, fmt.Sprintf("dev 34-%d", i), "lorem ipsum dolor ...")
	//}

	// ==== ADD attach

	//for _, att := range pageServ.GetPageAttaches(locUrl, 1474565).Results {
	//	fmt.Println(att.Id)
	//	//pageServ.AddFileAsAttach(locUrl, 1474561, att.Id)
	//}

	// ==== Get attaches
	fmt.Println(pageServ.GetPageAttaches(locUrl, 1474565))

	// === GET attach
	//fmt.Println(pageServ.GetAttach(locUrl, 1671169))

	// === Download attach
	//pageServ.DownloadAttach(locUrl, 1671171)

	// == END
	fmt.Printf("Script took %d secs", time.Now().Sub(now).Seconds())
}
