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

func (as AttachService) DownloadAttachmentById(url string, token string, aid string) models.Attachment {
	// todo
	return models.Attachment{}
}

func (as AttachService) DownloadAttachments(url string, token string, pid string) []models.Attachment {
	log.Printf("Downloading page '%s' attachments", pid)
	reqUrl := fmt.Sprintf("%s/rest/api/content/%s/child/attachment", url, pid)

	req, err := http.NewRequest(http.MethodGet, reqUrl, nil)
	req.Header.Add("Authorization", "Basic "+token)
	req.Header.Add("Accept", "application/json")

	client := myClient()
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error when performing GET request: %s", err)
	}
	defer resp.Body.Close()

	var attaches models.AttachmentsResult
	bts, err := io.ReadAll(resp.Body)
	err = json.Unmarshal(bts, &attaches)

	downloaded := make([]models.Attachment, 0)

	for _, att := range attaches.Results {
		dLink := fmt.Sprintf("%s%s", url, att.Links.Download)
		response, err := client.Get(dLink)
		if err != nil {
			log.Printf("Error when getting attachment via GET HTTP request. Err: %s", err)
		}
		//defer response.Body.Close()

		var attach models.Attachment
		bts, err := io.ReadAll(response.Body)
		err = json.Unmarshal(bts, &attach)
		downloaded = append(downloaded, attach)
	}

	return downloaded
}
