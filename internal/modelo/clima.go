package modelo

import (
	"fmt"
	"time"
)

type Clima struct {
	coordenada            Coordenada
	registroMetereologico map[time.Time]Meteorologia
}

func NuevoClima(coord Coordenada) Clima {
	return Clima{coordenada: coord, registroMetereologico: make(map[time.Time]Meteorologia)}
}

func (clima *Clima) AgregarMeteorologia(meteorologia Meteorologia) error {
	_, existe := clima.registroMetereologico[meteorologia.fecha]
	if existe {
		return fmt.Errorf("ya existe un registro metereológico para la fecha %s", meteorologia.fecha.Format("2006-01-02"))
	}

	clima.registroMetereologico[meteorologia.fecha] = meteorologia
	return nil
}

func (clima *Clima) ModificarMetereologia(fecha time.Time, agua float64, temperaturaMin float64, temperaturaMax float64) error {
	meteorologia, err := NuevaMeteorologia(fecha, agua, temperaturaMin, temperaturaMax)
	if err == nil {
		clima.registroMetereologico[meteorologia.fecha] = meteorologia
		return nil
	}
	return err
}
