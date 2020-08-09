package items

//Tile .
type Tile struct {
	Letter string `json:"letter"`
	Number int    `json:"number"`
	Taken  bool   `json:"taken"`
}
