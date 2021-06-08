package applib

type ApiKeyProvider interface {
	GetApiKey() string
	SetApiKey(string)
}
