package gollama

import (
	"encoding/base64"
	"fmt"
	"os"
)

func base64EncodeFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	if len(data) == 0 {
		return "", fmt.Errorf("%s contains 0 bytes", filePath)
	}
	encoded := base64.StdEncoding.EncodeToString(data)
	return encoded, nil
}
