package tokens

import (
	"encoding/base64"
	"fmt"
)

type TokenService struct{}

func (t TokenService) GetToken(u string, p string) string {
	auth := fmt.Sprintf("%s:%s", u, p)
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
