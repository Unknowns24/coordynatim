# Gominatim - Go library to access nominatim geocoding services

[![Latest Release](https://img.shields.io/github/release/unknowns24/gominatim.svg)](https://github.com/unknowns24/gominatim/releases)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/unknowns24/gominatim)
[![Coverage Status](https://coveralls.io/repos/github/unknowns24/gominatim/badge.svg?branch=master)](https://coveralls.io/github/unknowns24/gominatim?branch=master)


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
	geocoder, err := gominatim.NewGominatim(gominatim.DefaultConfig())
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := geocoder.Search(gominatim.SearchParameters{Street: "Falcon 357", Country: "Argentina", City: "San Nicolas de los Arroyos", PostalCode: "2900"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// DO SOMETHING WITH RES..
}
```
