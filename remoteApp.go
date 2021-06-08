package applib

import (
	"github.com/asaskevich/EventBus"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

type RemoteApp struct {
	apiKeyProvider ApiKeyProvider
	c              *websocket.Conn
	eventBus       EventBus.Bus
	authorized     bool
}

func NewTsApp(provider ApiKeyProvider) RemoteApp {
	eventBus := EventBus.New()
	return RemoteApp{
		apiKeyProvider: provider,
		eventBus:       eventBus,
		authorized:     false,
	}
}

func (app *RemoteApp) Disconnect() {
	if app.c != nil {
		_ = app.c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		_ = app.c.Close()
		app.c = nil
	}

	app.authorized = false
	app.eventBus.Publish(string(DisconnectEvent))
}

func (app *RemoteApp) Connect() error {
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
