package event

import (
	"github.com/jkoenig134/ts-remote-app/types"
)

type AppEvent string

const (
	ErrorEvent                     AppEvent = "ERROR"
	AuthEvent                      AppEvent = "AUTH"
	DisconnectEvent                AppEvent = "DISCONNECT"
	ClientSelfPropertyUpdatedEvent AppEvent = "CLIENT_SELF_PROPERTY_UPDATED"
	ClientPropertyUpdatedEvent     AppEvent = "CLIENT_PROPERTY_UPDATED"
	ChannelPropertiesUpdatedEvent  AppEvent = "CHANNEL_PROPERTIES_UPDATED"
	ConnectStatusChangedEvent      AppEvent = "CONNECT_STATUS_CHANGED_EVENT"
	ChannelsSubscribedEvent        AppEvent = "CHANNELS_SUBSCRIBED"
	ClientMovedEvent               AppEvent = "CLIENT_MOVED"
	ClientChannelGroupChangedEvent AppEvent = "CLIENT_CHANNEL_GROUP_CHANGED"
	ChannelsEvent                  AppEvent = "CHANNELS"
	GroupInfoEvent                 AppEvent = "GROUP_INFO"
	PermissionListEvent            AppEvent = "PERMISSION_LIST"
	ServerPropertiesUpdatedEvent   AppEvent = "SERVER_PROPERTIES_UPDATED"
	NeededPermissionsEvent         AppEvent = "NEEDED_PERMISSIONS"
	ButtonPressEvent               AppEvent = "BUTTON_PRESS"
)

type ErrorEventPayload = error

type AuthEventPayload struct {
	ApiKey              string             `json:"apiKey"`
	Connections         []types.Connection `json:"connections"`
	CurrentConnectionId int                `json:"currentConnectionId"`
}

type ClientSelfPropertyUpdatedEventPayload struct {
	ConnectionId int    `json:"connectionId"`
	Flag         string `json:"flag"`
	NewValue     string `json:"newValue"`
	OldValue     string `json:"oldValue"`
}

type ClientPropertyUpdatedEventPayload struct {
	ClientId     int                    `json:"clientId"`
	ConnectionId int                    `json:"connectionId"`
	Properties   types.ClientProperties `json:"properties"`
}

type ChannelPropertiesUpdatedEventPayload struct {
	ChannelId    string                 `json:"channelId"`
	ConnectionId int                    `json:"connectionId"`
	Properties   map[string]interface{} `json:"properties"`
}

type ConnectStatusChangedEventPayload struct {
	ConnectionId int `json:"connectionId"`
	Error        int `json:"error"`
	Info         struct {
		ClientId   int    `json:"clientId"`
		ServerName string `json:"serverName"`
		ServerUid  string `json:"serverUid"`
	} `json:"info"`
	Status int `json:"status"`
}

type ChannelsSubscribedEventPayload struct {
	ChannelIds   interface{} `json:"channelIds"` //TODO
	ConnectionId int         `json:"connectionId"`
}

type ClientMovedEventPayload struct {
	ClientId     int                    `json:"clientId"`
	ConnectionId int                    `json:"connectionId"`
	Message      string                 `json:"message"`
	NewChannelId string                 `json:"newChannelId"`
	OldChannelId string                 `json:"oldChannelId"`
	Type         int                    `json:"type"`
	Visibility   int                    `json:"visibility"`
	Properties   types.ClientProperties `json:"properties,omitempty"`
}

type ClientChannelGroupChangedEventPayload struct {
	ChannelGroupId                 string `json:"channelGroupId"`
	ChannelGroupInheritedChannelId string `json:"channelGroupInheritedChannelId"`
	ChannelId                      string `json:"channelId"`
	ClientId                       int    `json:"clientId"`
	ConnectionId                   int    `json:"connectionId"`
}

type ChannelsEventPayload struct {
	ConnectionId int               `json:"connectionId"`
	Info         types.ChannelInfo `json:"info"`
}

type GroupInfoEventPayload struct {
	ConnectionId int               `json:"connectionId"`
	Data         []types.GroupInfo `json:"data"`
	Type         int               `json:"type"`
}

type PermissionListEventPayload struct {
	ConnectionId int                `json:"connectionId"`
	Data         []types.Permission `json:"data"`
	GroupEndIds  interface{}        `json:"groupEndIds"` //TODO
}

type ServerPropertiesUpdatedEventPayload struct {
	ConnectionId int                    `json:"connectionId"`
	Properties   map[string]interface{} `json:"properties"`
}

type NeededPermissionsEventPayload struct {
	ConnectionId int            `json:"connectionId"`
	Data         map[string]int `json:"data"`
}

type ButtonPressEventPayload struct {
	Button     string
	State      bool
	ReturnCode string
	Status     struct {
		Code    int
		Message string
	}
}
