package bassist

import "fmt"

type Bassist interface {
	FollowTheDrums(name string)
	ChangeRhythm(name string)
	StopPlaying(name string)
}

type bassist struct {
	name string
}

func (b *bassist) FollowTheDrums(name string) {
	fmt.Println(name, "следует за барабанами")
}

func (b *bassist) ChangeRhythm(name string) {
	fmt.Println(name, "перешел на ритм ")
}

func (b *bassist) StopPlaying(name string) {
	fmt.Println(name, "заканчивает играть")
}

func NewBassist() Bassist {
	return &bassist{name: "Мартин"}
}
