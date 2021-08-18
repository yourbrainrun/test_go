package models

type Channel struct {
	Id             int    `gorm:"column:id;primaryKey"`
	ProjectId      int    `gorm:"column:project_id"`
	AppId          string `gorm:"column:app_id"`
	StreamChannel  string `gorm:"column:stream_channel"`
	StreamUrl      string `gorm:"column:stream_url"`
	Status         int    `gorm:"column:status"`
	OrganizationId int    `gorm:"column:organization_id"`
	SdnUdpIp       string `gorm:"column:sdn_udp_ip"`
	SdnUdpPort     string `gorm:"column:sdn_udp_port"`
	ServerSign     string `gorm:"column:server_sign"`
}

func (Channel) TableName() string {
	return "cdn_channels"
}
