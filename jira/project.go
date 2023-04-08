package jira

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type ProjectService struct{}

func (ps ProjectService) GetProjects(url string, token string) []Project {
	var projects []Project
	reqUrl := fmt.Sprintf("%s/rest/api/3/project/", url)
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
	if err != nil {
		fmt.Printf("Error when creating request. Err: %s", err)
	}
	response, err := GetClient().Do(req)
	if err != nil {
		log.Printf("Error when GETting projects: %s", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error when closing request body: %s", err)
		}
	}(response.Body)
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error when parsing request body: %s", err)
	}
	err3 := json.Unmarshal(data, &projects)
	if err3 != nil {
		log.Printf("Error when unmarshalling data from request: %s", err)
	}
	return projects
}

func (ps ProjectService) GetProject(url string, token string, key string) Project {
	var project Project
	reqUrl := fmt.Sprintf("%s/rest/api/3/project/%s", url, key)
	log.Printf("The request URL is %s", reqUrl)
	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", token))
	if err != nil {
		fmt.Printf("Error when creating request. Err: %s", err)
	}
	response, err := GetClient().Do(req)
	if err != nil {
		log.Printf("Error when getting ptoject %s. Error: %s", key, err)
		return Project{}
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error when closing request body: %s", err)
		}
	}(response.Body)
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("Error when parsing request body: %v", err)
	}
	err3 := json.Unmarshal(data, &project)
	if err3 != nil {
		log.Printf("Error when unmarshalling data from request: %s", err)
	}

	return project
}

func httpClient() *http.Client {
	return &http.Client{Timeout: 20 * time.Second}
}
