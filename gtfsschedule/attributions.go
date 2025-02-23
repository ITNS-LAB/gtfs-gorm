package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Attribution struct {
	AttributionId    *string
	AgencyId         *string
	RouteId          *string
	TripId           *string
	OrganizationName string `gorm:"not null"`
	IsProducer       *int
	IsOperator       *int
	IsAuthority      *int
	AttributionURL   *string
	AttributionEmail *string
	AttributionPhone *string
}

func ParseAttribution(path string) ([]Attribution, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open Attribution CSV: %w", err)
	}

	// データを解析して Attribution 構造体のスライスを作成
	var attributions []Attribution
	for i := 0; i < len(df.Records); i++ {
		attributionID, _ := df.GetStringPtr(i, "attribution_id") // 必須でない場合はエラーを無視
		agencyID, _ := df.GetStringPtr(i, "agency_id")
		routeID, _ := df.GetStringPtr(i, "route_id")
		tripID, _ := df.GetStringPtr(i, "trip_id")

		organizationName, err := df.GetString(i, "organization_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'organization_name' at row %d: %w", i, err)
		}

		isProducer, err := df.GetIntPtr(i, "is_producer")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'is_producer' at row %d: %w", i, err)
		}

		isOperator, err := df.GetIntPtr(i, "is_operator")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'is_operator' at row %d: %w", i, err)
		}

		isAuthority, err := df.GetIntPtr(i, "is_authority")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'is_authority' at row %d: %w", i, err)
		}

		attributionURL, _ := df.GetStringPtr(i, "attribution_url")
		attributionEmail, _ := df.GetStringPtr(i, "attribution_email")
		attributionPhone, _ := df.GetStringPtr(i, "attribution_phone")

		// Attribution 構造体を作成しリストに追加
		attributions = append(attributions, Attribution{
			AttributionId:    attributionID,
			AgencyId:         agencyID,
			RouteId:          routeID,
			TripId:           tripID,
			OrganizationName: organizationName,
			IsProducer:       isProducer,
			IsOperator:       isOperator,
			IsAuthority:      isAuthority,
			AttributionURL:   attributionURL,
			AttributionEmail: attributionEmail,
			AttributionPhone: attributionPhone,
		})
	}

	return attributions, nil
}

type AttributionGeom struct {
	AttributionId    *string
	AgencyId         *string
	RouteId          *string
	TripId           *string
	OrganizationName string `gorm:"not null"`
	IsProducer       *int
	IsOperator       *int
	IsAuthority      *int
	AttributionURL   *string
	AttributionEmail *string
	AttributionPhone *string
}

func ParseAttributionGeom(path string) ([]AttributionGeom, error) {
	// CSVを開く
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open Attribution CSV: %w", err)
	}

	// データを解析して Attribution 構造体のスライスを作成
	var attributions []AttributionGeom
	for i := 0; i < len(df.Records); i++ {
		attributionID, _ := df.GetStringPtr(i, "attribution_id") // 必須でない場合はエラーを無視
		agencyID, _ := df.GetStringPtr(i, "agency_id")
		routeID, _ := df.GetStringPtr(i, "route_id")
		tripID, _ := df.GetStringPtr(i, "trip_id")

		organizationName, err := df.GetString(i, "organization_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'organization_name' at row %d: %w", i, err)
		}

		isProducer, err := df.GetIntPtr(i, "is_producer")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'is_producer' at row %d: %w", i, err)
		}

		isOperator, err := df.GetIntPtr(i, "is_operator")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'is_operator' at row %d: %w", i, err)
		}

		isAuthority, err := df.GetIntPtr(i, "is_authority")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'is_authority' at row %d: %w", i, err)
		}

		attributionURL, _ := df.GetStringPtr(i, "attribution_url")
		attributionEmail, _ := df.GetStringPtr(i, "attribution_email")
		attributionPhone, _ := df.GetStringPtr(i, "attribution_phone")

		// Attribution 構造体を作成しリストに追加
		attributions = append(attributions, AttributionGeom{
			AttributionId:    attributionID,
			AgencyId:         agencyID,
			RouteId:          routeID,
			TripId:           tripID,
			OrganizationName: organizationName,
			IsProducer:       isProducer,
			IsOperator:       isOperator,
			IsAuthority:      isAuthority,
			AttributionURL:   attributionURL,
			AttributionEmail: attributionEmail,
			AttributionPhone: attributionPhone,
		})
	}

	return attributions, nil
}
