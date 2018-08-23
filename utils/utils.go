package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func EnMd5(EncryptedStr  string) string  {
	h := md5.New()
	h.Write([]byte(EncryptedStr))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}