package facade

import "fmt"

type drummer struct {
	name string
}

func (d *drummer) startPlaying(name string) {
	fmt.Println(name, "начинает играть")
}

func (d *drummer) stopPlaying() {
	fmt.Println("заканчивает играть")
}
