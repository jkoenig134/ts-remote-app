package inMemory

type InMemoryApiKeyProvider struct {
	ApiKey string
}

func (p InMemoryApiKeyProvider) GetApiKey() string {
	return p.ApiKey
}

func (p InMemoryApiKeyProvider) SetApiKey(apiKey string) {
	p.ApiKey = apiKey
}
