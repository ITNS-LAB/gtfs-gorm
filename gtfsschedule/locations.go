package gtfsschedule

type GeoJSON struct {
	Type     string    `json:"type"`     // "FeatureCollection"
	Features []Feature `json:"features"` // Featureオブジェクトのコレクション
}

type Feature struct {
	Type       string      `json:"type"`       // "Feature"
	ID         string      `json:"id"`         // 場所を識別
	Properties Properties  `json:"properties"` // 場所のプロパティ
	Geometry   Geometry    `json:"geometry"`   // 場所のジオメトリ
	StopTimes  []StopTimes `gorm:"foreignKey:LocationId;references:LocationId "`
}

type Properties struct {
	StopName string `json:"stop_name,omitempty"` // 乗客に表示される場所の名前
	StopDesc string `json:"stop_desc,omitempty"` // 乗客の方向を示す説明
}

type Geometry struct {
	Type        string      `json:"type"`        // "Polygon" または "MultiPolygon"
	Coordinates [][]float64 `json:"coordinates"` // 地理座標 (緯度、経度)
}
