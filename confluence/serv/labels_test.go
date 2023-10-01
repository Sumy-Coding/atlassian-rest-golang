package serv

import (
	token2 "atlas-rest-golang/srv"
	"log"
	"testing"
)

func TestGetLabel(t *testing.T) {
	host, user, pass := GetAuthData()
	ls := LabelService{}

	tokService := token2.TokenService{}
	token := tokService.GetToken(user, pass)

	labels := ls.GetPageLabels(host, token, "519276711")

	log.Println(labels)
}
