package main

import (
	"atlas-rest-golang/confluence/serv"
	"atlas-rest-golang/srv"
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	start := time.Now()

	// args
	argsWithoutProg := os.Args[1:]
	//args := flag.Args()
	//flag.Parse()

	var instanceType string
	var action string
	var id string

	for a := 0; a < len(argsWithoutProg); a++ {
		if argsWithoutProg[a] == "--type" {
			instanceType = argsWithoutProg[a+1]
		}
		if argsWithoutProg[a] == "--action" {
			action = argsWithoutProg[a+1]
		}
		if argsWithoutProg[a] == "--id" {
			id = argsWithoutProg[a+1]
		}
	}

	fmt.Println(instanceType)
	fmt.Println(action)
	fmt.Println(id)

	//locJiraUrl := "http://localhost:9500"

	tokService := token.TokenService{}
	anmaWiki := os.Getenv("ANMA_URL")
	anmaTok := os.Getenv("ANMA_PASS")

	//locUser := os.Getenv("CONF_LOC_U")
	//locPass, _ := os.LookupEnv("CONF_LOC_P")
	anmaToken := tokService.GetToken(os.Getenv("ANMA_USER"), anmaTok)

	//randomString := ranServ.RandomString(10)
	runtime.GOMAXPROCS(50)

	//is := jira.IssueService{}
	//issue := is.GetIssue(locUrl, lToken, "AAA-3")
	//fmt.Println(issue)

	//created := is.CreateIssue(locUrl, lToken, jira.CreateIssue{Fields: jira.CreateFields{
	//	Project:     jira.CreateIssueProject{Id: "10000"},
	//	Summary:     "Test",
	//	Issuetype:   jira.CIIssuetype{Id: "10006"},
	//	Assignee:    jira.Assignee{Name: "admin"},
	//	Reporter:    jira.Reporter{Name: "admin"},
	//	Priority:    jira.Priority{Id: "3"},
	//	Labels:      []string{"aa", "bb"},
	//	Description: "Go test",
	//	Duedate:     "2023-04-10",
	//}})
	//fmt.Println(created)

	ps := serv.PageService{}
	page := ps.GetPage(anmaWiki, anmaToken, id)
	fmt.Println(page)

	// == END
	fmt.Printf("Operations took '%f' secs", time.Now().Sub(start).Seconds())
}
