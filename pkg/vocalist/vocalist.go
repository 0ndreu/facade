package vocalist

import "fmt"

// Vocalist can only sing a couplet or a chords
type Vocalist interface {
	SingCouplet(name string)
	SingChorus(name string)
}

type vocalist struct {
	name string
}

// Sing a couplet
func (v *vocalist) SingCouplet(name string) {
	fmt.Println(name, "спел куплет")
}

// Sing a chorus
func (v *vocalist) SingChorus(name string) {
	fmt.Println(name, "спел припев")
}

// Create new vocalist
func NewVocalist() Vocalist {
	return &vocalist{}
}
