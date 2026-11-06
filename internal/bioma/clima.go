package bioma

import (
	"time"
)

type Clima struct {
	meteorologiaPorFecha map[time.Time]Meteorologia
}
