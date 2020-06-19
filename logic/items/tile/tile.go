package tile

type Tile struct {
	Letter       string
	Number       int
	taken        bool
	neighborhood *Octagon
}
