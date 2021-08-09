package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/open-exoplanet-catalogue/oec"
)

func main() {

	var ex oec.Exoplanet
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get("https://gist.githubusercontent.com/joelbirchler/66cf8045fcbb6515557347c05d789b4a/raw/9a196385b44d4288431eef74896c0512bad3defe/exoplanets")
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(resp.Body)
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	for dec.More() {
		err = dec.Decode(&ex)
		oec.NumberOfOrphanPlanets(ex)
		oec.NameOfThePlanetOrbittingTheHottestStar(ex)
		oec.DiscoveryByYearAndSize(ex)
	}

	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler)
	log.Printf("HTTP server will be listening on :8080. e.g. Use curl -s localhost:8080 to see results")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Number of Orphan Planets i.e. if HostStarTempK is 0: %v\n", oec.OrhpanPlanetCount)
	fmt.Fprintf(w, "The Planet with the hottest star %#v\n", oec.HottestStar.PlanetIdentifier)
	keys := make([]int, 0, len(oec.Discoveries))
	for key := range oec.Discoveries {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	fmt.Fprintf(w, "Planets discovered by year classified by Jupiter radii - small < 1, medium < 2, the rest large\n")

	for _, k := range keys {
		fmt.Fprintf(w, "key: %v, value: %#v\n", k, *oec.Discoveries[k])
	}
}
