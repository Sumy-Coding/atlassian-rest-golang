package serv

import (
	"bytes"
	"confluence-rest-golang/models"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
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

//func basicAuth(username, password string) string {
//	auth := username + ":" + password
//	return base64.StdEncoding.EncodeToString([]byte(auth))
//}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	locUser, _ := os.LookupEnv("CONF_LOC_U")
	locPass, _ := os.LookupEnv("CONF_LOC_P")
	tokServ := TokenService{}
	tok := tokServ.GetToken(locUser, locPass)
	req.Header.Add("Authorization", "Basic "+tok)
	return nil
}

func myClient(url string, token string) *http.Client {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	return client
}

func (s PageService) GetPage(url string, tok string, id string) models.Content {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	expand := "expand=body.storage,history,version"

	reqUrl := fmt.Sprintf("%s/rest/api/content/%s?%s", url, id, expand)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	var content models.Content
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &content)

	return content

}

func (s PageService) GetChildren(url string, tok string, id string) models.ContentArray {

	expand := "expand=body.storage,history,version"
	reqUrl := fmt.Sprintf("%s/rest/api/content/%s/child/page?%s", url, id, expand)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	client := myClient(reqUrl, tok)
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	var cnArray models.ContentArray
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &cnArray)

	return cnArray

}

func (s PageService) GetDescendants(url string, tok string, id string) models.ContentArray {

	//expand := "?expand=body.storage,history,version"

	reqUrl := fmt.Sprintf("%s/rest/api/content/search?cql=ancestor=%s&limit=300", url, id)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	client := myClient(reqUrl, tok)
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	var cnArray models.ContentArray
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &cnArray)

	return cnArray

}

func (s PageService) PageContains(url string, tok string, id string, find string) bool {
	value := s.GetPage(url, tok, id).Body.Storage.Value
	return strings.Contains(value, find)
}

func (s PageService) GetSpacePages(url string, tok string, key string) models.ContentArray {

	expand := "expand=body.storage,history,version"
	reqUrl := fmt.Sprintf("%s/rest/api/content?type=page&spaceKey=%s&%s&limit=300", url, key, expand)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	client := myClient(reqUrl, tok)
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

func (s PageService) GetSpaceBlogs(url string, tok string, key string) models.ContentArray {

	reqUrl := fmt.Sprintf("%s/rest/api/content?type=blogpost&spaceKey=%s&limit=300", url, key) //limit=300
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	client := myClient(reqUrl, tok)
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	var cnArray models.ContentArray
	bytes, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bytes, &cnArray)

	return cnArray

}

func (s PageService) DeletePageLabels(url string, tok string, id string, labels []string) string {
	for _, lab := range labels {
		reqUrl := fmt.Sprintf("%s/rest/api/content/%d/label/%s", url, id, lab) //limit=300
		req, err := http.NewRequest("DELETE", reqUrl, nil)
		req.Header.Add("Authorization", "Basic "+tok)
		client := myClient(reqUrl, tok)
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

func (s PageService) DeletePage(url string, tok string, id string) models.Content {
	reqUrl := fmt.Sprintf("%s/rest/api/content/%d", url, id) //limit=300
	req, err := http.NewRequest("DELETE", reqUrl, nil)
	req.Header.Add("Authorization", "Basic "+tok)
	client := myClient(reqUrl, tok)
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
func (s PageService) ScrollTemplates(url string, tok string, key string) []string {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	reqUrl := fmt.Sprintf("%s/plugins/servlet/scroll-office/api/templates?spaceKey=%s", url, key)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
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

func (s PageService) CreateSpace(url string, tok string, key string) []string {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	reqUrl := fmt.Sprintf("%s/rest/api/space", url)
	req, err := http.NewRequest("POST", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
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
func (s PageService) CreatePage(url string, tok string, key string, parent string, title string, bd string) models.Content {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	reqUrl := fmt.Sprintf("%s/rest/api/content", url)
	ancts := []models.Ancestor{{Id: parent}} // parent
	cntb := models.CreatePage{
		//Id:    "",
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
	req.Header.Add("Authorization", "Basic "+tok)
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

func (s PageService) GetPageAttaches(url string, tok string, pid string) models.ContentArray {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	expand := "expand=body.storage,history,version"
	reqUrl := fmt.Sprintf("%s/rest/api/content/%s/child/attachment?%s", url, pid, expand)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
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

func (s PageService) GetAttach(url string, tok string, aid string) models.Content {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	expand := "expand=body.storage,history,version"
	reqUrl := fmt.Sprintf("%s/rest/api/content/%s?%s", url, aid, expand)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
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

func (s PageService) DownloadAttach(url string, tok string, atId string) string {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	attach := s.GetAttach(url, tok, atId)

	fileDir, _ := os.Getwd()
	fName := attach.Title

	// download link
	dwLink := attach.Links.Base + attach.Links.Download

	//btb := &bytes.Buffer{} // byte buffer
	//mime.ParseMediaType()

	//reader := multipart.NewReader(btb, "")
	//reader.NextPart()

	r1, _ := http.NewRequest("GET", dwLink, nil)
	r1.Header.Add("Authorization", "Basic "+tok)
	//r1.Header.Add("Content-Type", writer.FormDataContentType())
	r1.Header.Add("X-Atlassian-Token", "nocheck")

	resp, err := client.Do(r1)
	fmt.Println(resp, err) // log response

	bts, err := ioutil.ReadAll(resp.Body)
	err = os.WriteFile(fName, bts, fs.ModeAppend)

	filePath := path.Join(fileDir, fName) // file path
	fmt.Println("File path is " + filePath)

	return filePath
}

func (s PageService) AddFileAsAttach(url string, tok string, pid string, atId string) string {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	reqUrl := fmt.Sprintf("%s/rest/api/content/%s/child/attachment", url, pid)

	//req, err := http.NewRequest("POST", reqUrl, bytes.NewReader(mrsCtn))
	// =========

	// get attach
	attach := s.GetAttach(url, tok, atId)

	fileDir, _ := os.Getwd()
	fName := attach.Title
	filePath := path.Join(fileDir, fName) // equals atFilePath
	fmt.Println("File path is " + filePath)

	atFilePath := s.DownloadAttach(url, tok, atId) // attach.Id = nil

	//os.WriteFile(fName)
	//ioutil.WriteFile(fName)
	// save attach
	fl, _ := os.Open(atFilePath) // atFilePath
	defer fl.Close()

	btb := &bytes.Buffer{} // byte buffer
	writer := multipart.NewWriter(btb)
	part, _ := writer.CreateFormFile("file", filepath.Base(fl.Name()))
	io.Copy(part, fl)
	clErr := writer.Close()
	if clErr != nil {
		log.Panicln(clErr)
	}

	r, _ := http.NewRequest("POST", reqUrl, btb)
	cntType := writer.FormDataContentType()
	fmt.Println("Content type is " + cntType)

	r.Header.Add("Authorization", "Basic "+tok)
	r.Header.Add("Content-Type", cntType)
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
	// delete attach
	defer os.Remove(atFilePath)

	return "attachment added " + string(bts)

}

func (s PageService) GetComment(url string, tok string, cid string) models.Content {
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	expand := "expand=body.storage,history,version"
	reqUrl := fmt.Sprintf("%s/rest/api/content/%s?%s", url, cid, expand)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
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
