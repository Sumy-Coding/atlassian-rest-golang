package jira

import (
	token "atlas-rest-golang/srv"
	"fmt"
	"os"
	"testing"
)

func TestIssuesService(t *testing.T) {
	tokService := token.TokenService{}
	url := os.Getenv("ATLAS_URL")

	pass := os.Getenv("ATLAS_PASS")
	anmaToken := tokService.GetToken(os.Getenv("ATLAS_USER"), pass)
	is := IssueService{}
	iss := is.GetIssue(url, anmaToken, "AAA-2")

	fmt.Println(iss)
}

func TestCreateIssue(t *testing.T) {
	tokService := token.TokenService{}
	url := os.Getenv("ATLAS_URL")
	pass := os.Getenv("ATLAS_PASS")
	aToken := tokService.GetToken(os.Getenv("ATLAS_USER"), pass)
	is := IssueService{}
	created := is.CreateIssue(url, aToken, &CreateIssue{
		Fields: CreateFields{
			Project:     CreateIssueProject{Id: "10000"},
			Summary:     "summary",
			Issuetype:   CIIssuetype{Id: "10006"},
			Assignee:    Assignee{Name: "admin"},
			Reporter:    Reporter{Name: "admin"},
			Labels:      []string{"aaa"},
			Description: "description",
		}})

	fmt.Println(created)
}

func TestCreateIssues(t *testing.T) {
	tokService := token.TokenService{}
	url := os.Getenv("ATLAS_URL")
	pass := os.Getenv("ATLAS_PASS")
	aToken := tokService.GetToken(os.Getenv("ATLAS_USER"), pass)
	is := IssueService{}
	for i := 0; i < 20; i++ {
		created := is.CreateIssue(url, aToken, &CreateIssue{
			Fields: CreateFields{
				Project:     CreateIssueProject{Id: "10000"},
				Summary:     fmt.Sprintf("summary %d", i),
				Issuetype:   CIIssuetype{Id: "10006"},
				Assignee:    Assignee{Name: "admin"},
				Reporter:    Reporter{Name: "admin"},
				Labels:      []string{"aaa"},
				Description: "description",
			}})

		fmt.Println(created)
	}
}
