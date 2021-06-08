package applib

import "errors"

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

func (app *RemoteApp) SendJson(request interface{}) error {
	if app.c == nil {
		return errors.New("not connected")
	}

	if !app.authorized {
		return errors.New("not authorized")
	}

	return app.c.WriteJSON(request)
}

func (app *RemoteApp) SendAuthRequest(identifier, version, name, description string) error {
	authRequest := Request{
		Type: "auth",
		Payload: AuthRequestPayload{
			Identifier:  identifier,
			Version:     version,
			Name:        name,
			Description: description,
			Content: AuthRequestPayloadContent{
				ApiKey: app.apiKeyProvider.GetApiKey(),
			},
		},
	}

	return app.SendJson(authRequest)
}
