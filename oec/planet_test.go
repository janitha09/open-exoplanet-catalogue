package oec

import (
	"encoding/json"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJsonStreaming(t *testing.T) {

	var ex Exoplanet
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get("https://gist.githubusercontent.com/joelbirchler/66cf8045fcbb6515557347c05d789b4a/raw/9a196385b44d4288431eef74896c0512bad3defe/exoplanets")
	if err != nil {
		t.Fatal(err)
	}
	dec := json.NewDecoder(resp.Body)
	_, err = dec.Token()
	if err != nil {
		t.Fatal(err)
	}

	for dec.More() {
		err = dec.Decode(&ex)
		// t.Log(ex)
		NumberOfOrphanPlanets(ex)
		NameOfThePlanetOrbittingTheHottestStar(ex)
		DiscoveryByYearAndSize(ex)
		// NameOfThePlanetOrbittingTheHottestStar(ex)
	}

	_, err = dec.Token()
	if err != nil {
		t.Fatal(err)
	}
	// t.Log(ex)
	t.Log(OrhpanPlanetCount)
	t.Log(HottestStar)
	t.Logf("%#v", Discoveries)
}

func TestNameOfThePlanetOrbittingTheHottestStar(t *testing.T) {
	testCases := []struct {
		data     []Exoplanet
		expected string
	}{
		{
			data: []Exoplanet{{
				PlanetIdentifier: "a",
				HostStarTempK:    0,
				DiscoveryYear:    2020,
				RadiusJpt:        1,
			}},
			expected: "a",
		},
		{
			data: []Exoplanet{
				{
					PlanetIdentifier: "a",
					HostStarTempK:    1,
					DiscoveryYear:    2020,
					RadiusJpt:        1,
				},
				{
					PlanetIdentifier: "b",
					HostStarTempK:    2,
					DiscoveryYear:    2020,
					RadiusJpt:        1,
				},
			},
			expected: "b",
		},
	}
	for i, tc := range testCases {
		HottestStar = Exoplanet{}
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			for _, d := range tc.data {
				NameOfThePlanetOrbittingTheHottestStar(d)
			}
			assert.Equal(t, tc.expected, HottestStar.PlanetIdentifier)
		})
	}
}

func TestOrphanPlanetCount(t *testing.T) {
	testCases := []struct {
		data     []Exoplanet
		expected int
	}{
		{
			data: []Exoplanet{{
				PlanetIdentifier: "a",
				HostStarTempK:    0,
				DiscoveryYear:    2020,
				RadiusJpt:        1,
			}},
			expected: 1,
		},
		{
			data: []Exoplanet{
				{
					PlanetIdentifier: "a",
					HostStarTempK:    0,
					DiscoveryYear:    2020,
					RadiusJpt:        1,
				},
				{
					PlanetIdentifier: "b",
					HostStarTempK:    0,
					DiscoveryYear:    2020,
					RadiusJpt:        1,
				},
			},
			expected: 2,
		},
	}
	for i, tc := range testCases {
		OrhpanPlanetCount = 0
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			for _, d := range tc.data {
				NumberOfOrphanPlanets(d)
			}
			assert.Equal(t, tc.expected, OrhpanPlanetCount)
		})
	}
}

func TestNumberOfPlanetsDiscoveredPerYear(t *testing.T) {
	testCases := []struct {
		data     []Exoplanet
		expected map[int]*SizeCounts
	}{
		{
			data: []Exoplanet{{
				PlanetIdentifier: "a",
				HostStarTempK:    0,
				DiscoveryYear:    2020,
				RadiusJpt:        1,
			}},
			expected: map[int]*SizeCounts{2020: {
				small:  0,
				medium: 1,
				large:  0,
			}},
		},
		{
			data: []Exoplanet{
				{
					PlanetIdentifier: "a",
					HostStarTempK:    0,
					DiscoveryYear:    2020,
					RadiusJpt:        1,
				},
				{
					PlanetIdentifier: "b",
					HostStarTempK:    0,
					DiscoveryYear:    2020,
					RadiusJpt:        1,
				},
			},
			expected: map[int]*SizeCounts{2020: {
				small:  0,
				medium: 2,
				large:  0,
			}},
		},
		{
			data: []Exoplanet{
				{
					PlanetIdentifier: "a",
					HostStarTempK:    0,
					DiscoveryYear:    2019,
					RadiusJpt:        5,
				},
				{
					PlanetIdentifier: "b",
					HostStarTempK:    0,
					DiscoveryYear:    2020,
					RadiusJpt:        1,
				},
				{
					PlanetIdentifier: "b",
					HostStarTempK:    0,
					DiscoveryYear:    2020,
					RadiusJpt:        1,
				},
			},
			expected: map[int]*SizeCounts{
				2020: {
					small:  0,
					medium: 2,
					large:  0,
				},
				2019: {
					small:  0,
					medium: 0,
					large:  1,
				},
			},
		},
	}
	for i, tc := range testCases[0:] {
		Discoveries = make(map[int]*SizeCounts)
		Discoveries[2020] = &SizeCounts{0, 0, 0}
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			for _, d := range tc.data {
				DiscoveryByYearAndSize(d)
			}
			assert.Equal(t, tc.expected, Discoveries)
		})
	}
}
