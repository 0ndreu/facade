package drummer

import "fmt"

// Drummer only starts and stops playing
type Drummer interface {
	StartPlaying(name string)
	StopPlaying(name string)
}

type drummer struct {
	name string
}

// Start the song
func (d *drummer) StartPlaying(name string) {
	fmt.Println(name, "начинает играть")
}

// End the song
func (d *drummer) StopPlaying(name string) {
	fmt.Println(name, "заканчивает играть")
}

// Create new drummer
func NewDrummer() Drummer {
	return &drummer{}
}
