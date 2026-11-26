package bioma

import (
	"errors"
	"testing"
)

func TestCoordenadaValida(t *testing.T) {
	coordenada, err := NuevaCoordenadaDecimal(38.79, -4.87)

	if err != nil {
		t.Fatalf("no esperaba error: %v", err)
	}
	if coordenada.latitudY != 38.79 || coordenada.longitudX != -4.87 {
		t.Errorf("coordenada incorrecta: %+v", coordenada)
	}
}

func TestLatitudFueraDeRango(t *testing.T) {
	_, err := NuevaCoordenadaDecimal(91, 0)

	if !errors.Is(err, ErrLatitudFueraDeRango) {
		t.Errorf("esperaba ErrLatitudFueraDeRango, obtuve %v", err)
	}
}

func TestLongitudFueraDeRango(t *testing.T) {
	_, err := NuevaCoordenadaDecimal(0, -181)

	if !errors.Is(err, ErrLongitudFueraDeRango) {
		t.Errorf("esperaba ErrLongitudFueraDeRango, obtuve %v", err)
	}
}