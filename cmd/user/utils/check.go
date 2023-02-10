package utils

import (
	"encoding/base64"
	"log"

	"golang.org/x/crypto/scrypt"
)

// ScryptPwd 加密
func ScryptPwd(password string) string {
	const PwdHashByte = 10
	salt := make([]byte, 8)
	salt = []byte{200, 20, 9, 29, 15, 50, 80, 7}

	key, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, PwdHashByte)
	if err != nil {
		log.Fatal(err)
	}
	FinPwd := base64.StdEncoding.EncodeToString(key)
	return FinPwd
}
