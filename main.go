package main

import (
	"confluence-rest-golang/serv"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	//locUrl := "http://localhost:7150"
	bhtUrl := os.Getenv("BHT_URL")
	pageServ := serv.PageService{}
	//labServ := serv.LabelService{}
	//spaceServ := serv.SpaceService{}
	ranServ := serv.RandService{}
	tokService := serv.TokenService{}

	//ncName, _ := os.LookupEnv("NC_ID")
	//ncPass, b := os.LookupEnv("NC_PASS")
	//locUser := os.Getenv("CONF_LOC_U")
	//locPass, _ := os.LookupEnv("CONF_LOC_P")
	bhUser := os.Getenv("BHT_USER")
	bhPass, _ := os.LookupEnv("BHT_TOKEN")
	//bhUser := os.Getenv("BHT_USER")
	//bhPass, _ := os.LookupEnv("BHT_TOKEN")
	//lToken := tokService.GetToken(locUser, locPass)
	bhToken := tokService.GetToken(bhUser, bhPass)

	ranServ.RandomString(10)

	// == Get Page
	fmt.Println(pageServ.GetPage(bhtUrl, bhToken, "468287489"))

	// === Children
	//child := pageServ.GetSpacePages(bhtUrl, bhToken, "DEMO")
	//fmt.Println(child)

	// === GET descendants
	//for _, page := range pageServ.GetDescendants(bhtUrl, bhToken, "75890693", 200).Results {
	//	fmt.Println(page)
	//}

	// === Create Pages
	//pageServ.CreateContent(bhtUrl, bhToken, "page", "TEST", "467369985", "HHHHH", ranServ.RandomString(15))
	// == several
	//for i := 1; i < 15; i++ {
	//	bod := ranServ.RandomString(15)
	//	pageServ.CreateContent(bhtUrl, bhToken, "page", "TEST", "467369985", fmt.Sprintf("RST - %d", i), bod)
	//}

	// COMPLEX HIERARCHY
	//var count int
	//for _, page := range pageServ.GetChildren(locUrl, lToken, "66221").Results {
	//	for i := 1; i < 40; i++ {
	//		bod := ranServ.RandomString(50)
	//		pageServ.CreateContent(locUrl, lToken, "page", "DEV3", page.Id, fmt.Sprintf("%s - %d", page.Title, i), bod)
	//		count += i
	//	}
	//}
	//log.Printf("%d pages created", count)

	// === GET space
	//fmt.Println(spaceServ.GetSpace(locUrl, lToken, "DEV"))

	// === CREATE SPACE
	//fmt.Println(spaceServ.CreateSpace(locUrl, lToken, "DEV2", "DEV2"))

	// == Edit Page
	//fmt.Println(pageServ.UpdatePage(locUrl, lToken, "2719745", "lorem", "lorem lorem lorem"))

	// == COPY page
	//pageServ.CopyPage(locUrl, lToken, "65603", "2326554", true, true, false)

	// == COPY Hierarchy
	//log.Println(pageServ.CopyPageDescs(bhtUrl, bhToken, "", "", "", true, true, false))

	// ==== ADD attach
	//for _, att := range pageServ.GetPageAttaches(locUrl, lToken, "2719753").Results {
	//	fmt.Println(att.Id)
	//	pageServ.CopyAttach(locUrl, lToken, "2981909", att.Id)
	//}

	// ==== Get attaches
	//fmt.Println(pageServ.GetPageAttaches(locUrl, 1474565))

	// === GET attach
	//fmt.Println(pageServ.GetAttach(locUrl, 1671169))

	// === Download attach
	//pageServ.DownloadAttach(locUrl, 1671171)

	// === GET labels
	//fmt.Println(labServ.GetPageLabels(locUrl, lToken, "2719745"))

	// === COPY labels
	//labels := []string{"aaa", "bbb"}
	//labServ.AddLabels(locUrl, lToken, "2555907", labels)

	// === Comments
	//pageServ.AddComment(locUrl, lToken, "2555915", "2555911")

	// == END
	fmt.Printf("Operations took '%f' secs", time.Now().Sub(start).Seconds())
}
