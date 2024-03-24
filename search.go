package coordynatim

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"googlemaps.github.io/maps"
)

func (g *Coordynatim) Search(parameters SearchParameters) (GeoJSONResult, error) {
	var geoJSONResp GeoJSONResult

	respBytes, err := g.request(parameters)
	if err != nil {
		return geoJSONResp, err
	}

	err = json.Unmarshal(respBytes, &geoJSONResp)
	if err != nil {
		return geoJSONResp, err
	}

	return geoJSONResp, nil
}

func (g *Coordynatim) GetHouseAddressCoords(searchParams SearchParameters) (AddressCoords, error) {
	res, err := g.Search(searchParams)

	if err != nil {
		return AddressCoords{}, nil
	}

	addr := AddressCoords{}

	for _, f := range res.Features {
		if f.Properties.Geocoding.Type == "house" && len(f.Geometry.Coordinates) == 2 {
			addr.Lon = f.Geometry.Coordinates[0]
			addr.Lat = f.Geometry.Coordinates[1]
		}
	}

	if (addr.Lat == 0 || addr.Lon == 0) && g.config.GoogleMapsAPI == "" {
		return addr, errors.New("cannot obtain address cordinates")
	}

	if (addr.Lat == 0 || addr.Lon == 0) && g.config.GoogleMapsAPI != "" {
		results, err := g.googleClient.Geocode(context.Background(), &maps.GeocodingRequest{
			Address: searchParams.ToString(),
		})

		if err != nil {
			return addr, fmt.Errorf("cannot obtain address cordinates either from google, error: %s", err)
		}

		for _, result := range results {
			addr.Lon = Coordinate(result.Geometry.Location.Lng)
			addr.Lat = Coordinate(result.Geometry.Location.Lat)
			break
		}

		if addr.Lat == 0 || addr.Lon == 0 {
			return addr, fmt.Errorf("cannot obtain address cordinates either from google, error: %s", err)
		}
	}

	return addr, nil
}
