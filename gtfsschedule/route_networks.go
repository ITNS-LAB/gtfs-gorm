package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type RouteNetwork struct {
	NetworkId string `gorm:"primary_key"` // networks.network_id を参照する外部 ID
	RouteId   string `gorm:"not null"`
}

func ParseRouteNetwork(path string) ([]RouteNetwork, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open route_network CSV: %w", err)
	}

	// Parse the data and create a slice of RouteNetwork structs
	var routeNetworks []RouteNetwork
	for i := 0; i < len(df.Records); i++ {
		networkID, err := df.GetString(i, "network_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'network_id' at row %d: %w", i, err)
		}

		routeID, err := df.GetString(i, "route_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_id' at row %d: %w", i, err)
		}

		// Create the RouteNetwork struct and append to the list
		routeNetworks = append(routeNetworks, RouteNetwork{
			NetworkId: networkID,
			RouteId:   routeID,
		})
	}

	return routeNetworks, nil
}

type RouteNetworkGeom struct {
	NetworkId string `gorm:"primary_key"` // networks.network_id を参照する外部 ID
	RouteId   string `gorm:"not null"`
}

func ParseRouteNetworkGeom(path string) ([]RouteNetworkGeom, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open route_network CSV: %w", err)
	}

	// Parse the data and create a slice of RouteNetwork structs
	var routeNetworks []RouteNetworkGeom
	for i := 0; i < len(df.Records); i++ {
		networkID, err := df.GetString(i, "network_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'network_id' at row %d: %w", i, err)
		}

		routeID, err := df.GetString(i, "route_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'route_id' at row %d: %w", i, err)
		}

		// Create the RouteNetwork struct and append to the list
		routeNetworks = append(routeNetworks, RouteNetworkGeom{
			NetworkId: networkID,
			RouteId:   routeID,
		})
	}

	return routeNetworks, nil
}
