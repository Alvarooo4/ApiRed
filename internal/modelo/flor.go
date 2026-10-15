package modelo

import "errors"

type Flor struct {
	nombre           string
	aguaIdeal        float64
	temperaturaIdeal float64
}

func NuevaFlor(nombreIntroducido string, agua float64, temperatura float64) (Flor, error) {
	if nombreIntroducido == "" {
		return Flor{}, errors.New("el nombre de la flor no puede ser nulo")
	}

	if agua <= 0 {
		return Flor{}, errors.New("el agua óptima no puede ser menor o igual a cero")
	}

	if temperatura < -50 || temperatura > 60 {
		return Flor{}, errors.New("la temperatura óptima debe estar en el rango [-50, 60]")
	}

	return Flor{nombre: nombreIntroducido, aguaIdeal: agua, temperaturaIdeal: temperatura}, nil
}
