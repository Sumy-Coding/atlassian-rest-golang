package main

import (
	atlas "confluence-rest-golang"
	"confluence-rest-golang/jira"
	"fmt"
	"runtime"
	"time"
)

func main() {
	start := time.Now()

	locUrl := "http://localhost:9500"

	//ranServ := serv2.RandService{}
	tokService := atlas.TokenService{}

	//locUser := os.Getenv("CONF_LOC_U")
	//locPass, _ := os.LookupEnv("CONF_LOC_P")
	dcUser := "admin"
	dcPass := "admin"
	lToken := tokService.GetToken(dcUser, dcPass)

	//randomString := ranServ.RandomString(10)
	runtime.GOMAXPROCS(50)

	is := jira.IssueService{}
	//issue := is.GetIssue(locUrl, lToken, "AAA-3")
	//fmt.Println(issue)

	created := is.CreateIssue(locUrl, lToken, jira.CreateIssue{Fields: jira.CreateFields{
		Project:     jira.CreateIssueProject{Id: "10000"},
		Summary:     "Test",
		Issuetype:   jira.CIIssuetype{Id: "10006"},
		Assignee:    jira.Assignee{Name: "admin"},
		Reporter:    jira.Reporter{Name: "admin"},
		Priority:    jira.Priority{Id: "3"},
		Labels:      []string{"aa", "bb"},
		Description: "Go test",
		Duedate:     "2023-04-10",
	}})
	fmt.Println(created)

	// == END
	fmt.Printf("Operations took '%f' secs", time.Now().Sub(start).Seconds())
}
