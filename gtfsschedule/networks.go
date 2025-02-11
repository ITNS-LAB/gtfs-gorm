package gtfsschedule

import (
	"fmt"
	"github.com/ITNS-LAB/gtfs-gorm/pkg/csvutil"
)

type Network struct {
	NetworkID   string `gorm:"primary_key"` // ユニーク ID: networks.txt 内で一意
	NetworkName string `gorm:"not null"`    // ネットワークの名前
	//FareLeg                       FareLeg          `gorm:"foreignKey:NetworkId;references:NetworkId "`
	FareLegJoinRulesFromNetworkID FareLegJoinRules `gorm:"foreignKey:NetworkId;references:FromNetworkID "`
	FareLegJoinRulesToNetworkID   FareLegJoinRules `gorm:"foreignKey:NetworkId;references:ToNetworkID "`
	RouteNetwork                  []RouteNetwork   `gorm:"foreignKey:NetworkId;references:NetworkID "`
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
			NetworkID:   networkID,
			NetworkName: networkName,
		})
	}

	return networks, nil
}
