package vocalist

import "fmt"

type Vocalist interface {
	SingCouplet(name string)
	SingChorus(name string)
}

type vocalist struct {
	name string
}

func (v *vocalist) SingCouplet(name string) {
	fmt.Println(name, "спел куплет")
}

func (v *vocalist) SingChorus(name string) {
	fmt.Println(name, "спел припев")
}

func NewVocalist() Vocalist {
	return &vocalist{name: "Джонни"}
}
