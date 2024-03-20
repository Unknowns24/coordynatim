package gominatim_test

import (
	"fmt"
	"testing"

	"github.com/unknowns24/gominatim"
)

func TestSearch(t *testing.T) {
	geocoder, err := gominatim.NewGominatim(gominatim.DefaultConfig())
	if err != nil {
		t.Error(err)
	}

	res, err := geocoder.Search(gominatim.SearchParameters{Country: "Germany", City: "Hamburg"})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(res)
}
