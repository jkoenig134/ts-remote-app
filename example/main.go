package main

import (
	"fmt"
	"github.com/jkoenig134/ts-remote-app"
	"github.com/jkoenig134/ts-remote-app/apiKey/windows"
	"github.com/jkoenig134/ts-remote-app/event"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	provider := windows.NewApiKeyProvider("target")
	app := tsremote.NewTsApp(provider)

	_ = app.SubscribeEvent(event.ErrorEvent, func(v event.ErrorEventPayload) {
		fmt.Printf("error event: %s\n", v)
	})

	_ = app.SubscribeEvent(event.AuthEvent, func(v event.AuthEventPayload) {
		fmt.Printf("auth event: %+v\n", v.ApiKey)
	})

	_ = app.SubscribeEvent(event.DisconnectEvent, func() {
		fmt.Println("Disconnect Event")

		interrupt <- os.Interrupt
	})

	_ = app.SubscribeEvent(event.ClientSelfPropertyUpdatedEvent, func(v event.ClientSelfPropertyUpdatedEventPayload) {
		// fmt.Printf("client self prop updated: %+v\n", v)
	})

	_ = app.SubscribeEvent(event.ClientPropertyUpdatedEvent, func(v event.ClientPropertyUpdatedEventPayload) {
		// fmt.Printf("client prop updated: %+v\n", v)
	})

	_ = app.SubscribeEvent(event.ChannelPropertiesUpdatedEvent, func(v event.ChannelPropertiesUpdatedEventPayload) {
		// fmt.Printf("channel props updated: %+v\n", v)
	})

	_ = app.SubscribeEvent(event.ConnectStatusChangedEvent, func(v event.ConnectStatusChangedEventPayload) {
		// fmt.Printf("channel props updated: %+v\n", v)
	})

	_ = app.SubscribeEvent(event.ChannelsSubscribedEvent, func(v event.ChannelsSubscribedEventPayload) {
		// fmt.Printf("channel props updated: %+v\n", v)
	})

	_ = app.SubscribeEvent(event.ClientMovedEvent, func(v event.ClientMovedEventPayload) {
		// fmt.Printf("client moved: %+v\n", v)
	})

	_ = app.SubscribeEvent(event.ClientChannelGroupChangedEvent, func(v event.ClientChannelGroupChangedEventPayload) {
		// fmt.Printf("client channel group change %+v\n", v)
	})

	_ = app.SubscribeEvent(event.ChannelsEvent, func(v event.ChannelsEventPayload) {
		// fmt.Printf("channels %+v\n", v)
	})

	_ = app.SubscribeEvent(event.GroupInfoEvent, func(v event.GroupInfoEventPayload) {
		// fmt.Printf("group info %+v\n", v)
	})

	_ = app.SubscribeEvent(event.PermissionListEvent, func(v event.PermissionListEventPayload) {
		// fmt.Printf("permission list %+v\n", v)
	})

	_ = app.SubscribeEvent(event.ServerPropertiesUpdatedEvent, func(v event.ServerPropertiesUpdatedEventPayload) {
		// fmt.Printf("server properties updated %+v\n", v)
	})

	_ = app.SubscribeEvent(event.NeededPermissionsEvent, func(v event.NeededPermissionsEventPayload) {
		// fmt.Printf("needed permissions %+v\n", v)
	})

	_ = app.SubscribeEvent(event.ButtonPressEvent, func(v event.ButtonPressEventPayload) {
		fmt.Printf("button press %+v\n", v)
	})

	err := app.Connect()
	if err != nil {
		log.Fatal(err)
	}

	err = app.SendAuthRequest("de.cool", "1.0.0", "Cool", "Really cool app")
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)
	log.Println("enable button")
	_ = app.SendButtonPress("testButton", true)

	time.Sleep(1 * time.Second)
	log.Println("disable button")
	_ = app.SendButtonPress("testButton", false)

	<-interrupt
}
