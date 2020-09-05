package token

import (
	"crypto/md5"
	"encoding/hex"
)

func GetMD5HashToken(userID, email string) string {

	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(userID + email))
	return hex.EncodeToString(hash.Sum(nil))
}
