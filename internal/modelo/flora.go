package modelo

import "fmt"

type Flora struct {
	flores map[string]Flor
}

func NuevaFlora() *Flora { return &Flora{flores: make(map[string]Flor)} }

func (flora *Flora) AgregarFlor(flor Flor) error {
	nombre := flor.nombre
	_, existe := flora.flores[nombre]
	if existe {
		return fmt.Errorf("la flor %s ya existe", nombre)
	}

	flora.flores[nombre] = flor
	return nil
}

func (flora *Flora) ModificarFlor(nombre string, agua float64, temperatura float64) error {
	flor, err := NuevaFlor(nombre, agua, temperatura)
	if err != nil {
		return err
	}

	flora.flores[flor.nombre] = flor
	return nil
}
