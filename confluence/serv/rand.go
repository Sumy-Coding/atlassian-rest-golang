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

func (r RandService) RandomString(leng int) string {
	rStr := "Lorem ipsum dolor sit amet consectetur adipiscing, elit ultricies aliquam nostra sapien turpis vivamus, natoque in eleifend"
	split := strings.Split(rStr, " ")
	var fStr string
	for i := 0; i < leng; i++ {
		rInt := rand.Intn(len(split) - 1)
		fStr = fStr + " " + split[rInt]
	}
	return fStr
}
