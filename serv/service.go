package serv

import (
	"bytes"
	"confluence-rest-golang/models"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
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

func (s PageService) DeletePageLabels(url string, id int64, labels []string) string {
	for _, lab := range labels {
		reqUrl := fmt.Sprintf("%s/rest/api/content/%d/label/%s", url, id, lab) //limit=300
		req, err := http.NewRequest("DELETE", reqUrl, nil)
		req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
		client := myClient(reqUrl, basicAuth("admin", "admin"))
		resp, err := client.Do(req)
		if err != nil {
			log.Panicln(err)
		}
		fmt.Println(resp)
		//bts, err := ioutil.ReadAll(resp.Body)
		//err = json.Unmarshal(bytes, &cnArray)
	}

	return "labels deleted "
}

func (s PageService) DeletePage(url string, id int64) models.Content {
	reqUrl := fmt.Sprintf("%s/rest/api/content/%d", url, id) //limit=300
	req, err := http.NewRequest("DELETE", reqUrl, nil)
	req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
	client := myClient(reqUrl, basicAuth("admin", "admin"))
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	var cnt models.Content
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &cnt)
	return cnt
}

//"${CONF_URL}/plugins/servlet/scroll-office/api/templates?spaceKey=${spaceKey}"
func (s PageService) ScrollTemplates(url string, key string) []string {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	reqUrl := fmt.Sprintf("%s/plugins/servlet/scroll-office/api/templates?spaceKey=%s", url, key)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	tms := []string{}
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &tms)

	return tms

}

// todo

func (s PageService) CreateSpace(url string, key string) []string {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	reqUrl := fmt.Sprintf("%s/plugins/servlet/scroll-office/api/templates?spaceKey=%s", url, key)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	rspb, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Println(err2)
	}
	fmt.Println("Response is " + string(rspb))
	return nil

}

//createPage(CONF_URL, TOKEN, space, parentId, title, body)
func (s PageService) CreatePage(url string, key string, parent int64, title string, bd string) models.Content {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	reqUrl := fmt.Sprintf("%s/rest/api/content", url)
	ancts := []models.Ancestor{{Id: parent}} // parent
	cntb := models.Content{
		Type:  "page",
		Title: title,
		Space: models.Space{
			Key: key,
		}, Body: models.Body{
			Storage: models.Storage{
				Representation: "storage", Value: bd},
		},
		Ancestors: ancts,
	}
	mrsCtn, err2 := json.Marshal(cntb)
	if err2 != nil {
		log.Panicln(err2)
	}
	req, err := http.NewRequest("POST", reqUrl, bytes.NewReader(mrsCtn))
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	var content models.Content
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &content)
	fmt.Println(string(bts))

	return content

}

func (s PageService) GetPageAttaches(url string, pid int64) models.ContentArray {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	reqUrl := fmt.Sprintf("%s/rest/api/content/%d/child/attachment", url, pid)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	var carr models.ContentArray
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &carr)
	fmt.Println(string(bts))

	return carr

}

func (s PageService) GetAttach(url string, aid int64) models.Content {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	reqUrl := fmt.Sprintf("%s/rest/api/content/%d", url, aid)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}

	var content models.Content
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &content)
	fmt.Println(string(bts))

	return content

}

func (s PageService) AddFileAsAttach(url string, fn string, atId int64) string {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	reqUrl := fmt.Sprintf("%s/rest/api/content/%d/child/attachment", url, atId)

	//req, err := http.NewRequest("POST", reqUrl, bytes.NewReader(mrsCtn))
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	//req.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
	//req.Header.Add("Content-Type", "application/json")
	//ioutil.WriteFile()
	// =========

	//attach._links.base + attach._links.download

	fileDir, _ := os.Getwd()
	fileName := fn
	filePath := path.Join(fileDir, fileName)

	fl, _ := os.Open(filePath)
	defer fl.Close()

	btb := &bytes.Buffer{} // byte buffer
	writer := multipart.NewWriter(btb)
	part, _ := writer.CreateFormFile("file", filepath.Base(fl.Name()))
	io.Copy(part, fl)
	writer.Close()

	r, _ := http.NewRequest("POST", reqUrl, btb)
	r.Header.Add("Authorization", "Basic "+basicAuth("admin", "admin"))
	r.Header.Add("Content-Type", writer.FormDataContentType())
	r.Header.Add("X-Atlassian-Token", "nocheck")
	//client := &http.Client{}
	resp, err := client.Do(r)
	//if err != nil {
	//	log.Panicln(err)
	//}
	/// ====
	if err != nil {
		log.Panicln(err)
	}

	//var content models.Content
	bts, err := ioutil.ReadAll(resp.Body)
	//err = json.Unmarshal(bts, &content)
	fmt.Println(string(bts))

	return "attachment added " + string(bts)

}
