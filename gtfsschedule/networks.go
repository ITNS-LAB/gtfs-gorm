package gtfsschedule

type Network struct {
	NetworkID                     string           `gorm:"primary_key"` // ユニーク ID: networks.txt 内で一意
	NetworkName                   string           `gorm:"not null"`    // ネットワークの名前
	FareLeg                       FareLeg          `gorm:"foreignKey:NetworkId;references:NetworkId "`
	FareLegJoinRulesFromNetworkID FareLegJoinRules `gorm:"foreignKey:NetworkId;references:FromNetworkID "`
	FareLegJoinRulesToNetworkID   FareLegJoinRules `gorm:"foreignKey:NetworkId;references:ToNetworkID "`
	RouteNetwork                  []RouteNetwork   `gorm:"foreignKey:NetworkId;references:NetworkID "`
}
