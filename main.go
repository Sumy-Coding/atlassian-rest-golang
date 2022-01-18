package main

import (
	"confluence-rest-golang/serv"
	"fmt"
	"os"
	"time"
)

func main() {
	now := time.Now()
	locUrl := "http://localhost:7150"
	pageServ := serv.PageService{}
	//labServ := serv.LabelService{}
	ranServ := serv.RandService{}
	tokService := serv.TokenService{}

	//ncName, _ := os.LookupEnv("NC_ID")
	//ncPass, b := os.LookupEnv("NC_PASS")
	locUser, _ := os.LookupEnv("CONF_LOC_U")
	locPass, _ := os.LookupEnv("CONF_LOC_P")
	lToken := tokService.GetToken(locUser, locPass)

	ranServ.RandomString(10)

	// == Get Page
	//fmt.Println(pageServ.GetPage(locUrl, lToken, "1474565"))

	// === Children
	//child := pageServ.GetSpacePages(locUrl, lToken, "DEV")
	//fmt.Println(child)

	// === GET descendants
	//for _, page := range pageServ.GetDescendants(locUrl, lToken, "65603", 200).Results {
	//	fmt.Println(page)
	//}

	// === Create Pages
	//pageServ.CreatePage(locUrl, lToken, "DEV", "2719745", "t 333", ranServ.RandomString(10))
	// == several
	//for i := 1; i < 15; i++ {
	//	bod := ranServ.RandomString()
	//	pageServ.CreatePage(locUrl, lToken, "DEV", "1769480", fmt.Sprintf("dev 1-2-1--%d", i), bod)
	//}

	// == Edit Page
	//fmt.Println(pageServ.UpdatePage(locUrl, lToken, "2719745", "lorem", "lorem lorem lorem"))

	// == COPY page
	pageServ.CopyPage(locUrl, lToken, "2719747", "2719745")

	// ==== ADD attach

	//for _, att := range pageServ.GetPageAttaches(locUrl, "1474565").Results {
	//	fmt.Println(att.Id)
	//	pageServ.AddFileAsAttach(locUrl, "1474561", att.Id)
	//}

	// ==== Get attaches
	//fmt.Println(pageServ.GetPageAttaches(locUrl, 1474565))

	// === GET attach
	//fmt.Println(pageServ.GetAttach(locUrl, 1671169))

	// === Download attach
	//pageServ.DownloadAttach(locUrl, 1671171)

	// === GET labels
	//fmt.Println(labServ.GetPageLabels(locUrl, lToken, "2719745"))

	// == END
	fmt.Printf("Script took %d secs", time.Now().Sub(now).Seconds())
}
