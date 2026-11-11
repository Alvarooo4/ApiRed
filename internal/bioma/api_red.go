package bioma

type ApiRed struct {
	climaPorCoordenada map[CoordenadaDecimal]Clima
	heuristicaPorFlor  map[string]HeuristicaFloracion
}
