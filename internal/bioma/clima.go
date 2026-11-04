package bioma

import (
	"time"
)

type Clima struct {
	coordenada            Coordenada
	flores                []string
	registroMeteorologico map[time.Time]Meteorologia
}
