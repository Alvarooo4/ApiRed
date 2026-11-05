package bioma

type ApiRed struct {
	clima_por_coordenada map[Coordenada]Clima
	flores               map[Flor]struct{}
}
