package inMemory

type ApiKeyProvider struct {
	ApiKey string
}

func (p ApiKeyProvider) GetApiKey() string {
	return p.ApiKey
}

func (p ApiKeyProvider) SetApiKey(apiKey string) {
	p.ApiKey = apiKey
}
