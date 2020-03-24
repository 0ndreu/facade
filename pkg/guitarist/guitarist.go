package guitarist

import "fmt"

// guitarist can play openings, riffs, solos and final accord
type Guitarist interface {
	PlayCoolOpening(name string)
	PlayCoolRiffs(name string)
	PlayAnotherCoolRiffs(name string)
	PlayIncrediblyCoolSolo(name string)
	PlayFinalAccord(name string)
}

type guitarist struct {
	name string
}

// Plays opening of the song
func (g *guitarist) PlayCoolOpening(name string) {
	fmt.Println(name, "играет крутое вступление")
}

// Plays 1st riffs in song
func (g *guitarist) PlayCoolRiffs(name string) {
	fmt.Println(name, "играет крутые риффы")
}

// Plays 2nd riffs in song
func (g *guitarist) PlayAnotherCoolRiffs(name string) {
	fmt.Println(name, "играет другие крутые риффы")
}

// Plays the bestest solo in the song
func (g *guitarist) PlayIncrediblyCoolSolo(name string) {
	fmt.Println(name, "выдает невероятно крутое соло")
}

// Plays final accord of the song
func (g *guitarist) PlayFinalAccord(name string) {
	fmt.Println(name, "заканчивает песню мощным аккордом")
}

// Create new guitarist
func NewGuitarist() Guitarist {
	return &guitarist{}
}
