package gominatim

import (
	"strings"
)

var (
	server string
)

type Address struct {
	House       string `json:"house_number"`
	Road        string `json:"road"`
	Village     string `json:"village"`
	Suburb      string `json:"suburb"`
	Town        string `json:"town"`
	City        string `json:"city"`
	State       string `json:"state"`
	County      string `json:"state_district"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

func SetServer(srv string) {
	srv = strings.TrimRight(srv, "/")
	server = srv
}
