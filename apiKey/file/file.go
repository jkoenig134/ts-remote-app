package file

import (
	"errors"
	"fmt"
	"os"
)

type FileApiKeyProvider struct {
	filename string
}

func NewApiKeyProvider(filename string) FileApiKeyProvider {
	return FileApiKeyProvider{filename: filename}
}

func (p FileApiKeyProvider) GetApiKey() string {
	if _, err := os.Stat(""); errors.Is(err, os.ErrNotExist) {
		return ""
	}

	content, _ := os.ReadFile(p.filename)
	return string(content)
}

func (p FileApiKeyProvider) SetApiKey(apiKey string) {
	file, err := os.OpenFile(p.filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	file.WriteString(apiKey)
}
