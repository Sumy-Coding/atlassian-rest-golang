package serv

import "os"

func GetAuthData() (string, string, string) {
	confUrl := os.Getenv("ATLAS_URL")
	user := os.Getenv("ATLAS_USER")
	pass := os.Getenv("ATLAS_PASS")
	return confUrl, user, pass
}
