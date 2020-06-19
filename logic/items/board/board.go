package items

import (
	tile "bingo-tokenring/logic/items/tile"
)

type Board struct {
	Name  string
	Tiles []tile.Tile
}
