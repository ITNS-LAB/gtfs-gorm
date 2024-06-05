package ormstatic

import "database/sql"

type Route struct {
	RouteId           string `gorm:"primaryKey;index;not null"`
	AgencyId          sql.NullString
	RouteShortName    sql.NullString
	RouteLongName     sql.NullString
	RouteDesc         sql.NullString
	RouteType         int16 `gorm:"index;not null"`
	RouteUrl          sql.NullString
	RouteColor        sql.NullString
	RouteTextColor    sql.NullString
	RouteSortOrder    sql.NullInt64 `gorm:"index"`
	ContinuousPickup  int16         `gorm:"default:1"`
	ContinuousDropOff int16         `gorm:"default:1"`
	Trip              Trip          `gorm:"foreignKey:RouteId"`
	FareRule          FareRule      `gorm:"foreignKey:RouteId"`
}

func (Route) TableName() string {
	return "routes"
}

func (r Route) GetRouteId() any {
	return r.RouteId
}

func (r Route) GetAgencyId() any {
	if r.AgencyId.Valid {
		return r.AgencyId.String
	}
	return nil
}

func (r Route) GetRouteShortName() any {
	if r.RouteShortName.Valid {
		return r.RouteShortName.String
	}
	return nil
}

func (r Route) GetRouteLongName() any {
	if r.RouteLongName.Valid {
		return r.RouteLongName.String
	}
	return nil
}

func (r Route) GetRouteDesc() any {
	if r.RouteDesc.Valid {
		return r.RouteDesc.String
	}
	return nil
}

func (r Route) GetRouteType() any {
	return r.RouteType
}

func (r Route) GetRouteUrl() any {
	if r.RouteUrl.Valid {
		return r.RouteUrl.String
	}
	return nil
}

func (r Route) GetRouteColor() any {
	if r.RouteColor.Valid {
		return r.RouteColor.String
	}
	return nil
}

func (r Route) GetRouteTextColor() any {
	if r.RouteTextColor.Valid {
		return r.RouteTextColor.String
	}
	return nil
}

func (r Route) GetRouteSortOrder() any {
	if r.RouteSortOrder.Valid {
		return r.RouteSortOrder.Int64
	}
	return nil
}

func (r Route) GetContinuousPickup() any {
	return r.ContinuousPickup
}

func (r Route) GetContinuousDropOff() any {
	return r.ContinuousDropOff
}
