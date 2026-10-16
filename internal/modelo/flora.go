package modelo

import "fmt"

type Flora struct {
	flores map[string]Flor
}

func NuevaFlora() *Flora { return &Flora{make(map[string]Flor)} }

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
	flor, err := nuevaFlor(nombre, agua, temperatura)
	if err == nil {
		flora.flores[flor.nombre] = flor
		return nil
	}
	return err
}
