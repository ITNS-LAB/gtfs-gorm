package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Network struct {
	NetworkId   string `gorm:"primary_key"` // ユニーク ID: networks.txt 内で一意
	NetworkName string `gorm:"not null"`    // ネットワークの名前
	//FareLeg                       FareLeg          `gorm:"foreignKey:NetworkId;references:NetworkId"`
	FareLegJoinRulesFromNetworkID FareLegJoinRules `gorm:"foreignKey:FromNetworkId;references:NetworkId"`
	FareLegJoinRulesToNetworkID   FareLegJoinRules `gorm:"foreignKey:ToNetworkId;references:NetworkId"`
	RouteNetwork                  []RouteNetwork   `gorm:"foreignKey:NetworkId;references:NetworkId"`
}

func (Network) TableName() string {
	return "network"
}

func ParseNetwork(path string) ([]Network, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open network CSV: %w", err)
	}

	// Parse the data and create a slice of Network structs
	var networks []Network
	for i := 0; i < len(df.Records); i++ {
		networkID, err := df.GetString(i, "network_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'network_id' at row %d: %w", i, err)
		}

		networkName, err := df.GetString(i, "network_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'network_name' at row %d: %w", i, err)
		}

		// Create the Network struct and append to the list
		networks = append(networks, Network{
			NetworkId:   networkID,
			NetworkName: networkName,
		})
	}

	return networks, nil
}

type NetworkGeom struct {
	NetworkId   string `gorm:"primary_key"` // ユニーク ID: networks.txt 内で一意
	NetworkName string `gorm:"not null"`    // ネットワークの名前
	//FareLeg                       FareLeg          `gorm:"foreignKey:NetworkId;references:NetworkId"`
	FareLegJoinRulesFromNetworkID FareLegJoinRulesGeom `gorm:"foreignKey:FromNetworkId;references:NetworkId"`
	FareLegJoinRulesToNetworkID   FareLegJoinRulesGeom `gorm:"foreignKey:ToNetworkId;references:NetworkId"`
	RouteNetwork                  []RouteNetworkGeom   `gorm:"foreignKey:NetworkId;references:NetworkId"`
}

func (NetworkGeom) TableName() string {
	return "network"
}

func ParseNetworkGeom(path string) ([]NetworkGeom, error) {
	// Open the CSV file
	df, err := csvutil.OpenCSV(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open network CSV: %w", err)
	}

	// Parse the data and create a slice of Network structs
	var networks []NetworkGeom
	for i := 0; i < len(df.Records); i++ {
		networkID, err := df.GetString(i, "network_id")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'network_id' at row %d: %w", i, err)
		}

		networkName, err := df.GetString(i, "network_name")
		if err != nil {
			return nil, fmt.Errorf("failed to get 'network_name' at row %d: %w", i, err)
		}

		// Create the Network struct and append to the list
		networks = append(networks, NetworkGeom{
			NetworkId:   networkID,
			NetworkName: networkName,
		})
	}

	return networks, nil
}
