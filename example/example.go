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

	fmt.Println(res)
}
