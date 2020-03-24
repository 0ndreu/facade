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

// Vocalist sings a couplet
func (v *vocalist) SingCouplet(name string) {
	fmt.Println(name, "спел куплет")
}

// Vocalist sings a chorus
func (v *vocalist) SingChorus(name string) {
	fmt.Println(name, "спел припев")
}

// New vocalist creating
func NewVocalist() Vocalist {
	return &vocalist{}
}
