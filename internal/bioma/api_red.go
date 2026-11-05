package bioma

type ApiRed struct {
	coordenadas_registradas map[Coordenada]struct{}
	clima_por_coordenada    map[Coordenada]Clima
	flores                  map[Flor]struct{}
}
