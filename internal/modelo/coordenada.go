package modelo

import "fmt"

type Coordenada struct {
	latitudY  float64
	longitudX float64
}

func NuevaCoordenada(lY, lX float64) (Coordenada, error) {
	if lY < -90 || lY > 90 {
		return Coordenada{}, fmt.Errorf("latitud del eje y fuera de rango: %f vs [-90, 90]", lY)
	}

	if lX < -180 || lX > 180 {
		return Coordenada{}, fmt.Errorf("longitud del eje x fuera de rango: %f vs [-180, 180]", lX)
	}

	return Coordenada{latitudY: lY, longitudX: lX}, nil
}
