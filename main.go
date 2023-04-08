package main

import (
	"atlas-rest-golang/confluence/models"
	"atlas-rest-golang/confluence/serv"
	"atlas-rest-golang/jira"
	"atlas-rest-golang/srv"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {
	runtime.GOMAXPROCS(50)
	start := time.Now()

	argsWithoutProg := os.Args[1:]
	//args := flag.Args()
	//flag.Parse()

	var instanceType string
	var action string

	// confluence
	var pageId string
	var pageTitle string
	var spaceKey string
	var parent string
	var body string
	var labels string

	// jira
	var projKey string
	var projId string
	var issueKey string

	//misc
	var file string

	// create issue data
	var summary string
	var description string
	var issTypeId string // 10006
	var dueDate string   // "2023-04-10"
	var issueLabels []string
	var assignee string
	var reporter string
	var priorityName string // normal: id=3

	for a := 0; a < len(argsWithoutProg); a++ {
		if argsWithoutProg[a] == "--type" {
			instanceType = argsWithoutProg[a+1]
		}
		if argsWithoutProg[a] == "--action" {
			action = argsWithoutProg[a+1]
		}

		// confluence
		if argsWithoutProg[a] == "--id" {
			pageId = argsWithoutProg[a+1]
		}
		if argsWithoutProg[a] == "--space" {
			spaceKey = argsWithoutProg[a+1]
		}
		if argsWithoutProg[a] == "--title" {
			pageTitle = argsWithoutProg[a+1]
		}
		if argsWithoutProg[a] == "--parent" {
			parent = argsWithoutProg[a+1]
		}
		if argsWithoutProg[a] == "--body" {
			body = argsWithoutProg[a+1]
		}
		if argsWithoutProg[a] == "--file" {
			file = argsWithoutProg[a+1]
		}
		if argsWithoutProg[a] == "--labels" {
			labels = argsWithoutProg[a+1]
		}

		// jira
		if argsWithoutProg[a] == "--key" {
			projKey = argsWithoutProg[a+1]
		}
		if argsWithoutProg[a] == "--summary" {
			summary = argsWithoutProg[a+1]
		}
		if argsWithoutProg[a] == "--priority" {
			priorityName = argsWithoutProg[a+1]
		}
		if argsWithoutProg[a] == "--descr" {
			description = argsWithoutProg[a+1]
		}
	}

	tokService := token.TokenService{}
	url := os.Getenv("ATLAS_URL")
	if instanceType == "confluence" {
		url += "/wiki"
	}
	pass := os.Getenv("ATLAS_PASS")
	anmaToken := tokService.GetToken(os.Getenv("ATLAS_USER"), pass)

	// Confluence instance
	switch instanceType {
	case "confluence":
		ps := serv.PageService{}
		ss := serv.SpaceService{}
		ls := serv.LabelService{}
		// find space's home page
		if pageTitle == "@home" {
			space := ss.GetSpace(url, anmaToken, spaceKey)
			pageTitle = space.Homepage.Title
		}
		switch action {
		case "getPage":
			if pageId != "" {
				page := ps.GetPage(url, anmaToken, pageId)
				printPage(page)
			} else {
				page := ps.GetPageTitleKey(url, anmaToken, spaceKey, pageTitle)
				printPage(page)
			}
		case "getSpace":
			space := ss.GetSpace(url, anmaToken, spaceKey)
			fmt.Println(space)
		case "createPage":
			created := ps.CreateContent(url, anmaToken, "page", spaceKey, parent, pageTitle, body)
			if labels != "" {
				ls.AddLabels(url, anmaToken, created.Id, strings.Split(labels, ","))
			}
			printPage(created)
		case "addLabel":
			page := ps.GetPage(url, anmaToken, pageId)
			if labels != "" {
				ls.AddLabels(url, anmaToken, page.Id, strings.Split(labels, ","))
			}
			log.Printf("Added labels '%s' to page '%s'", labels, page.Id)
		}

	// Jira instance
	case "jira":
		is := jira.IssueService{}
		ps := jira.ProjectService{}
		switch action {
		case "getIssue":
			issue := is.GetIssue(url, anmaToken, issueKey)
			fmt.Println(issue)
		case "getProject":
			proj := ps.GetProject(url, anmaToken, projKey)
			fmt.Println(proj)
		case "createIssue":
			created := is.CreateIssue(url, anmaToken, jira.CreateIssue{Fields: jira.CreateFields{
				Project:     jira.CreateIssueProject{Id: projId}, //10000
				Summary:     summary,
				Issuetype:   jira.CIIssuetype{Id: issTypeId},
				Assignee:    jira.Assignee{Name: assignee},
				Reporter:    jira.Reporter{Name: reporter},
				Priority:    jira.Priority{Name: priorityName},
				Labels:      issueLabels,
				Description: description,
				Duedate:     dueDate,
			}})
			fmt.Println(created)

		}

	}
	log.Println(file) // todo

	// == END
	fmt.Printf("Operations took '%f' sec", time.Now().Sub(start).Seconds())
}

func printPage(page models.Content) {
	log.Println("============ Content =============")
	log.Printf("\nType: %s\nTitle: %s\nSpace: %s\nBody: %s",
		page.Type, page.Title, page.Space.Name, page.Body.Storage.Value)
}
