package util

import (
	"github.com/tidwall/geodesic"
)

// KarneyGrs80 GPSの座標から距離(m)を返します。
func KarneyGrs80(prevLat, prevLon, nextLat, nextLon float64) (distance float64) {
	geodesic.WGS84.Inverse(prevLat, prevLon, nextLat, nextLon, &distance, nil, nil)
	return distance
}

// KarneyWgs84 GPSの座標から距離(m)を返します。
func KarneyWgs84(prevLat, prevLon, nextLat, nextLon float64) (distance float64) {
	geodesic.WGS84.Inverse(prevLat, prevLon, nextLat, nextLon, &distance, nil, nil)
	return distance
}
