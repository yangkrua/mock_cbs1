package utils

import (
	"encoding/base64"
)

func Base64EncodeToString(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func Base64DecodeString(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
