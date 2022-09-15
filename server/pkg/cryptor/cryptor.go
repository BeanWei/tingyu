package cryptor

import (
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

// HashUserPwd 用户明文密码加密
func HashUserPwd(password, salt string) string {
	passwd := pbkdf2.Key([]byte(password), []byte(salt), 10000, 50, sha256.New)
	return fmt.Sprintf("%x", passwd)
}
