package bioma

import (
	"time"
)

type Meteorologia struct {
	fecha          time.Time
	aguaLlovida    float64
	temperaturaMin float64
	temperaturaMax float64
}
