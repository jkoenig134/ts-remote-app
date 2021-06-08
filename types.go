package applib

type Connection struct {
	ChannelInfos ChannelInfo            `json:"channelInfos"`
	ClientId     int                    `json:"clientId"`
	ClientInfos  []ClientInfo           `json:"clientInfos"`
	Id           int                    `json:"id"`
	Properties   map[string]interface{} `json:"properties"`
	Status       int                    `json:"status"`
}

type Channel struct {
	Id         string                 `json:"id"`
	Order      string                 `json:"order"`
	ParentId   string                 `json:"parentId"`
	Properties map[string]interface{} `json:"properties"`
}

type ChannelInfo struct {
	RootChannels []Channel            `json:"rootChannels"`
	SubChannels  map[string][]Channel `json:"subChannels"`
}

type ClientProperties struct {
	ActiveIntegrationsInfo         string  `json:"activeIntegrationsInfo"`
	Away                           int     `json:"away"`
	AwayMessage                    string  `json:"awayMessage"`
	Badges                         string  `json:"badges"`
	ChannelGroupId                 string  `json:"channelGroupId"`
	ChannelGroupInheritedChannelId string  `json:"channelGroupInheritedChannelId"`
	Country                        string  `json:"country"`
	Created                        int     `json:"created"`
	DatabaseId                     string  `json:"databaseId"`
	DefaultChannel                 string  `json:"defaultChannel"`
	DefaultChannelPassword         string  `json:"defaultChannelPassword"`
	DefaultToken                   string  `json:"defaultToken"`
	Description                    string  `json:"description"`
	EstimatedLocation              string  `json:"estimatedLocation"`
	FlagAvatar                     string  `json:"flagAvatar"`
	FlagTalking                    int     `json:"flagTalking"`
	IconId                         int     `json:"iconId"`
	IdleTime                       int     `json:"idleTime"`
	InputDeactivated               int     `json:"inputDeactivated"`
	InputHardware                  int     `json:"inputHardware"`
	InputMuted                     int     `json:"inputMuted"`
	Integrations                   string  `json:"integrations"`
	IsChannelCommander             int     `json:"isChannelCommander"`
	IsMuted                        int     `json:"isMuted"`
	IsPrioritySpeaker              int     `json:"isPrioritySpeaker"`
	IsRecording                    int     `json:"isRecording"`
	IsTalker                       int     `json:"isTalker"`
	LastConnected                  int     `json:"lastConnected"`
	MetaData                       string  `json:"metaData"`
	MonthBytesDownloaded           int     `json:"monthBytesDownloaded"`
	MonthBytesUploaded             int     `json:"monthBytesUploaded"`
	MyteamspeakAvatar              string  `json:"myteamspeakAvatar"`
	MyteamspeakId                  string  `json:"myteamspeakId"`
	NeededServerqueryViewPower     int     `json:"neededServerqueryViewPower"`
	Nickname                       string  `json:"nickname"`
	NicknamePhonetic               string  `json:"nicknamePhonetic"`
	OutputHardware                 int     `json:"outputHardware"`
	OutputMuted                    int     `json:"outputMuted"`
	OutputonlyMuted                int     `json:"outputonlyMuted"`
	PermissionHints                int     `json:"permissionHints"`
	Platform                       string  `json:"platform"`
	ServerPassword                 string  `json:"serverPassword"`
	Servergroups                   string  `json:"servergroups"`
	SignedBadges                   string  `json:"signedBadges"`
	TalkPower                      int     `json:"talkPower"`
	TalkRequest                    int     `json:"talkRequest"`
	TalkRequestMsg                 string  `json:"talkRequestMsg"`
	TotalBytesDownloaded           int     `json:"totalBytesDownloaded"`
	TotalBytesUploaded             int     `json:"totalBytesUploaded"`
	TotalConnections               int     `json:"totalConnections"`
	Type                           int     `json:"type"`
	UniqueIdentifier               string  `json:"uniqueIdentifier"`
	UnreadMessages                 int     `json:"unreadMessages"`
	Version                        string  `json:"version"`
	VolumeModificator              float64 `json:"volumeModificator"`
}

type ClientInfo struct {
	ChannelId  int              `json:"channelId"`
	Id         int              `json:"id"`
	Properties ClientProperties `json:"properties"`
}

type Permission struct {
	Description string `json:"description"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
}

type GroupInfo struct {
	IconId                  int    `json:"iconId"`
	Id                      string `json:"id"`
	Name                    string `json:"name"`
	NameMode                int    `json:"nameMode"`
	NeededMemberAddPower    int    `json:"neededMemberAddPower"`
	NeededMemberRemovePower int    `json:"neededMemberRemovePower"`
	NeededModifyPower       int    `json:"neededModifyPower"`
	SaveDb                  bool   `json:"saveDb"`
	SortId                  string `json:"sortId"`
	Type                    int    `json:"type"`
}
