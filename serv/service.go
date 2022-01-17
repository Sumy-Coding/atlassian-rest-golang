package serv

import (
	"confluence-rest-golang/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type PageService struct {
}

func (s PageService) GetPage(url string, id int64) models.Content {

	reqUrl := fmt.Sprintf("%s/rest/api/content/%d", url, id)
	resp, err := http.Get(reqUrl)
	if err != nil {
		log.Panicln(err)
	}
	var content models.Content
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &content)

	return content

}
