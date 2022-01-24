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

var (
	labServ = LabelService{}
)

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

	expand := "expand=space,body.storage,history,version"

	reqUrl := fmt.Sprintf("%s/rest/api/content/%s?%s", url, id, expand)
	log.Println("GET REQ URL is " + reqUrl)

	req, err := http.NewRequest("GET", reqUrl, nil)
	//defer req.Body.Close()	// close body
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Panicln(err)
		}
	}(resp.Body)

	var content models.Content
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &content)

	return content

}

func (s PageService) GetChildren(url string, tok string, id string) models.ContentArray {

	expand := "expand=space,body.storage,history,version"
	reqUrl := fmt.Sprintf("%s/rest/api/content/%s/child/page?%s", url, id, expand)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//		log.Panicln(err)
	//	}
	//}(req.Body)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	client := myClient(reqUrl, tok)
	resp, err := client.Do(req)
	defer resp.Body.Close() // todo -close body
	if err != nil {
		log.Panicln(err)
	}
	var cnArray models.ContentArray
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &cnArray)

	return cnArray

}

func (s PageService) GetDescendants(url string, tok string, id string, lim int) models.ContentArray {

	//expand := "?expand=body.storage,history,version"

	reqUrl := fmt.Sprintf("%s/rest/api/content/search?cql=ancestor=%s&limit=%d", url, id, lim)
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	client := myClient(reqUrl, tok)
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close() // close request's body

	var cnArray models.ContentArray
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &cnArray)

	return cnArray

}

func (s PageService) PageContains(url string, tok string, id string, find string) bool {
	body := s.GetPage(url, tok, id).Body.Storage.Value
	return strings.Contains(body, find)
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
	defer resp.Body.Close()
	var cnArray models.ContentArray
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &cnArray)

	return cnArray

}

func (p PageService) GetSpacePagesByLabel(url string, tok string, key string, lb string) models.ContentArray { // todo
	//?cql=space+%3D+"DEV"+and+label+%3D+"aa"
	reqUrl := fmt.Sprintf("%s/rest/api/search?cql=space=\"%s\"+and+label+%3D+\"%s\"", url, key, lb)
	req, err := http.NewRequest("GET", reqUrl, nil)
	req.Header.Add("Authorization", "Basic "+tok)
	client := myClient(reqUrl, tok)
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	var cnArray models.ContentArray
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &cnArray)

	return cnArray
}
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
	defer resp.Body.Close()
	var cnArray models.ContentArray
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &cnArray)

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
		defer resp.Body.Close()
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
	defer resp.Body.Close()
	var cnt models.Content
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &cnt)
	return cnt
}

// Scroll url "${CONF_URL}/plugins/servlet/scroll-office/api/templates?spaceKey=${spaceKey}"
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
	defer resp.Body.Close()
	tms := make([]string, 0)
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &tms)

	return tms

}

//createPage(CONF_URL, TOKEN, space, parentId, title, body)
func (s PageService) CreateContent(url string, tok string, ctype string, key string, parent string, title string, bd string) models.Content {

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	reqUrl := fmt.Sprintf("%s/rest/api/content", url)
	ancts := []models.Ancestor{{Id: parent}} // parent
	cntb := models.CreatePage{
		//Id:    "",
		Type:  ctype,
		Title: title,
		CreatePageSpace: models.CreatePageSpace{
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
	defer resp.Body.Close()
	var content models.Content
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &content)
	fmt.Println(string(bts))

	return content
}

func (s PageService) CopyPage(url string, tok string, pid string, tid string,
	copyLabs bool, copyAtt bool, copyCo bool) models.Content {
	// todo - copyLabels, copyComments, copyAttaches

	log.Println("Copying page " + pid)

	orPage := s.GetPage(url, tok, pid)
	//time.Sleep(time.Duration(sservice.ResponseTime) * time.Millisecond) // sleep till getting target parent
	parPage := s.GetPage(url, tok, tid)

	/* reqUrl := fmt.Sprintf("%s/rest/api/createdPage", url)
	ancts := []models.Ancestor{{Id: parent}} // parent
	cntb := models.CreateContent{
		//Id:    "",
		Type:  "page",
		Title: orPage.Title,
		Space: models.Space{
			Key: orPage.Space.Key,
		}, Body: models.Body{
			Storage: models.Storage{
				Representation: "storage", Value: orPage.Body.Storage.Value},
		},
		Ancestors: ancts,
	} */
	var ttl string
	if orPage.Space.Key == parPage.Space.Key {
		ttl = "Copy of " + orPage.Title
	} else {
		ttl = orPage.Title
	}

	var createdPage models.Content
	createdPage = s.CreateContent(url, tok, "page", parPage.Space.Key, tid, ttl, orPage.Body.Storage.Value)

	// copy labels
	if copyLabs {
		lArr := labServ.GetPageLabels(url, tok, pid)
		lbls := make([]string, 0)
		for _, l := range lArr.Results {
			lbls = append(lbls, l.Name)
		}
		labServ.AddLabels(url, tok, createdPage.Id, lbls)
	}
	// attachment
	if copyAtt {
		log.Printf("Copying %s page attachments", pid)
		attaches := s.GetPageAttaches(url, tok, pid).Results // todo - can be more than 100 set currently
		for _, att := range attaches {
			s.CopyAttach(url, tok, createdPage.Id, att.Id)
		}
	}
	// comments
	if copyCo {
		// todo
	}

	return createdPage
}

func (s PageService) CopyPageDescs(url string, tok string, pid string, tgt string, nTitle string,
	copyLabs bool, copyCo bool, copyAtt bool) []models.Content {

	// todo - copyLabels, copyComments, copyAttaches + later 'TargetServer'
	log.Printf("Copying %s page descendants", pid)
	cntList := make([]models.Content, 0)

	//root := s.GetPage(url, tok, pid)
	childs := s.GetChildren(url, tok, pid).Results
	rootCp := s.CopyPage(url, tok, pid, tgt, copyLabs, copyCo, copyAtt)

	log.Printf("ROOT page %s copied as %s", pid, rootCp.Id)

	for _, child := range childs {
		var ttl string
		if nTitle == "" {
			// todo - check current space or different
			if child.Space.Key == rootCp.Space.Key {
				ttl = "Copy of " + child.Title
			} else {
				ttl = child.Title
			}
		} else {
			ttl = nTitle + child.Title
		}
		log.Println("Copying child page " + child.Id + " under " + rootCp.Id)

		// recursion NOT working for GO as in Groovy - use Async ?
		s.CopyPageDescs(url, tok, child.Id, rootCp.Id, ttl, copyLabs, copyCo, copyAtt)
		cpPage := s.CopyPage(url, tok, child.Id, rootCp.Id, copyLabs, copyCo, copyAtt)

		cntList = append(cntList, cpPage)
	}

	return cntList
}

func (s PageService) UpdatePage(url string, tok string, pid string, find string, repl string) models.Content {

	log.Printf("Updating %s page", pid)
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}
	reqUrl := fmt.Sprintf("%s/rest/api/content/%s", url, pid)
	log.Println("Request URL = " + reqUrl)
	log.Println("Edited pageID = " + pid)

	page := s.GetPage(url, tok, pid)
	pBody := page.Body.Storage.Value
	fBody := strings.Replace(pBody, find, repl, -1)
	cntb := models.EditPage{
		Id:    page.Id,
		Title: page.Title,
		Type:  "page",
		Body: models.Body{
			Storage: models.Storage{
				Representation: "storage", Value: fBody},
		},
		Version: models.VersionE{Number: page.Version.Number + 1},
	}
	pageBytes, err2 := json.Marshal(cntb)
	if err2 != nil {
		log.Panicln(err2)
	}
	req, err := http.NewRequest("PUT", reqUrl, bytes.NewReader(pageBytes))
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	var content models.Content
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &content)
	fmt.Println(string(bts))

	return content

}

func (s PageService) GetPageAttaches(url string, tok string, pid string) models.ContentArray {

	log.Printf("Getting %s page attachments", pid)
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	expand := "expand=body.storage,history,version"
	reqUrl := fmt.Sprintf("%s/rest/api/content/%s/child/attachment?limit=100&%s", url, pid, expand) // limit=100
	req, err := http.NewRequest("GET", reqUrl, nil)
	//req.SetBasicAuth("admin", "admin")
	//resp, err := http.Get(reqUrl)
	req.Header.Add("Authorization", "Basic "+tok)
	resp, err := client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	var carr models.ContentArray
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &carr)
	fmt.Println(string(bts))

	return carr

}

func (s PageService) GetAttach(url string, tok string, aid string) models.Content {

	log.Printf("Getting %s attachment", aid)
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
	defer resp.Body.Close()
	var content models.Content
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &content)
	fmt.Println(string(bts))

	return content

}

func (s PageService) DownloadAttach(url string, tok string, atId string) string {

	log.Printf("Getting %s attachment", atId)
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
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()

	bts, rErr := ioutil.ReadAll(resp.Body)
	if rErr != nil {
		log.Panicln(rErr)
	}
	err = os.WriteFile(fName, bts, fs.ModePerm) // mode 0777
	if err != nil {
		log.Panicln(err)
	}
	filePath := path.Join(fileDir, fName) // file path
	fmt.Println("File path is " + filePath)

	return filePath
}

func (s PageService) CopyAttach(url string, tok string, tpid string, atId string) string {

	log.Println("Adding attach: " + atId + " to page: " + tpid)

	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	reqUrl := fmt.Sprintf("%s/rest/api/content/%s/child/attachment", url, tpid)

	//req, err := http.NewRequest("POST", reqUrl, bytes.NewReader(mrsCtn))
	// =========

	// get attach
	attach := s.GetAttach(url, tok, atId)

	fileDir, _ := os.Getwd()
	fName := attach.Title
	filePath := path.Join(fileDir, fName) // equals atFilePath
	fmt.Println("File path of attach is: " + filePath)

	atFilePath := s.DownloadAttach(url, tok, atId) // attach.Id = nil

	//os.WriteFile(fName)
	//ioutil.WriteFile(fName)
	// save attach
	fl, _ := os.Open(atFilePath) // atFilePath / fileDir
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
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	//var content models.Content
	bts, err := ioutil.ReadAll(resp.Body)
	//err = json.Unmarshal(bts, &content)
	fmt.Println(string(bts))
	// delete attach
	defer os.Remove(atFilePath)

	return "attachment added " + string(bts)

}

func (s PageService) GetComment(url string, tok string, cid string) models.Content {

	log.Printf("Getting %s page coment", cid)
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
	defer resp.Body.Close()
	var content models.Content
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &content)
	fmt.Println(string(bts))

	return content
}

func (p PageService) AddComment(url string, tok string, cid string, pid string) models.Content {
	log.Printf("Copying %s comment to %s page", cid, pid)
	client := &http.Client{
		CheckRedirect: redirectPolicyFunc,
	}

	cmm := p.GetComment(url, tok, cid)
	page := p.GetPage(url, tok, pid)

	reqUrl := fmt.Sprintf("%s/rest/api/content", url)
	cntb := models.CreateComment{
		//Id:    "",
		Type:  "comment",
		Title: cmm.Title,
		CreatePageSpace: models.CreatePageSpace{
			Key: cmm.Space.Key,
		}, Body: models.Body{
			Storage: models.Storage{
				Representation: "storage", Value: cmm.Body.Storage.Value},
		},
		Container: page,
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
	defer resp.Body.Close()

	var content models.Content
	bts, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &content)
	fmt.Println(string(bts))

	return content

}

//func (p PageService) CopyComment(url string, tok string, cid string, pid string) models.Content {
//	log.Printf("Copying %s comment to %s page", cid, pid)
//
//	cmm := p.GetComment(url, tok, cid)
//	page := p.GetPage(url, tok, pid)
//
//}
