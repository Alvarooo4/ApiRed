package bioma

type ApiRed struct {
	coordenadasRegistradas map[Coordenada]struct{}
	climaPorCoordenada     map[Coordenada]Clima
	flores                 map[Flor]struct{}
}
