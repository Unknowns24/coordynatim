# Gominatim - Go library to access nominatim geocoding services

[![Latest Release](https://img.shields.io/github/release/muesli/gominatim.svg)](https://github.com/unknowns24/gominatim/releases)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/unknowns24/gominatim)
[![Coverage Status](https://coveralls.io/repos/github/muesli/gominatim/badge.svg?branch=master)](https://coveralls.io/github/unknowns24/gominatim?branch=master)

## Geocoding? WTF?

If you want to determine the coordinates of a certain location by only having its
name, you can do this via a geocoding service. If you want to do this in Go, you
probably want to use gominatim to do it.

## Features

The plan is to cover everything this site documents:
[Nominatim Wiki](http://wiki.openstreetmap.org/wiki/Nominatim)

-   [x] Search
-   [x] Reverse Geocoding

## Contributions

If you want to add anything, do it and submit a pull-request.
Please add Tests for your additions!

## Usage of the Openstreetmaps-Nominatim Server

Please refer to the [Nominatim Wiki](http://wiki.openstreetmap.org/wiki/Nominatim)
if you plan to use the nominatim service of openstreetmaps. If you plan to generate
high loads with geoqueries, it would be nice if you did it on your own infrastructure, not on
their server.

## Examples

```go
package main

import (
	"fmt"
	"github.com/unknowns24/gominatim"
)

func main() {
	gominatim.SetServer("https://nominatim.openstreetmap.org/")

	//Get by a Querystring
	qry := gominatim.SearchQuery{
		Q: "Hamburg",
	}
	resp, _ := qry.Get() // Returns []gominatim.SearchResult
	fmt.Printf("Found location: %s (%s, %s)\n", resp[0].DisplayName, resp[0].Lat, resp[0].Lon)

	//Get by City
	qry = gominatim.SearchQuery{
		City: "Berlin",
	}
	resp, _ = qry.Get()
	fmt.Printf("Found location: %s (%s, %s)\n", resp[0].DisplayName, resp[0].Lat, resp[0].Lon)

	//Reverse Geocoding
	rqry := gominatim.ReverseQuery{
		Lat: "52.5170365",
		Lon: "13.3888599",
	}
	rresp, _ := rqry.Get()
	fmt.Printf("Found %s\n", rresp.DisplayName)
}
```
