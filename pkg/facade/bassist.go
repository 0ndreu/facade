package facade

import "fmt"

type bassist struct {
	name string
}

func (b *bassist) followTheDrums(name string) {
	fmt.Println(name, "следует за барабанами")
}

func (b *bassist) changeRhythm() {
	fmt.Println("перешел на ритм ")
}

func (b *bassist) stopPlaying() {
	fmt.Println("заканчивает играть")
}
