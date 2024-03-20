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

	res, err := geocoder.Search(gominatim.SearchParameters{Street: "Almirante Brown 79", Country: "Argentina", City: "San Nicolas de los Arroyos", PostalCode: "2900"})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, feature := range res.Features {
		fmt.Println(feature.Geometry.Coordinates)
	}
}
