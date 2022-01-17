package serv

import (
	"confluence-rest-golang/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type PageService struct {
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
	return nil
}

func myClient(url string, token string) *http.Client {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	return client
}

func (s PageService) GetPage(url string, id int64) models.Content {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	reqUrl := fmt.Sprintf("%s/rest/api/content/%d", url, id)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	var content models.Content
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &content)

	return content

}

func (s PageService) GetChildren(url string, id int64) models.ContentArray {

	reqUrl := fmt.Sprintf("%s/rest/api/content/%d/child/page", url, id)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
	client := myClient(reqUrl, basicAuth("admin", "admin"))
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	var cnArray models.ContentArray
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &cnArray)

	return cnArray

}

func (s PageService) GetDescendants(url string, id int64) models.ContentArray {

	reqUrl := fmt.Sprintf("%s/rest/api/content/search?cql=ancestor=%d&limit=300", url, id)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
	client := myClient(reqUrl, basicAuth("admin", "admin"))
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	var cnArray models.ContentArray
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &cnArray)

	return cnArray

}

func (s PageService) PageContains(url string, id int64, find string) bool {
	value := s.GetPage(url, id).Body.Storage.Value
	return strings.Contains(value, find)
}

func (s PageService) GetSpacePages(url string, key string) models.ContentArray {

	reqUrl := fmt.Sprintf("%s/rest/api/content?type=page&spaceKey=%s&limit=300", url, key)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
	client := myClient(reqUrl, basicAuth("admin", "admin"))
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	var cnArray models.ContentArray
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &cnArray)

	return cnArray

}

//getSpacePagesByLabel() {                                 // todo

func (s PageService) GetSpaceBlogs(url string, key string) models.ContentArray {

	reqUrl := fmt.Sprintf("%s/rest/api/content?type=blogpost&spaceKey=%s&limit=300", url, key) //limit=300
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
	client := myClient(reqUrl, basicAuth("admin", "admin"))
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	var cnArray models.ContentArray
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &cnArray)

	return cnArray

}
