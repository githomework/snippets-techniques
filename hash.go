package main

import (
	"crypto/sha1"
	"encoding/base64"
	"io"
)

func hash(code string, salt string) string {
	h := sha1.New()
	io.WriteString(h, code+salt)
	return base64.StdEncoding.EncodeToString(h.Sum(nil)) + salt
}
