package utils

import (
	"encoding/base64"
)

func ConvertToBase64(value string) string {
	bytes := []byte(value)
	return base64.StdEncoding.EncodeToString(bytes)
}
