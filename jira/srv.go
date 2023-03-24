package jira

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
