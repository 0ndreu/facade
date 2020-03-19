package facade

import "fmt"

type guitarist struct {
	name string
}

func (g *guitarist) playCoolOpening(name string) {
	fmt.Println(name, "играет крутое вступление")
}

func (g *guitarist) playCoolRiffs() {
	fmt.Println("играет крутые риффы")
}

func (g *guitarist) playAnotherCoolRiffs() {
	fmt.Println("играет другие крутые риффы")
}

func (g *guitarist) playIncrediblyCoolSolo() {
	fmt.Println("выдает невероятно крутое соло")
}

func (g *guitarist) playFinalAccord() {
	fmt.Println("заканчивает песню мощным аккордом")
}
