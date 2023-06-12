package jira

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type IssueService struct{}

func GetClient() *http.Client {
	return &http.Client{Timeout: 10 * time.Second}
}

func (is IssueService) GetIssue(url string, token string, key string) Issue {
	var issue Issue
	reqUrl := fmt.Sprintf("%s/rest/api/2/issue/%s", url, key)
	client := GetClient()
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))

	if err != nil {
		fmt.Printf("Error when creating request. Err: %s", err)
	}
	response, err := client.Do(req)
	if err != nil {
		return Issue{}
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Issue{}
	}
	json.Unmarshal(data, &issue)

	return issue
}

func (is IssueService) CreateIssue(url string, token string, data *CreateIssue) CreatedIssue {
	fmt.Println(data)
	log.Printf(">> Creating issue in project %s", data.Fields.Project.Id)
	reqUrl := fmt.Sprintf("%s/rest/api/2/issue/", url)
	client := GetClient()
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error when creating issue for host %s. %v", url, err)
	}
	req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewReader(jsonData))

	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Printf("Error when creating request. Err: %s", err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	bData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	var issue CreatedIssue
	err2 := json.Unmarshal(bData, &issue)
	if err != nil {
		log.Println(err2)
	}

	return issue
}

func (is IssueService) UpdateField(url string, token string, id string, fieldId string, data interface{}) string {
	reqUrl := fmt.Sprintf("%s/rest/api/2/issue/%s", url, id)
	client := GetClient()
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
	if err != nil {
		fmt.Printf("Error when creating request. Err: %s", err)
	}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()

	bData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}
	var issue Issue
	err2 := json.Unmarshal(bData, &issue)
	if err != nil {
		log.Println(err2)
	}

	// update
	reqUrl2 := fmt.Sprintf("%s/rest/api/2/issue/%s", url, id)
	req2, err := http.NewRequest(http.MethodPut, reqUrl2, strings.NewReader(
		fmt.Sprintf("{\"fields\":{ %s: %s }}", fieldId, data)))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
	if err != nil {
		fmt.Printf("Error when creating request. Err: %s", err)
	}
	resp, err := client.Do(req2)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	return "field updated"
}
