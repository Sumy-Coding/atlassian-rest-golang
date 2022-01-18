package serv

import (
	"math/rand"
	"strings"
)

type RandService struct {
}

//func rangeIn(low, hi int) int {
//	return low + rand.Intn(hi-low)
//}

func (r RandService) RandomString() string {
	rStr := "Lorem ipsum dolor sit amet consectetur adipiscing, elit ultricies aliquam nostra sapien turpis vivamus, natoque in eleifend iaculis senectus"
	split := strings.Split(rStr, " ")
	rInt := rand.Intn(len(split) - 1)
	return split[rInt]
}
