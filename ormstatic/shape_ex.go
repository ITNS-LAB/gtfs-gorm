package ormstatic

import geomdatatypes "github.com/ITNS-LAB/gtfs-gorm/internal/gormdatatypes"

type ShapeEx struct {
	TripId            *string  `gorm:"primaryKey"`
	ShapeId           *string  `gorm:"primaryKey"`
	ShapePtLat        *float64 `gorm:"not null"`
	ShapePtLon        *float64 `gorm:"not null"`
	ShapePtSequence   *int     `gorm:"primaryKey"`
	ShapeDistTraveled *float64
	StopId            *string
	Geom              *geomdatatypes.Geometry `gorm:"index"`
	Stop              Stop                    `gorm:"foreignKey:StopId"`
}

func (ShapeEx) TableName() string {
	return "shapes_ex"
}

// CreateShapeEx より短い間隔のshapesを生成します。intervalは新たに作るshapesの間隔の値(m)です。
func CreateShapeEx(shapes []Shape, interval int) []ShapeEx {
	var shapesEx []ShapeEx
	shapesEx = append(shapesEx, ShapeEx{
		ShapeId:         shapes[0].ShapeId,
		ShapePtLat:      shapes[0].ShapePtLat,
		ShapePtLon:      shapes[0].ShapePtLon,
		ShapePtSequence: shapes[0].ShapePtSequence,
	})
	return shapesEx
}
