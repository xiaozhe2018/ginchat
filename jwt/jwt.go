package jwt

import (
	"gitchat/utils"
	"time"
)

func GenToken(id string, name string) (token string) {
	token = utils.EncodeMD5(id + name + time.Now().Format(time.RFC3339))
	return token
}
