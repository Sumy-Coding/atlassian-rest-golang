package serv

import (
	"bytes"
	"confluence-rest-golang/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type LabelService struct {
}

func (l LabelService) GetPageLabels(url string, tok string, pid string) models.LabelArray {

	reqUrl := fmt.Sprintf("%s/rest/api/content/%s/label", url, pid)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	client := myClient(reqUrl, tok)
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	var cnArray models.LabelArray
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &cnArray)

	return cnArray

}

func (l LabelService) AddLabels(url string, tok string, pid string, labels []string) models.LabelArray {

	log.Println("Adding labels to " + pid)

	lbls := make([]models.Label, 0)

	for _, l := range labels {
		lbObj := models.Label{Prefix: "global", Name: l}
		lbls = append(lbls, lbObj)
	}
	reqUrl := fmt.Sprintf("%s/rest/api/content/%s/label", url, pid)

	lJson, err := json.Marshal(lbls)
	req, err := http.NewRequest("POST", reqUrl, bytes.NewReader(lJson))
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	client := myClient(reqUrl, tok)
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	var label models.LabelArray
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &label)

	return label

}

//func (l LabelService) CopyLabels(url string, tok string, pid string, tgt string) models.LabelArray {
//
//	log.Printf("Copying %s page labels to %s", pid, tgt)
//
//	labels := l.GetPageLabels(url, tok, pid).Results
//
//	l.AddLabel(url, tok, tgt, labels)
//
//	var label models.LabelArray
//	bts, err := ioutil.ReadAll(resp.Body)
//	err = json.Unmarshal(bts, &label)
//
//	return label
//
//}

//DELETE /rest/api/content/{id}/label/{label}

func (l LabelService) DeleteLabel(url string, tok string, pid string, lid string) string {

	reqUrl := fmt.Sprintf("%s/rest/api/content/%s/label/%s", url, pid, lid)
	req, err := http.NewRequest("DELETE", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	client := myClient(reqUrl, tok)
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	var rStr string
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &rStr)

	return rStr

}
