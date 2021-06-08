package tsremote

import (
	"encoding/json"
	"fmt"
	"github.com/jkoenig134/ts-remote-app/event"
	"github.com/jkoenig134/ts-remote-app/publisher"
	"log"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

type RemoteApp struct {
	apiKeyProvider ApiKeyProvider
	c              *websocket.Conn
	authorized     bool
	eventPublisher publisher.AppClientEventPublisher
}

func NewTsApp(provider ApiKeyProvider) RemoteApp {
	return RemoteApp{
		apiKeyProvider: provider,
		authorized:     false,
		eventPublisher: publisher.NewParser(),
	}
}

func (app *RemoteApp) Disconnect() {
	if app.c != nil {
		_ = app.c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		_ = app.c.Close()
		app.c = nil
	}

	app.authorized = false
	app.eventPublisher.PublishEvent(event.DisconnectEvent)
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

func (app *RemoteApp) SubscribeEvent(event event.AppEvent, fn interface{}) error {
	return app.eventPublisher.SubscribeEvent(event, fn)
}

func (app *RemoteApp) UnsubscribeEvent(event event.AppEvent, handler interface{}) error {
	return app.eventPublisher.UnsubscribeEvent(event, handler)
}

type TypeScan struct {
	Type string `json:"type"`
}

func (app *RemoteApp) receive() {
	defer app.Disconnect()
	for {
		_, message, err := app.c.ReadMessage()
		if err != nil {
			return
		}

		resp := TypeScan{}
		if err = json.Unmarshal(message, &resp); err != nil {

			app.eventPublisher.PublishEvent(event.ErrorEvent, fmt.Errorf("deserialize: %s", message))
			continue
		}

		switch resp.Type {
		case "auth":
			if key := app.eventPublisher.HandleAuthResponse(message); key != nil {
				app.authorized = true
				app.apiKeyProvider.SetApiKey(*key)
			}
		case "clientSelfPropertyUpdated":
			app.eventPublisher.HandleClientSelfPropertyUpdated(message)
		case "clientPropertiesUpdated":
			app.eventPublisher.HandleClientPropertyUpdated(message)
		case "connectStatusChanged":
			app.eventPublisher.HandleConnectStatusChanged(message)
		case "channelsSubscribed":
			app.eventPublisher.HandleChannelsSubscribed(message)
		case "clientMoved":
			app.eventPublisher.HandleClientMoved(message)
		case "channelPropertiesUpdated":
			app.eventPublisher.HandleChannelPropertiesUpdated(message)
		case "clientChannelGroupChanged":
			app.eventPublisher.HandleClientChannelGroupChanged(message)
		case "channels":
			app.eventPublisher.HandleChannels(message)
		case "groupInfo":
			app.eventPublisher.HandleGroupInfo(message)
		case "permissionList":
			app.eventPublisher.HandlePermissionList(message)
		case "serverPropertiesUpdated":
			app.eventPublisher.HandleServerPropertiesUpdated(message)
		case "neededPermissions":
			app.eventPublisher.HandleNeededPermissions(message)
		case "buttonPress":
			app.eventPublisher.HandleButtonPress(message)

		default:
			log.Printf("not recognized event %s: %s\n", resp.Type, string(message))
		}
	}
}
