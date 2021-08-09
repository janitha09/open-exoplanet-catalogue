package oec

var OrhpanPlanetCount int = 0
var HottestStar Exoplanet
var Discoveries map[int]*SizeCounts = make(map[int]*SizeCounts)

type Exoplanet struct {
	PlanetIdentifier string
	HostStarTempK    float64
	DiscoveryYear    int
	RadiusJpt        float64
}

type SizeCounts struct {
	small  int
	medium int
	large  int
}

func NameOfThePlanetOrbittingTheHottestStar(d Exoplanet) {
	if HottestStar.HostStarTempK <= d.HostStarTempK {
		HottestStar = d
	}
}

func NumberOfOrphanPlanets(d Exoplanet) {
	if d.HostStarTempK == 0 {
		OrhpanPlanetCount++
	}
}

func DiscoveryByYearAndSize(d Exoplanet) {

	elem, ok := Discoveries[d.DiscoveryYear]
	if ok {
		incrementSize(d, elem)
	} else {
		initial := &SizeCounts{0, 0, 0}
		incrementSize(d, initial)
		Discoveries[d.DiscoveryYear] = initial

	}
}

func incrementSize(d Exoplanet, v *SizeCounts) {
	switch {
	case d.RadiusJpt < 1:
		v.small++
	case d.RadiusJpt < 2:
		v.medium++
	default:
		v.large++
	}
}
