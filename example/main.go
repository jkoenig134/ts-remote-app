package main

import (
	"fmt"
	applib "github.com/jkoenig134/go-ts5-applib"
	"github.com/jkoenig134/go-ts5-applib/apiKey/windows"
	"log"
	"os"
	"os/signal"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	provider := windows.NewApiKeyProvider("target")
	app := applib.NewTsApp(provider)

	_ = app.SubscribeEvent(applib.ErrorEvent, func(v error) {
		fmt.Printf("error event: %s\n", v)
	})

	_ = app.SubscribeEvent(applib.ApiKeyEvent, func(v string) {
		// fmt.Printf("apiKey event: %s\n", v)
	})

	_ = app.SubscribeEvent(applib.AuthEvent, func(v applib.AuthResponsePayload) {
		fmt.Printf("auth event: %+v\n", v.ApiKey)
	})

	_ = app.SubscribeEvent(applib.DisconnectEvent, func() {
		fmt.Println("Disconnect Event")
	})

	_ = app.SubscribeEvent(applib.ClientSelfPropertyUpdatedEvent, func(v applib.ClientSelfPropertyUpdatedResponsePayload) {
		// fmt.Printf("client self prop updated: %+v\n", v)
	})

	_ = app.SubscribeEvent(applib.ClientPropertyUpdatedEvent, func(v applib.ClientPropertyUpdatedResponsePayload) {
		// fmt.Printf("client prop updated: %+v\n", v)
	})

	_ = app.SubscribeEvent(applib.ChannelPropertiesUpdatedEvent, func(v applib.ChannelPropertiesUpdatedResponsePayload) {
		// fmt.Printf("channel props updated: %+v\n", v)
	})

	_ = app.SubscribeEvent(applib.ConnectStatusChangedEvent, func(v applib.ConnectStatusChangedResponsePayload) {
		// fmt.Printf("channel props updated: %+v\n", v)
	})

	_ = app.SubscribeEvent(applib.ChannelsSubscribedEvent, func(v applib.ChannelsSubscribedResponsePayload) {
		// fmt.Printf("channel props updated: %+v\n", v)
	})

	_ = app.SubscribeEvent(applib.ClientMovedEvent, func(v applib.ClientMovedResponsePayload) {
		// fmt.Printf("client moved: %+v\n", v)
	})

	_ = app.SubscribeEvent(applib.ClientChannelGroupChangedEvent, func(v applib.ClientChannelGroupChangedResponsePayload) {
		// fmt.Printf("client channel group change %+v\n", v)
	})

	_ = app.SubscribeEvent(applib.ChannelsEvent, func(v applib.ChannelsResponsePayload) {
		// fmt.Printf("channels %+v\n", v)
	})

	_ = app.SubscribeEvent(applib.GroupInfoEvent, func(v applib.GroupInfoResponsePayload) {
		// fmt.Printf("group info %+v\n", v)
	})

	_ = app.SubscribeEvent(applib.PermissionListEvent, func(v applib.PermissionListResponsePayload) {
		// fmt.Printf("permission list %+v\n", v)
	})

	_ = app.SubscribeEvent(applib.ServerPropertiesUpdatedEvent, func(v applib.ServerPropertiesUpdatedResponsePayload) {
		// fmt.Printf("server properties updated %+v\n", v)
	})

	_ = app.SubscribeEvent(applib.NeededPermissionsEvent, func(v applib.NeededPermissionsResponsePayload) {
		// fmt.Printf("needed permissions %+v\n", v)
	})

	err := app.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = app.SendAuthRequest("de.cool", "1.0.0", "Cool", "Really cool app")
	if err != nil {
		log.Fatal(err)
	}

	<-interrupt
}
