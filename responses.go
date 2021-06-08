package applib

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

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
			app.publishEvent(ErrorEvent, errors.New(fmt.Sprintf("deserialize: %s", message)))
			continue
		}

		switch resp.Type {
		case "auth":
			app.handleAuthResponse(message)
		case "clientSelfPropertyUpdated":
			app.handleClientSelfPropertyUpdated(message)
		case "clientPropertiesUpdated":
			app.handleClientPropertyUpdated(message)
		case "connectStatusChanged":
			app.handleConnectStatusChanged(message)
		case "channelsSubscribed":
			app.handleChannelsSubscribed(message)
		case "clientMoved":
			app.handleClientMoved(message)
		case "channelPropertiesUpdated":
			app.handleChannelPropertiesUpdated(message)
		case "clientChannelGroupChanged":
			app.handleClientChannelGroupChanged(message)
		case "channels":
			app.handleChannels(message)
		case "groupInfo":
			app.handleGroupInfo(message)
		case "permissionList":
			app.handlePermissionList(message)
		case "serverPropertiesUpdated":
			app.handleServerPropertiesUpdated(message)
		case "neededPermissions":
			app.handleNeededPermissions(message)

		default:
			log.Printf("not recognized event %s: %s\n", resp.Type, string(message))
		}
	}
}

func (app *RemoteApp) deserializeResponse(message []byte, v interface{}) error {
	err := json.Unmarshal(message, v)
	if err != nil {
		app.publishEvent(ErrorEvent, errors.New(fmt.Sprintf("deserialize: %s", message)))
	}

	return err
}

type AuthResponse struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Payload AuthResponsePayload `json:"payload"`
}

type AuthResponsePayload struct {
	ApiKey              string       `json:"apiKey"`
	Connections         []Connection `json:"connections"`
	CurrentConnectionId int          `json:"currentConnectionId"`
}

func (app *RemoteApp) handleAuthResponse(message []byte) {
	auth := AuthResponse{}
	if err := app.deserializeResponse(message, &auth); err != nil {
		return
	}

	if auth.Status.Code != 0 {
		app.publishEvent(ErrorEvent, errors.New(auth.Status.Message))
		return
	}

	app.authorized = true
	key := auth.Payload.ApiKey
	app.apiKeyProvider.SetApiKey(key)
	app.publishEvent(ApiKeyEvent, key)
	app.publishEvent(AuthEvent, auth.Payload)
}

type ClientSelfPropertyUpdatedResponse struct {
	Payload ClientSelfPropertyUpdatedResponsePayload `json:"payload"`
}

type ClientSelfPropertyUpdatedResponsePayload struct {
	ConnectionId int    `json:"connectionId"`
	Flag         string `json:"flag"`
	NewValue     string `json:"newValue"`
	OldValue     string `json:"oldValue"`
}

func (app *RemoteApp) handleClientSelfPropertyUpdated(message []byte) {
	response := ClientSelfPropertyUpdatedResponse{}
	if err := app.deserializeResponse(message, &response); err != nil {
		return
	}

	app.publishEvent(ClientSelfPropertyUpdatedEvent, response.Payload)
}

type ConnectStatusChangedResponse struct {
	Payload ConnectStatusChangedResponsePayload `json:"payload"`
}

type ConnectStatusChangedResponsePayload struct {
	ConnectionId int `json:"connectionId"`
	Error        int `json:"error"`
	Info         struct {
		ClientId   int    `json:"clientId"`
		ServerName string `json:"serverName"`
		ServerUid  string `json:"serverUid"`
	} `json:"info"`
	Status int `json:"status"`
}

func (app *RemoteApp) handleConnectStatusChanged(message []byte) {
	response := ConnectStatusChangedResponse{}
	if err := app.deserializeResponse(message, &response); err != nil {
		return
	}

	app.publishEvent(ConnectStatusChangedEvent, response.Payload)
}

type ChannelsSubscribedResponse struct {
	Payload ChannelsSubscribedResponsePayload `json:"payload"`
}

type ChannelsSubscribedResponsePayload struct {
	ChannelIds   interface{} `json:"channelIds"` //TODO
	ConnectionId int         `json:"connectionId"`
}

func (app *RemoteApp) handleChannelsSubscribed(message []byte) {
	response := ChannelsSubscribedResponse{}
	if err := app.deserializeResponse(message, &response); err != nil {
		return
	}

	app.publishEvent(ChannelsSubscribedEvent, response.Payload)
}

type ClientMovedResponse struct {
	Payload ClientMovedResponsePayload `json:"payload"`
}

type ClientMovedResponsePayload struct {
	ClientId     int              `json:"clientId"`
	ConnectionId int              `json:"connectionId"`
	Message      string           `json:"message"`
	NewChannelId string           `json:"newChannelId"`
	OldChannelId string           `json:"oldChannelId"`
	Type         int              `json:"type"`
	Visibility   int              `json:"visibility"`
	Properties   ClientProperties `json:"properties,omitempty"`
}

func (app *RemoteApp) handleClientMoved(message []byte) {
	response := ClientMovedResponse{}
	if err := app.deserializeResponse(message, &response); err != nil {
		return
	}

	app.publishEvent(ClientMovedEvent, response.Payload)
}

type ClientPropertyUpdatedResponse struct {
	Payload ClientPropertyUpdatedResponsePayload `json:"payload"`
}

type ClientPropertyUpdatedResponsePayload struct {
	ClientId     int              `json:"clientId"`
	ConnectionId int              `json:"connectionId"`
	Properties   ClientProperties `json:"properties"`
}

func (app *RemoteApp) handleClientPropertyUpdated(message []byte) {
	response := ClientPropertyUpdatedResponse{}
	if err := app.deserializeResponse(message, &response); err != nil {
		return
	}

	app.publishEvent(ClientPropertyUpdatedEvent, response.Payload)
}

type ClientChannelGroupChangedResponse struct {
	Payload ClientChannelGroupChangedResponsePayload `json:"payload"`
}

type ClientChannelGroupChangedResponsePayload struct {
	ChannelGroupId                 string `json:"channelGroupId"`
	ChannelGroupInheritedChannelId string `json:"channelGroupInheritedChannelId"`
	ChannelId                      string `json:"channelId"`
	ClientId                       int    `json:"clientId"`
	ConnectionId                   int    `json:"connectionId"`
}

func (app *RemoteApp) handleClientChannelGroupChanged(message []byte) {
	response := ClientChannelGroupChangedResponse{}
	if err := app.deserializeResponse(message, &response); err != nil {
		return
	}

	app.publishEvent(ClientChannelGroupChangedEvent, response.Payload)
}

type ChannelPropertiesUpdatedResponse struct {
	Payload ChannelPropertiesUpdatedResponsePayload `json:"payload"`
}

type ChannelPropertiesUpdatedResponsePayload struct {
	ChannelId    string                 `json:"channelId"`
	ConnectionId int                    `json:"connectionId"`
	Properties   map[string]interface{} `json:"properties"`
}

func (app *RemoteApp) handleChannelPropertiesUpdated(message []byte) {
	response := ChannelPropertiesUpdatedResponse{}
	if err := app.deserializeResponse(message, &response); err != nil {
		return
	}

	app.publishEvent(ChannelPropertiesUpdatedEvent, response.Payload)
}

type ChannelsResponse struct {
	Payload ChannelsResponsePayload `json:"payload"`
}

type ChannelsResponsePayload struct {
	ConnectionId int         `json:"connectionId"`
	Info         ChannelInfo `json:"info"`
}

func (app *RemoteApp) handleChannels(message []byte) {
	response := ChannelsResponse{}
	if err := app.deserializeResponse(message, &response); err != nil {
		return
	}

	app.publishEvent(ChannelsEvent, response.Payload)
}

type GroupInfoResponse struct {
	Payload GroupInfoResponsePayload `json:"payload"`
}

type GroupInfoResponsePayload struct {
	ConnectionId int         `json:"connectionId"`
	Data         []GroupInfo `json:"data"`
	Type         int         `json:"type"`
}

func (app *RemoteApp) handleGroupInfo(message []byte) {
	response := GroupInfoResponse{}
	if err := app.deserializeResponse(message, &response); err != nil {
		return
	}

	app.publishEvent(GroupInfoEvent, response.Payload)
}

type PermissionListResponse struct {
	Payload PermissionListResponsePayload `json:"payload"`
	Type    string                        `json:"type"`
}

type PermissionListResponsePayload struct {
	ConnectionId int          `json:"connectionId"`
	Data         []Permission `json:"data"`
	GroupEndIds  interface{}  `json:"groupEndIds"` //TODO
}

func (app *RemoteApp) handlePermissionList(message []byte) {
	response := PermissionListResponse{}
	if err := app.deserializeResponse(message, &response); err != nil {
		return
	}

	app.publishEvent(PermissionListEvent, response.Payload)
}

type ServerPropertiesUpdatedResponse struct {
	Payload ServerPropertiesUpdatedResponsePayload `json:"payload"`
}

type ServerPropertiesUpdatedResponsePayload struct {
	ConnectionId int                    `json:"connectionId"`
	Properties   map[string]interface{} `json:"properties"`
}

func (app *RemoteApp) handleServerPropertiesUpdated(message []byte) {
	response := ServerPropertiesUpdatedResponse{}
	if err := app.deserializeResponse(message, &response); err != nil {
		return
	}

	app.publishEvent(ServerPropertiesUpdatedEvent, response.Payload)
}

type NeededPermissionsResponse struct {
	Payload NeededPermissionsResponsePayload `json:"payload"`
}

type NeededPermissionsResponsePayload struct {
	ConnectionId int            `json:"connectionId"`
	Data         map[string]int `json:"data"`
}

func (app *RemoteApp) handleNeededPermissions(message []byte) {
	response := NeededPermissionsResponse{}
	if err := app.deserializeResponse(message, &response); err != nil {
		return
	}

	app.publishEvent(NeededPermissionsEvent, response.Payload)
}
