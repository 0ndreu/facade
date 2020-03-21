package drummer

import "fmt"

type Drummer interface {
	StartPlaying(name string)
	StopPlaying(name string)
}

type drummer struct {
	name string
}

func (d *drummer) StartPlaying(name string) {
	fmt.Println(name, "начинает играть")
}

func (d *drummer) StopPlaying(name string) {
	fmt.Println(name, "заканчивает играть")
}

func NewDrummer() Drummer {
	return &drummer{}
}
