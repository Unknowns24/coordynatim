package coordynatim_test

import (
	"testing"

	"github.com/unknowns24/coordynatim"
)

func TestSearchWithGoogleMapsAPI(t *testing.T) {
	geocoder, err := coordynatim.NewCoordynatim(coordynatim.DefaultConfigWithGoogleMapsAPI("SOME_API_KEY"))
	if err != nil {
		t.Error("something wron is happening, this should fail")
	}

	_, err = geocoder.Search(coordynatim.SearchParameters{Country: "Germany", City: "Hamburg"})
	if err != nil {
		t.Error(err)
	}
}

func TestSearch(t *testing.T) {
	geocoder, err := coordynatim.NewCoordynatim(coordynatim.DefaultConfig())
	if err != nil {
		t.Error(err)
	}

	_, err = geocoder.Search(coordynatim.SearchParameters{Q: "Hamburg, Germany"})
	if err != nil {
		t.Error(err)
	}
}

func TestGetHouseAddressCoords(t *testing.T) {
	geocoder, err := coordynatim.NewCoordynatim(coordynatim.DefaultConfig())
	if err != nil {
		t.Error(err)
	}

	_, err = geocoder.GetHouseAddressCoords(coordynatim.SearchParameters{Street: "Ameghino 120", City: "San Nicolas de los Arroyos", PostalCode: "2900", Country: "Argentina"})
	if err != nil {
		t.Error(err)
	}
}

func TestSearchParametersToString(t *testing.T) {
	sp := coordynatim.SearchParameters{Street: "Ameghino 120", City: "San Nicolas de los Arroyos", PostalCode: "2900", Country: "Argentina"}
	expected := "Ameghino 120, San Nicolas de los Arroyos, Argentina"
	if sp.ToString() != expected {
		t.Errorf("Unexpected to string answer, %s expected. %s received", expected, sp.ToString())
	}
}
