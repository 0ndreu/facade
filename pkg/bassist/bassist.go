package bassist

import "fmt"

// Bassist can follow the drums, change rhythm and stop playing
type Bassist interface {
	FollowTheDrums(name string)
	ChangeRhythm(name string)
	StopPlaying(name string)
}

type bassist struct {
	name string
}

// Bassist follows the drums
func (b *bassist) FollowTheDrums(name string) {
	fmt.Println(name, "следует за барабанами")
}

// Bassist changes the rhythm
func (b *bassist) ChangeRhythm(name string) {
	fmt.Println(name, "перешел на ритм ")
}

// Bassist stops playing
func (b *bassist) StopPlaying(name string) {
	fmt.Println(name, "заканчивает играть")
}

// Creating new bassist
func NewBassist() Bassist {
	return &bassist{}
}
