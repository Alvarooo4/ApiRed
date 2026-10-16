package modelo

import (
	"errors"
	"fmt"
	"time"
)

type Meteorologia struct {
	fecha          int
	aguaLlovida    float64
	temperaturaMin float64
	temperaturaMax float64
}

func NuevaMeteorologia(fechaIntroducida time.Time, agua float64, tempMin float64, tempMax float64) (Meteorologia, error) {
	if agua < 0 {
		return Meteorologia{}, errors.New("el agua llovida no puede ser negativa")
	}

	if tempMin > tempMax {
		return Meteorologia{}, fmt.Errorf("la temperatura mínima (%f) no puede ser mayor que la máxima (%f)", tempMin, tempMax)
	}

	if tempMin < -273.15 {
		return Meteorologia{}, fmt.Errorf("la temperatura mínima (%f) debe ser físicamente posible", tempMin)
	}

	fecha00 := fechaIntroducida.Year()*10000 + int(fechaIntroducida.Month())*100 + fechaIntroducida.Day()

	return Meteorologia{fecha: fecha00, aguaLlovida: agua, temperaturaMin: tempMin, temperaturaMax: tempMax}, nil
}
