package tile

type Octagon struct {
	Up        *Tile
	RightUp   *Tile
	Right     *Tile
	RightDown *Tile
	Down      *Tile
	DownLeft  *Tile
	Left      *Tile
	LeftUp    *Tile
}
