package serv

import (
	"confluence-rest-golang/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
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
