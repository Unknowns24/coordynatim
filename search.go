package gominatim

import (
	"encoding/json"
)

func (g *Gominatim) Search(parameters SearchParameters) (GeoJSONResult, error) {
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
