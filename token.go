package atlas

import "encoding/base64"

type TokenService struct {
}

func (t TokenService) GetToken(u string, p string) string {
	auth := u + ":" + p
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
