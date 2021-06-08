package tsremote

import "errors"

type Request struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

func (app *RemoteApp) sendJson(eventType string, payload interface{}) error {
	if app.c == nil {
		return errors.New("not connected")
	}

	if !app.authorized {
		return errors.New("not authorized")
	}

	return app.c.WriteJSON(Request{
		Type:    eventType,
		Payload: payload,
	})
}

type ButtonPressPayload struct {
	Button string `json:"button"`
	State  bool   `json:"state"`
}

func (app *RemoteApp) SendButtonPress(button string, state bool) error {
	payload := ButtonPressPayload{
		Button: button,
		State:  state,
	}

	return app.sendJson("buttonPress", payload)
}

type AuthRequestPayload struct {
	Identifier  string `json:"identifier"`
	Version     string `json:"version"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Content     struct {
		ApiKey string `json:"apiKey"`
	} `json:"content"`
}

func (app *RemoteApp) SendAuthRequest(identifier, version, name, description string) error {
	if app.c == nil {
		return errors.New("not connected")
	}

	payload := AuthRequestPayload{
		Identifier:  identifier,
		Version:     version,
		Name:        name,
		Description: description,
		Content: struct {
			ApiKey string `json:"apiKey"`
		}(struct{ ApiKey string }{ApiKey: app.apiKeyProvider.GetApiKey()}),
	}

	return app.c.WriteJSON(Request{
		Type:    "auth",
		Payload: payload,
	})
}
