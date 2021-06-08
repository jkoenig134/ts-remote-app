package publisher

import "github.com/jkoenig134/ts-remote-app/types"

type AuthResponse struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Payload AuthResponsePayload `json:"payload"`
}

type AuthResponsePayload struct {
	ApiKey              string             `json:"apiKey"`
	Connections         []types.Connection `json:"connections"`
	CurrentConnectionId int                `json:"currentConnectionId"`
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

type ChannelsSubscribedResponse struct {
	Payload ChannelsSubscribedResponsePayload `json:"payload"`
}

type ChannelsSubscribedResponsePayload struct {
	ChannelIds   interface{} `json:"channelIds"` //TODO
	ConnectionId int         `json:"connectionId"`
}

type ClientMovedResponse struct {
	Payload ClientMovedResponsePayload `json:"payload"`
}

type ClientMovedResponsePayload struct {
	ClientId     int                    `json:"clientId"`
	ConnectionId int                    `json:"connectionId"`
	Message      string                 `json:"message"`
	NewChannelId string                 `json:"newChannelId"`
	OldChannelId string                 `json:"oldChannelId"`
	Type         int                    `json:"type"`
	Visibility   int                    `json:"visibility"`
	Properties   types.ClientProperties `json:"properties,omitempty"`
}

type ClientPropertyUpdatedResponse struct {
	Payload ClientPropertyUpdatedResponsePayload `json:"payload"`
}

type ClientPropertyUpdatedResponsePayload struct {
	ClientId     int                    `json:"clientId"`
	ConnectionId int                    `json:"connectionId"`
	Properties   types.ClientProperties `json:"properties"`
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

type ChannelPropertiesUpdatedResponse struct {
	Payload ChannelPropertiesUpdatedResponsePayload `json:"payload"`
}

type ChannelPropertiesUpdatedResponsePayload struct {
	ChannelId    string                 `json:"channelId"`
	ConnectionId int                    `json:"connectionId"`
	Properties   map[string]interface{} `json:"properties"`
}

type ChannelsResponse struct {
	Payload ChannelsResponsePayload `json:"payload"`
}

type ChannelsResponsePayload struct {
	ConnectionId int               `json:"connectionId"`
	Info         types.ChannelInfo `json:"info"`
}

type GroupInfoResponse struct {
	Payload GroupInfoResponsePayload `json:"payload"`
}

type GroupInfoResponsePayload struct {
	ConnectionId int               `json:"connectionId"`
	Data         []types.GroupInfo `json:"data"`
	Type         int               `json:"type"`
}

type PermissionListResponse struct {
	Payload PermissionListResponsePayload `json:"payload"`
	Type    string                        `json:"type"`
}

type PermissionListResponsePayload struct {
	ConnectionId int                `json:"connectionId"`
	Data         []types.Permission `json:"data"`
	GroupEndIds  interface{}        `json:"groupEndIds"` //TODO
}

type ServerPropertiesUpdatedResponse struct {
	Payload ServerPropertiesUpdatedResponsePayload `json:"payload"`
}

type ServerPropertiesUpdatedResponsePayload struct {
	ConnectionId int                    `json:"connectionId"`
	Properties   map[string]interface{} `json:"properties"`
}

type NeededPermissionsResponse struct {
	Payload NeededPermissionsResponsePayload `json:"payload"`
}

type NeededPermissionsResponsePayload struct {
	ConnectionId int            `json:"connectionId"`
	Data         map[string]int `json:"data"`
}

type ButtonPressResponse struct {
	Payload struct {
		Button string `json:"button"`
		State  bool   `json:"state"`
	} `json:"payload"`
	ReturnCode string `json:"returnCode"`
	Status     struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Type string `json:"type"`
}
