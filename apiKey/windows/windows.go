package windows

import (
	"fmt"
	"log"

	"github.com/danieljoos/wincred"
)

type WindowsApiKeyProvider struct {
	targetName string
}

func NewApiKeyProvider(targetName string) WindowsApiKeyProvider {
	return WindowsApiKeyProvider{targetName: targetName}
}

func (p WindowsApiKeyProvider) GetApiKey() string {
	cred, err := wincred.GetGenericCredential(p.targetName)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(cred.CredentialBlob)
}

func (p WindowsApiKeyProvider) SetApiKey(apiKey string) {
	cred := wincred.NewGenericCredential(p.targetName)
	cred.CredentialBlob = []byte(apiKey)
	err := cred.Write()

	if err != nil {
		log.Println(err)
	}
}
