package tsremote

type ApiKeyProvider interface {
	GetApiKey() string
	SetApiKey(string)
}
