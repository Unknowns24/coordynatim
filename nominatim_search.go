package gominatim

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type SearchResultError struct {
	message string
}

// Implement the Error() method for CustomError
func (e SearchResultError) Error() string {
	return e.message
}

type SearchResult struct {
	PlaceId       int64      `json:"place_id"`
	License       string     `json:"license"`
	OsmType       string     `json:"osm_type"`
	OsmId         int64      `json:"osm_id"`
	Boundingbox   []string   `json:"boundingbox"`
	Polygonpoints [][]string `json:"polygonpoints"`
	Lat           string     `json:"lat"`
	Lon           string     `json:"lon"`
	DisplayName   string     `json:"display_name"`
	Class         string     `json:"class"`
	Type          string     `json:"type"`
	Address       Address    `json:"address"`
}

type SearchQuery struct {
	JsonCallback    interface{}
	AcceptLanguage  string
	Q               string
	Street          string
	City            string
	County          string
	State           string
	Postalcode      string
	Countrycodes    []string
	Viewbox         string
	Bounded         bool
	Polygon         bool
	Addressdetails  bool
	Email           string
	ExcludePlaceIds []string
	Limit           int
	PolygonGeojson  bool
	PolygonKml      bool
	PolygonText     bool
	PolygonSvg      bool
}

func (q *SearchQuery) specificFieldsUsed() bool {
	return q.Street != "" || q.City != "" || q.County != "" || q.State != "" || q.Postalcode != ""
}

func (q *SearchQuery) buildQuery() (string, error) {
	if server == "" {
		return "", errors.New("server is not set. Set via gominatim.SetServer(srv string)")
	}
	s := server
	s = s + "/search?format=json"
	if q.JsonCallback != nil {
		cb, err := json.Marshal(q.JsonCallback)
		if err != nil {
			return "", err
		}
		s += "&json_callback=" + string(cb)
	}
	if q.AcceptLanguage != "" {
		s += "&accept-language=" + url.QueryEscape(q.AcceptLanguage)
	}
	if q.Q != "" {
		s += "&q=" + url.QueryEscape(q.Q)
	} else {
		if q.specificFieldsUsed() {
			if q.Street != "" {
				s += "&street=" + url.QueryEscape(q.Street)
			}
			if q.City != "" {
				s += "&city=" + url.QueryEscape(q.City)
			}
			if q.County != "" {
				s += "&county=" + url.QueryEscape(q.County)
			}
			if q.State != "" {
				s += "&state=" + url.QueryEscape(q.State)
			}
			if q.Postalcode != "" {
				s += "&postalcode=" + url.QueryEscape(q.Postalcode)
			}
		} else {
			return "", errors.New("you must use either Q or one or more of Street, City, County, State, Postalcode. The latter will be ignored if the further is used")
		}
	}
	if q.Countrycodes != nil && len(q.Countrycodes) > 0 {
		als := ""
		first := true
		for i := range q.Countrycodes {
			if !first {
				als = als + ","
			}
			als = als + q.Countrycodes[i]
			if first {
				first = false
			}
		}
		s += "&countrycodes=" + url.QueryEscape(als)
	}
	if q.Viewbox != "" {
		s += "&viewbox=" + url.QueryEscape(q.Viewbox)
	}
	if q.Bounded {
		s += "&bounded=1"
	} else {
		s += "&bounded=0"
	}
	if q.Polygon {
		s += "&polygon=1"
	} else {
		s += "&polygon=0"
	}
	if q.Addressdetails {
		s += "&addressdetails=1"
	} else {
		s += "&addressdetails=0"
	}
	if q.Email != "" {
		s += "&email=" + url.QueryEscape(q.Email)
	}
	if q.ExcludePlaceIds != nil && len(q.ExcludePlaceIds) > 0 {
		als := ""
		first := true
		for i := range q.ExcludePlaceIds {
			if !first {
				als = als + ","
			}
			als = als + q.ExcludePlaceIds[i]
			if first {
				first = false
			}
		}
		s += "&exclude_place_ids=" + url.QueryEscape(als)
	}
	if q.Limit > 0 {
		s += "&limit=" + strconv.Itoa(q.Limit)
	}
	if q.PolygonGeojson {
		s += "&polygon_geojson=1"
	} else {
		s += "&polygon_geojson=0"
	}
	if q.PolygonKml {
		s += "&polygon_kml=1"
	} else {
		s += "&polygon_kml=0"
	}
	if q.PolygonSvg {
		s += "&polygon_svg=1"
	} else {
		s += "&polygon_svg=0"
	}
	if q.PolygonText {
		s += "&polygon_text=1"
	} else {
		s += "&polygon_text=0"
	}
	return s, nil
}

func (q *SearchQuery) Get() ([]SearchResult, error) {
	querystring, err := q.buildQuery()
	if err != nil {
		return nil, err
	}
	resp, err := http.Get(querystring)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := make([]SearchResult, 0)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, SearchResultError{err.Error()}
	}
	if len(result) == 0 {
		return nil, errors.New("no results;")
	}
	return result, nil
}
