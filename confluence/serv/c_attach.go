package serv

import (
	"atlas-rest-golang/confluence/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type AttachService struct{}

func (as AttachService) GetPageAttachments(url string, tok string, pid string) models.AttachmentsResult {
	log.Printf("Getting %s page attachments", pid)
	reqUrl := fmt.Sprintf("%s/rest/api/content/%s/child/attachment", url, pid)
	req, err := http.NewRequest("GET", reqUrl, nil)
	req.Header.Add("Authorization", "Basic "+tok)
	client := myClient()
	resp, err := client.Do(req)
	if err != nil {
		log.Panicf("Erroe when performing GET request: %s", err)
	}
	defer resp.Body.Close()
	var attachments models.AttachmentsResult
	bts, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &attachments)

	return attachments
}

func (as AttachService) AddAttachment(url string, tok string, pid string, attach string) models.Attachment {
	log.Printf("Adding attachment %s to page %s ", attach, pid)
	reqUrl := fmt.Sprintf("%s/rest/api/content/%s/child/attachment", url, pid)

	file, err := os.Open(attach)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("multipart/form-data", attach)
	bw, err := io.Copy(part, file)
	log.Printf("Bytes written: %d", bw)

	if err != nil {
		log.Printf("Error when copying bytes data from file: %s", err)
	}
	defer writer.Close()

	req, err := http.NewRequest("POST", reqUrl, body)
	req.Header.Add("Authorization", "Basic "+tok)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("X-Atlassian-Token", "nocheck")

	client := myClient()
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error when performing request: %s", err)
	}
	defer resp.Body.Close()

	var added models.Attachment
	bts, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &added)

	return added

}

func createForm(form map[string]string) (string, io.Reader, error) {
	body := new(bytes.Buffer)
	mp := multipart.NewWriter(body)
	defer mp.Close()
	for key, val := range form {
		if strings.HasPrefix(val, "@") {
			val = val[1:]
			file, err := os.Open(val)
			if err != nil {
				return "", nil, err
			}
			defer file.Close()
			part, err := mp.CreateFormFile(key, val)
			if err != nil {
				return "", nil, err
			}
			io.Copy(part, file)
		} else {
			mp.WriteField(key, val)
		}
	}
	return mp.FormDataContentType(), body, nil
}
