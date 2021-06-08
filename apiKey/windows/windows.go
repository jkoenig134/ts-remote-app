package windows

import (
	"fmt"
	"github.com/danieljoos/wincred"
	"log"
)

type ApiKeyProvider struct {
	targetName string
}

func NewApiKeyProvider(targetName string) ApiKeyProvider {
	return ApiKeyProvider{targetName: targetName}
}

func (p ApiKeyProvider) GetApiKey() string {
	cred, err := wincred.GetGenericCredential(p.targetName)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(cred.CredentialBlob)
}

func (p ApiKeyProvider) SetApiKey(apiKey string) {
	cred := wincred.NewGenericCredential(p.targetName)
	cred.CredentialBlob = []byte(apiKey)
	err := cred.Write()

	if err != nil {
		log.Println(err)
	}
}
