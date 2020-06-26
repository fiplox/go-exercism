package space

type Planet string

func Age(sec float64, planet Planet) float64 {
	switch planet {
	case "Earth":
		return (sec / 31557600)
	case "Mercury":
		return (sec / 31557600 / 0.2408467)
	case "Venus":
		return (sec / 31557600 / 0.61519726)
	case "Mars":
		return (sec / 31557600 / 1.8808158)
	case "Jupiter":
		return (sec / 31557600 / 11.862615)
	case "Saturn":
		return (sec / 31557600 / 29.447498)
	case "Uranus":
		return (sec / 31557600 / 84.016846)
	case "Neptune":
		return (sec / 31557600 / 164.79132)
	default:
		break
	}
	return 0.0
}
