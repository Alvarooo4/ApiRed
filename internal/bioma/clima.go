package bioma

import (
	"time"
)

type Clima struct {
	coordenada            Coordenada
	registroMeteorologico map[time.Time]Meteorologia
}
