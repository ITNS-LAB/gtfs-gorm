package gtfsschedule

type Network struct {
	NetworkID   string `gorm:"primary_key"` // ユニーク ID: networks.txt 内で一意
	NetworkName string `gorm:"not null"`    // ネットワークの名前
}
