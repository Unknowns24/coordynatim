# Coordynatim - Go library to get an address coords

[![Latest Release](https://img.shields.io/github/release/unknowns24/gominatim.svg)](https://github.com/unknowns24/coordynatim/releases)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/unknowns24/coordynatim)
[![Coverage Status](https://coveralls.io/repos/github/Unknowns24/gominatim/badge.svg?branch=main)](https://coveralls.io/github/Unknowns24/coordynatim?branch=main)

## Description

The purpose of this library is to provide developers with the necessary tools to easily find the coordinates behind an address.

> [!IMPORTANT]
> This library includes the implementation of nominatim API and Google Maps API.
> Why? Simple, nominatim is free but it cannot get all address cordinates so we complement the ones that nominatim cannot resolve with google maps geocode api.

> [!NOTE]
> At the date of writing this (24/03/2024) with the free cuota of google we have 40k free request to geocoding api.
> See the [Google Maps Pricing table](https://mapsplatform.google.com/pricing/) to know the actual price of geocoding api

## Examples

See this exaple to know how to start implementing this module in your application

> [!TIP]
> Access other exples code inside the examples folder.

```go
package main

import (
	"fmt"

	"github.com/unknowns24/coordynatim"
)

func main() {
	geocoder, err := coordynatim.NewCoordynatim(coordynatim.DefaultConfigWithGoogleMapsAPI("YOUR_API_KEY"))
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := geocoder.GetHouseAddressCoords(coordynatim.SearchParameters{Street: "Alvear 1053", Country: "Argentina", City: "San Nicolas de los Arroyos", Region: "Buenos Aires", PostalCode: "2900"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// DO SOMETHING WITH RES..
}
```

## Contributions

If you want to add anything, do it and submit a pull-request.
Please add Tests for your additions, otherwise it will be rejected!

## Usage of the Openstreetmaps-Nominatim Server

Please refer to the [Nominatim Wiki](http://wiki.openstreetmap.org/wiki/Nominatim)
if you plan to use the nominatim service of openstreetmaps. If you plan to generate
high loads with geoqueries, it would be nice if you did it on your own infrastructure, not on
their server.
