package applib

type AppEvent string

const (
	ErrorEvent                     AppEvent = "ERROR"
	ApiKeyEvent                    AppEvent = "API_KEY"
	AuthEvent                      AppEvent = "AUTH_EVENT"
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
)

func (app *RemoteApp) SubscribeEvent(event AppEvent, fn interface{}) error {
	return app.eventBus.Subscribe(string(event), fn)
}

func (app *RemoteApp) UnsubscribeEvent(event AppEvent, handler interface{}) error {
	return app.eventBus.Unsubscribe(string(event), handler)
}

func (app *RemoteApp) publishEvent(event AppEvent, args ...interface{}) {
	app.eventBus.Publish(string(event), args...)
}
