package facade

import "fmt"

type vocalist struct {
	name string
}

func (v *vocalist) singCouplet(name string) {
	fmt.Println(name, "спел куплет")
}

func (v *vocalist) singChorus(name string) {
	fmt.Println(name, "спел припев")
}
