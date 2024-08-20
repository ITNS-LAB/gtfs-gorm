package geomdatatypes

import (
	"encoding/hex"
	"fmt"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/encoding/ewkb"
	"github.com/paulmach/orb/encoding/wkt"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Geometry struct {
	Geom orb.Geometry
	Srid int
}

// GormDataType gormで使用するデータタイプを指定
func (g Geometry) GormDataType() string {
	return fmt.Sprintf("geometry(%s,%d)", g.GeometryType(), g.Srid)
}

// GormValue データベースに保存する際のデータ形式を指定
func (g Geometry) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	wkt_val := wkt.MarshalString(g.Geom)
	sql := fmt.Sprintf("ST_SetSRID(ST_GeomFromText(?), %d)", g.Srid)
	return clause.Expr{
		SQL:  sql,
		Vars: []interface{}{wkt_val},
	}
}

func (g *Geometry) Scan(v interface{}) error {
	if v == nil {
		g.Geom = nil
		return nil
	}

	// GEOMETRYデータをバイト配列に変換
	geomData, err := hex.DecodeString(v.(string))
	if err != nil {
		return err
	}

	// WKBからジオメトリデータをデコード
	geom, srid, err := ewkb.Unmarshal(geomData)
	if err != nil {
		return err
	}

	g.Geom = geom
	g.Srid = srid
	return nil
}

func (g Geometry) GeometryType() string {
	switch g.Geom.(type) {
	case orb.Point:
		return "POINT"
	case orb.MultiPoint:
		return "MULTIPOINT"
	case orb.LineString:
		return "LINESTRING"
	case orb.MultiLineString:
		return "MULTILINESTRING"
	case orb.Polygon:
		return "POLYGON"
	case orb.MultiPolygon:
		return "MULTIPOLYGON"
	case orb.Collection:
		return "GEOMETRYCOLLECTION"
	default:
		return "UNSUPPORTED GEOMETRY TYPE"
	}
}
