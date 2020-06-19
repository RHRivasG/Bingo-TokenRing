package tile

type Checked interface {
	Check(tile Tile) bool
}

type DiagonalDown struct{}
type DiagonalUp struct{}
type Vertical struct{}
type Horizontal struct{}

func (d *DiagonalDown) Check(tile Tile) {
}

func (d *DiagonalUp) Check(tile Tile) {
}

func (d *Vertical) Check(tile Tile) {
}

func (d *Horizontal) Check(tile Tile) {
}
