package utils

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"mini-min-tiktok/pkg/dal/model"
	"mini-min-tiktok/pkg/dal/query"

	"golang.org/x/crypto/scrypt"
)

func CheckUser(q *query.Query, username, password string) (res *model.TUser, err error) {
	user := q.TUser
	res, _ = user.WithContext(context.Background()).
		Where(user.Name.Eq(username)).
		First()
	if res == nil {
		err = fmt.Errorf("用户不存在: %v", username)
		return
	}
	if pwd := ScryptPwd(password); pwd != res.Password {
		err = fmt.Errorf("密码错误: %v", password)
		return
	}
	return
}

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
