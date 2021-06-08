package publisher

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/jkoenig134/ts-remote-app/event"
)

type AppClientEventPublisher struct {
	eventBus EventBus.Bus
}

func NewParser() AppClientEventPublisher {
	return AppClientEventPublisher{
		eventBus: EventBus.New(),
	}
}

func (publisher *AppClientEventPublisher) PublishEvent(event event.AppEvent, args ...interface{}) {
	publisher.eventBus.Publish(string(event), args...)
}

func (publisher *AppClientEventPublisher) SubscribeEvent(event event.AppEvent, fn interface{}) error {
	return publisher.eventBus.Subscribe(string(event), fn)
}

func (publisher *AppClientEventPublisher) UnsubscribeEvent(event event.AppEvent, handler interface{}) error {
	return publisher.eventBus.Unsubscribe(string(event), handler)
}

func (publisher *AppClientEventPublisher) deserializeResponse(message []byte, v interface{}) error {
	err := json.Unmarshal(message, v)
	if err != nil {
		publisher.PublishEvent(event.ErrorEvent, fmt.Errorf("deserialize: %s", message))
	}

	return err
}

func (publisher *AppClientEventPublisher) HandleAuthResponse(message []byte) *string {
	auth := AuthResponse{}
	if err := publisher.deserializeResponse(message, &auth); err != nil {
		return nil
	}

	if auth.Status.Code != 0 {
		publisher.PublishEvent(event.ErrorEvent, errors.New(auth.Status.Message))
		return nil
	}

	publisher.PublishEvent(event.AuthEvent, event.AuthEventPayload(auth.Payload))
	return &auth.Payload.ApiKey
}

func (publisher *AppClientEventPublisher) HandleClientSelfPropertyUpdated(message []byte) {
	response := ClientSelfPropertyUpdatedResponse{}
	if err := publisher.deserializeResponse(message, &response); err != nil {
		return
	}

	publisher.PublishEvent(event.ClientSelfPropertyUpdatedEvent, event.ClientSelfPropertyUpdatedEventPayload(response.Payload))
}

func (publisher *AppClientEventPublisher) HandleConnectStatusChanged(message []byte) {
	response := ConnectStatusChangedResponse{}
	if err := publisher.deserializeResponse(message, &response); err != nil {
		return
	}

	publisher.PublishEvent(event.ConnectStatusChangedEvent, event.ConnectStatusChangedEventPayload(response.Payload))
}

func (publisher *AppClientEventPublisher) HandleClientMoved(message []byte) {
	response := ClientMovedResponse{}
	if err := publisher.deserializeResponse(message, &response); err != nil {
		return
	}

	publisher.PublishEvent(event.ClientMovedEvent, event.ClientMovedEventPayload(response.Payload))
}

func (publisher *AppClientEventPublisher) HandleClientPropertyUpdated(message []byte) {
	response := ClientPropertyUpdatedResponse{}
	if err := publisher.deserializeResponse(message, &response); err != nil {
		return
	}

	publisher.PublishEvent(event.ClientPropertyUpdatedEvent, event.ClientPropertyUpdatedEventPayload(response.Payload))
}

func (publisher *AppClientEventPublisher) HandleClientChannelGroupChanged(message []byte) {
	response := ClientChannelGroupChangedResponse{}
	if err := publisher.deserializeResponse(message, &response); err != nil {
		return
	}

	publisher.PublishEvent(event.ClientChannelGroupChangedEvent, event.ClientChannelGroupChangedEventPayload(response.Payload))
}

func (publisher *AppClientEventPublisher) HandleChannelPropertiesUpdated(message []byte) {
	response := ChannelPropertiesUpdatedResponse{}
	if err := publisher.deserializeResponse(message, &response); err != nil {
		return
	}

	publisher.PublishEvent(event.ChannelPropertiesUpdatedEvent, event.ChannelPropertiesUpdatedEventPayload(response.Payload))
}

func (publisher *AppClientEventPublisher) HandleChannels(message []byte) {
	response := ChannelsResponse{}
	if err := publisher.deserializeResponse(message, &response); err != nil {
		return
	}

	publisher.PublishEvent(event.ChannelsEvent, event.ChannelsEventPayload(response.Payload))
}

func (publisher *AppClientEventPublisher) HandleChannelsSubscribed(message []byte) {
	response := ChannelsSubscribedResponse{}
	if err := publisher.deserializeResponse(message, &response); err != nil {
		return
	}

	publisher.PublishEvent(event.ChannelsSubscribedEvent, event.ChannelsSubscribedEventPayload(response.Payload))
}

func (publisher *AppClientEventPublisher) HandleGroupInfo(message []byte) {
	response := GroupInfoResponse{}
	if err := publisher.deserializeResponse(message, &response); err != nil {
		return
	}

	publisher.PublishEvent(event.GroupInfoEvent, event.GroupInfoEventPayload(response.Payload))
}

func (publisher *AppClientEventPublisher) HandlePermissionList(message []byte) {
	response := PermissionListResponse{}
	if err := publisher.deserializeResponse(message, &response); err != nil {
		return
	}

	publisher.PublishEvent(event.PermissionListEvent, event.PermissionListEventPayload(response.Payload))
}

func (publisher *AppClientEventPublisher) HandleServerPropertiesUpdated(message []byte) {
	response := ServerPropertiesUpdatedResponse{}
	if err := publisher.deserializeResponse(message, &response); err != nil {
		return
	}

	publisher.PublishEvent(event.ServerPropertiesUpdatedEvent, event.ServerPropertiesUpdatedEventPayload(response.Payload))
}

func (publisher *AppClientEventPublisher) HandleNeededPermissions(message []byte) {
	response := NeededPermissionsResponse{}
	if err := publisher.deserializeResponse(message, &response); err != nil {
		return
	}

	publisher.PublishEvent(event.NeededPermissionsEvent, event.NeededPermissionsEventPayload(response.Payload))
}

func (publisher *AppClientEventPublisher) HandleButtonPress(message []byte) {
	response := ButtonPressResponse{}
	if err := publisher.deserializeResponse(message, &response); err != nil {
		return
	}

	payload := event.ButtonPressEventPayload{
		Button:     response.Payload.Button,
		State:      response.Payload.State,
		ReturnCode: response.ReturnCode,
		Status: struct {
			Code    int
			Message string
		}(response.Status),
	}

	publisher.PublishEvent(event.ButtonPressEvent, payload)
}
