package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Md5String(src string) string {
	h := md5.New()
	io.WriteString(h, src)
	return hex.EncodeToString(h.Sum(nil))
}
