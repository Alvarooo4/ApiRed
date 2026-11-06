package bioma

type ApiRed struct {
	climaPorCoordenada map[Coordenada]Clima
	flores             map[Flor]struct{}
}
