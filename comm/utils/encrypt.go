package utils

import "encoding/base64"

func Base64(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}
