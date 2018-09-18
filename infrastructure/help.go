package infrastructure

import (
	"encoding/base64"
	"fmt"
)

func Help() string {
	return ""
}

func BasicAuthTokenEncode(username string, password string) string {
	data := []byte(fmt.Sprintf("%s:%s", username, password))
	return base64.StdEncoding.EncodeToString(data)
}

func BasicAuthTokenDecode(token string) []byte {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil
	}
	return data
}
