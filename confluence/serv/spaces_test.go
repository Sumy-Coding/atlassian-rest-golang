package serv

import (
	token2 "atlas-rest-golang/srv"
	"fmt"
	"log"
	"testing"
)

func TestSpaceService_CreateSpace(t *testing.T) {
	ss := SpaceService{}
	tokService := token2.TokenService{}
	token := tokService.GetToken(user, pass)
	createdSpace := ss.CreateSpace(host, token, "test2", "test2")
	log.Println(createdSpace)
}

func TestSpaceService_CreateSpaces(t *testing.T) {
	ss := SpaceService{}
	tokService := token2.TokenService{}
	token := tokService.GetToken(user, pass)
	for i := 1; i < 20; i++ {
		createdSpace := ss.CreateSpace(host, token, fmt.Sprintf("test%d", i), fmt.Sprintf("test%d", i))
		log.Println(createdSpace)
	}
}
