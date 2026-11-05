package bioma

import (
	"time"
)

type Clima struct {
	coordenada             Coordenada
	meteorologia_por_fecha map[time.Time]Meteorologia
}
