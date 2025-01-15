package gtfsschedule

type RouteNetwork struct {
	NetworkID string `gorm:"primary_key"` // networks.network_id を参照する外部 ID
	RouteID   string `gorm:"not null"`
}
