package bioma

type ApiRed struct {
	climaPorCoordenada map[CoordenadaDecimal]Clima
	flores             map[Flor]struct{}
}
