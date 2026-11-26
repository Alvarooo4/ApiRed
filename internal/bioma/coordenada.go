package bioma

import "errors"

var (
	ErrLatitudFueraDeRango  = errors.New("la latitud debe estar entre -90 y 90")
	ErrLongitudFueraDeRango = errors.New("la longitud debe estar entre -180 y 180")
)

type CoordenadaDecimal struct {
	latitudY  float64
	longitudX float64
}

func NuevaCoordenadaDecimal(latitud, longitud float64) (CoordenadaDecimal, error) {
	if latitud < -90 || latitud > 90 {
		return CoordenadaDecimal{}, ErrLatitudFueraDeRango
	}
	if longitud < -180 || longitud > 180 {
		return CoordenadaDecimal{}, ErrLongitudFueraDeRango
	}
	return CoordenadaDecimal{latitudY: latitud, longitudX: longitud}, nil
}
