package applib

type Request struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type AuthRequestPayload struct {
	Identifier  string                    `json:"identifier"`
	Version     string                    `json:"version"`
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Content     AuthRequestPayloadContent `json:"content"`
}

type AuthRequestPayloadContent struct {
	ApiKey string `json:"apiKey"`
}
