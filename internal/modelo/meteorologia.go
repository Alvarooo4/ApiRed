package modelo

import (
	"errors"
	"fmt"
	"time"
)

type Meteorologia struct {
	coordenada     Coordenada
	fecha          time.Time
	aguaLlovida    float64
	temperaturaMin float64
	temperaturaMax float64
}

func NuevaMeteorologia(coord Coordenada, fechaIntroducida time.Time, agua float64, tempMin float64, tempMax float64) (Meteorologia, error) {
	if agua < 0 {
		return Meteorologia{}, errors.New("el agua llovida no puede ser negativa")
	}

	if tempMin > tempMax {
		return Meteorologia{}, fmt.Errorf("la temperatura mínima (%f) no puede ser mayor que la máxima (%f)", tempMin, tempMax)
	}

	if tempMin < -273.15 {
		return Meteorologia{}, fmt.Errorf("la temperatura mínima (%f) debe ser físicamente posible", tempMin)
	}

	fecha00 := time.Date(fechaIntroducida.Year(), fechaIntroducida.Month(), fechaIntroducida.Day(), 0, 0, 0, 0, time.UTC)

	return Meteorologia{coordenada: coord, fecha: fecha00, aguaLlovida: agua, temperaturaMin: tempMin, temperaturaMax: tempMax}, nil
}
