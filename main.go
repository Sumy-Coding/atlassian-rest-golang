package main

import (
	serv2 "confluence-rest-golang/confluence/serv"
	"confluence-rest-golang/jira"
	"fmt"
	"runtime"
	"time"
)

func main() {
	start := time.Now()

	locUrl := "http://localhost:9500"

	//ranServ := serv2.RandService{}
	tokService := serv2.TokenService{}

	//locUser := os.Getenv("CONF_LOC_U")
	//locPass, _ := os.LookupEnv("CONF_LOC_P")
	dcUser := "admin"
	dcPass := "admin"
	lToken := tokService.GetToken(dcUser, dcPass)

	//randomString := ranServ.RandomString(10)
	runtime.GOMAXPROCS(50)

	is := jira.IssueService{}
	issue := is.GetIssue(locUrl, lToken, "AAA-3")
	fmt.Println(issue)

	// == END
	fmt.Printf("Operations took '%f' secs", time.Now().Sub(start).Seconds())
}
