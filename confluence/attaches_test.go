package confluence

import (
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestDownloadAttachment(t *testing.T) {
	client := &http.Client{}
	resp, err := client.Get(os.Getenv("FILE_URL"))
	if err != nil {
		log.Printf("Error getting file: %s", err)
	}

	//file, err := os.ReadFile(url)
	//if err != nil {
	//	return
	//}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	log.Println(string(data))
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile("file.txt", data, 0644)
	if err != nil {
		log.Println(err)
	}

}
