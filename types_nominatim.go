package coordynatim

import (
	"fmt"
	"net/url"
	"strings"
)

type (
	GeoJSONFeature struct {
		Type       string `json:"type"`
		Properties struct {
			Geocoding struct {
				PlaceID int    `json:"place_id"`
				OsmType string `json:"osm_type"`
				OsmID   int    `json:"osm_id"`
				Type    string `json:"type"`
				Name    string `json:"name"`
			} `json:"geocoding"`
		} `json:"properties"`
		Geometry struct {
			Type        string       `json:"type"`
			Coordinates []Coordinate `json:"coordinates"`
		} `json:"geometry"`
	}

	GeoJSONResult struct {
		Type      string `json:"type"`
		Geocoding struct {
			Version     string `json:"version"`
			Attribution string `json:"attribution"`
			Licence     string `json:"licence"`
			Query       string `json:"query"`
		} `json:"geocoding"`
		Features []GeoJSONFeature `json:"features"`
	}

	SearchParameters struct {
		Q          string
		City       string
		Street     string
		Region     string
		Country    string
		PostalCode string
	}
)

func (s *SearchParameters) ToQuery() string {

	queryURLParts := []string{}

	if s.Q != "" {
		queryURLParts = append(queryURLParts, fmt.Sprintf("q=%s", url.QueryEscape(s.Q)))
	} else {
		if s.City != "" {
			queryURLParts = append(queryURLParts, fmt.Sprintf("city=%s", url.QueryEscape(s.City)))
		}
		if s.Street != "" {
			queryURLParts = append(queryURLParts, fmt.Sprintf("street=%s", url.QueryEscape(s.Street)))
		}
		if s.Country != "" {
			queryURLParts = append(queryURLParts, fmt.Sprintf("country=%s", url.QueryEscape(s.Country)))
		}
		if s.PostalCode != "" {
			queryURLParts = append(queryURLParts, fmt.Sprintf("postalcode=%s", url.QueryEscape(s.PostalCode)))
		}
	}

	return strings.Join(queryURLParts, "&")

}

func (s *SearchParameters) ToString() string {

	queryURLParts := []string{}

	if s.Q != "" {
		queryURLParts = append(queryURLParts, s.Q)
	} else {
		if s.Street != "" {
			queryURLParts = append(queryURLParts, s.Street)
		}
		if s.City != "" {
			queryURLParts = append(queryURLParts, s.City)
		}
		if s.Region != "" {
			queryURLParts = append(queryURLParts, s.Region)
		}
		if s.Country != "" {
			queryURLParts = append(queryURLParts, s.Country)
		}
	}

	return strings.Join(queryURLParts, ", ")

}
