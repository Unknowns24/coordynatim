package coordynatim

import (
	"fmt"

	"googlemaps.github.io/maps"
)

type (
	Coordynatim struct {
		config       Config
		googleClient *maps.Client
	}

	Config struct {
		UserAgent string
		Endpoint  string

		GoogleMapsAPI string
	}

	AddressCoords struct {
		Lat Coordinate `json:"lat"`
		Lon Coordinate `json:"lon"`
	}

	Coordinate float64
)

func (c *Coordinate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.5f", float64(*c))), nil
}
