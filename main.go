package coordynatim

import (
	"fmt"
	"io"
	"net/http"

	"googlemaps.github.io/maps"
)

var defaultCfg = Config{
	UserAgent: "gominatim",
	Endpoint:  "https://nominatim.openstreetmap.org",
}

func DefaultConfig() Config {
	return defaultCfg
}

func DefaultConfigWithGoogleMapsAPI(GoogleMapsApiToken string) Config {
	config := defaultCfg
	config.GoogleMapsAPI = GoogleMapsApiToken

	return config
}

func NewCoordynatim(config Config) (*Coordynatim, error) {
	if config.Endpoint == "" {
		return nil, fmt.Errorf("endpoint must not be empty")
	}
	if config.UserAgent == "" {
		return nil, fmt.Errorf("userAgent must not be empty")
	}

	c := Coordynatim{
		config: config,
	}

	// Initialize maps client
	if config.GoogleMapsAPI != "" {
		mapsClient, err := maps.NewClient(maps.WithAPIKey(config.GoogleMapsAPI))
		if err != nil {
			return nil, fmt.Errorf("cannot initialize google maps api, error: %s", err)
		}

		c.googleClient = mapsClient
	}

	return &c, nil
}

func (g *Coordynatim) request(parameters SearchParameters) ([]byte, error) {
	requestURL := fmt.Sprintf("%s/search?%s&format=geocodejson&limit=1",
		g.config.Endpoint,
		parameters.ToQuery(),
	)

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", g.config.UserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}
