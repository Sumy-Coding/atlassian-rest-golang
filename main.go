package main

import (
	"confluence-rest-golang/serv"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {
	start := time.Now()

	locUrl := "http://localhost:7201"
	//bhtUrl := os.Getenv("BHT_URL")
	pageServ := serv.PageService{}
	//labServ := serv.LabelService{}
	spaceServ := serv.SpaceService{}
	ranServ := serv.RandService{}
	tokService := serv.TokenService{}

	locUser := os.Getenv("CONF_LOC_U")
	locPass, _ := os.LookupEnv("CONF_LOC_P")
	//bhUser := os.Getenv("BHT_USER")
	//bhPass, _ := os.LookupEnv("BHT_TOKEN")
	lToken := tokService.GetToken(locUser, locPass)

	ranServ.RandomString(10)
	runtime.GOMAXPROCS(50)

	//msrv := grpca.MyServer{}
	//msrv.InitServer()

	// 										== Get Page
	//fmt.Println(pageServ.GetPage(bhtUrl, bhToken, "468287489"))

	// 										=== Children
	//child := pageServ.GetSpacePages(bhtUrl, bhToken, "DEMO")
	//fmt.Println(child)

	// 										=== GET descendants
	//for _, page := range pageServ.GetDescendants(bhtUrl, bhToken, "75890693", 200).Results {
	//	fmt.Println(page)
	//}

	// 										=== Create Pages
	//space := spaceServ.GetSpace(locUrl, lToken, "test46")
	//pageServ.CreateContent(locUrl, lToken, "page", space.Key, space.Homepage.Id, "HHHHH", ranServ.RandomString(100))

	// 									=== ASYNC == several
	//var waitG sync.WaitGroup
	//for i := 1; i <= 100; i++ {
	//	waitG.Add(1)
	//	//bod := ranServ.RandomString(2)
	//	bod := "lorem asd asd das  d ds"
	//	go func(count int) {
	//		pageServ.CreateContentAsync(&waitG, locUrl, lToken, "page", "GOGO",
	//			"2785364",
	//			fmt.Sprintf("GOGO 100-1 - %d", count),
	//			bod)
	//	}(i)
	//}
	//waitG.Wait()

	// 										==== COMPLEX HIERARCHY
	var count int
	for i := 40; i <= 50; i++ {
		space := spaceServ.GetSpace(locUrl, lToken, fmt.Sprintf("test%d", i))
		homePage := space.Homepage
		for i := 1; i <= 100; i++ {
			bod := ranServ.RandomString(100)
			pageServ.CreateContent(locUrl, lToken, "page", space.Key, homePage.Id, fmt.Sprintf("%s - %d", homePage.Title, i), bod)
		}
		count += i
	}

	log.Printf("%d pages created", count)

	// === GET space
	//fmt.Println(spaceServ.GetSpace(locUrl, lToken, "DEV"))

	// === CREATE SPACE
	//fmt.Println(spaceServ.CreateSpace(locUrl, lToken, "DEV2", "DEV2"))

	//for i := 3; i < 20; i++ {
	//	fmt.Println(spaceServ.CreateSpace(locUrl, lToken, fmt.Sprintf("DEV%d", i), fmt.Sprintf("DEV%d", i)))
	//}

	// Operations took '124.458869' secs
	//for a := 3; a <= 20; a++ {
	//	//wg.Add(1)
	//	go func() {
	//		key := fmt.Sprintf("DEV%d", a)
	//		sp := spaceServ.GetSpace(locUrl, lToken, key)
	//		for i := 40; i < 45; i++ {
	//			bod := ranServ.RandomString(15)
	//			pageServ.CreateContentAsync(locUrl, lToken, "page",
	//				sp.Key,
	//				sp.Homepage.Id,
	//				fmt.Sprintf("RST - %d", i), bod)
	//		}
	//	}()

	//fmt.Println(spaceServ.CreateSpace(locUrl, lToken, fmt.Sprintf("DEV%d", a), fmt.Sprintf("DEV%d", a)))
	//}

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
