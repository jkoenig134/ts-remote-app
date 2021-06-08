package applib

import (
	"errors"
	"github.com/asaskevich/EventBus"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

type TsApp struct {
	isAuthenticated bool
	apiKeyProvider  ApiKeyProvider
	c               *websocket.Conn
	eventBus        EventBus.Bus
	authorized      bool
}

func NewTsApp(provider ApiKeyProvider) TsApp {
	eventBus := EventBus.New()
	return TsApp{
		isAuthenticated: false,
		apiKeyProvider:  provider,
		eventBus:        eventBus,
		authorized:      false,
	}
}

func (app *TsApp) SendAuthRequest(identifier, version, name, description string) error {
	if app.c == nil {
		return errors.New("not connected")
	}

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

	return app.c.WriteJSON(authRequest)
}

func (app *TsApp) Disconnect() {
	if app.c != nil {
		_ = app.c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		_ = app.c.Close()
		app.c = nil
	}

	app.authorized = false
	app.eventBus.Publish(string(DisconnectEvent))
}

func (app *TsApp) Connect() error {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, _, err := websocket.DefaultDialer.Dial("ws://localhost:5899", nil)
	if err != nil {
		return err
	}

	app.c = c

	go app.receive()

	return nil
}
