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

// Starts the song
func (d *drummer) StartPlaying(name string) {
	fmt.Println(name, "начинает играть")
}

// Ends the song
func (d *drummer) StopPlaying(name string) {
	fmt.Println(name, "заканчивает играть")
}

// Create drummer
func NewDrummer() Drummer {
	return &drummer{}
}
