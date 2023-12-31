package model

type CommonFlowNetwork struct {
	CommonUUID
	CommonSyncUUID
	CommonNameUnique
	CommonDescription
	GlobalUUID              string  `json:"global_uuid,omitempty" gorm:"unique"`
	ClientId                string  `json:"client_id,omitempty"`
	ClientName              string  `json:"client_name,omitempty"`
	SiteId                  string  `json:"site_id,omitempty"`
	SiteName                string  `json:"site_name,omitempty"`
	DeviceId                string  `json:"device_id,omitempty"`
	DeviceName              string  `json:"device_name,omitempty"`
	IsRemote                *bool   `json:"is_remote,omitempty"`
	FetchHistories          *bool   `json:"fetch_histories,omitempty"`
	FetchHistoriesFrequency int     `json:"fetch_hist_frequency,omitempty"`      // time example 15min
	DeleteHistoriesOnFetch  *bool   `json:"delete_histories_on_fetch,omitempty"` // drop the histories on the producer device on migration
	IsMasterSlave           *bool   `json:"is_master_slave,omitempty"`
	FlowHTTPS               *bool   `json:"flow_https,omitempty"`
	FlowUsername            *string `json:"flow_username,omitempty"`
	FlowPassword            *string `json:"flow_password,omitempty"`
	FlowToken               *string `json:"flow_token,omitempty"`
	IsTokenAuth             *bool   `json:"is_token_auth,omitempty"`
	IsError                 *bool   `json:"is_error" gorm:"default:false"`
	ErrorMsg                *string `json:"error_msg,omitempty"`
	CommonConnection
	CommonCreated
}

type FlowNetwork struct {
	CommonFlowNetwork
	FlowNetworkParentUUID string    `json:"flow_network_parent_uuid"`
	FlowIP                *string   `json:"flow_ip,omitempty" gorm:"uniqueIndex:idx_flow_networks_flow_ip_flow_port"`
	FlowPort              *int      `json:"flow_port,omitempty" gorm:"uniqueIndex:idx_flow_networks_flow_ip_flow_port"`
	FlowTokenLocal        *string   `json:"flow_token_local,omitempty"`
	FlowHTTSLocal         *bool     `json:"flow_https_local,omitempty"`
	FlowIPLocal           *string   `json:"flow_ip_local,omitempty"`
	FlowPortLocal         *int      `json:"flow_port_local,omitempty"`
	Streams               []*Stream `json:"streams" gorm:"many2many:flow_networks_streams;constraint:OnDelete:CASCADE"`
}

type FlowNetworkClone struct {
	CommonFlowNetwork
	CommonSourceUUID
	FlowIP       *string        `json:"flow_ip,omitempty" gorm:"uniqueIndex:idx_flow_network_clones_flow_ip_flow_port"`
	FlowPort     *int           `json:"flow_port,omitempty" gorm:"uniqueIndex:idx_flow_network_clones_flow_ip_flow_port"`
	StreamClones []*StreamClone `json:"stream_clones" gorm:"constraint:OnDelete:CASCADE"`
}
