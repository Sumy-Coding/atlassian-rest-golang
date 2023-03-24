package serv

import (
	"bytes"
	models2 "confluence-rest-golang/confluence/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type SpaceService struct {
}

func (s SpaceService) GetSpace(url string, tok string, key string) models2.Space {
	var client *http.Client = &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	var reqUrl string = fmt.Sprintf("%s/rest/api/space/%s?expand=homepage,metadata.labels", url, key)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	rspb, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		log.Println(err2)
	}
	var space models2.Space

	err = json.Unmarshal(rspb, &space)
	fmt.Println(string(rspb))

	return space

}

func (s SpaceService) CreateSpace(url string, tok string, key string, name string) models2.Space {

	log.Printf("Creating space %s", name)

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	space := models2.CreateSpace{Key: key, Name: name}
	bod, _ := json.Marshal(space)
	reqUrl := fmt.Sprintf("%s/rest/api/space", url)
	req, err := http.NewRequest("POST", reqUrl, bytes.NewReader(bod))
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	rspb, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		log.Println(err2)
	}
	log.Printf("Response is %s", string(rspb))
	var crSPace models2.Space
	err4 := json.Unmarshal(rspb, crSPace)
	if err != nil {
		log.Panicln(err4)
	}

	return crSPace
}
