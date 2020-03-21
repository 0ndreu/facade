package guitarist

import "fmt"

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

func (g *guitarist) PlayCoolOpening(name string) {
	fmt.Println(name, "играет крутое вступление")
}

func (g *guitarist) PlayCoolRiffs(name string) {
	fmt.Println(name, "играет крутые риффы")
}

func (g *guitarist) PlayAnotherCoolRiffs(name string) {
	fmt.Println(name, "играет другие крутые риффы")
}

func (g *guitarist) PlayIncrediblyCoolSolo(name string) {
	fmt.Println(name, "выдает невероятно крутое соло")
}

func (g *guitarist) PlayFinalAccord(name string) {
	fmt.Println(name, "заканчивает песню мощным аккордом")
}

func NewGuitarist() Guitarist {
	return &guitarist{}
}
