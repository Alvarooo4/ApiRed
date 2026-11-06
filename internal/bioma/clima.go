package bioma

import (
	"time"
)

type Clima struct {
	meteorologia_por_fecha map[time.Time]Meteorologia
}
