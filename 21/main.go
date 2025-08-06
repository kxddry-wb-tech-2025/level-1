package main

import "fmt"

// Interface 1
type Bird interface {
	Fly()
	MakeSound()
}

// implementation

type Pigeon struct{}

func (*Pigeon) Fly()       {}
func (*Pigeon) MakeSound() {}

// Interface 2
type Dog interface {
	Bark()
}

type Snoop struct{}

func (*Snoop) Bark() {}

type DogAdapter struct {
	dog Dog
}

func (*DogAdapter) Fly() {
	fmt.Println("dog can't fly")
}
func (d *DogAdapter) MakeSound() {
	d.dog.Bark()
}

func main() {
	pigeon := &Pigeon{}
	dog := &Snoop{}

	fmt.Println("pigeon")
	pigeon.Fly()
	pigeon.MakeSound()

	adapter := &DogAdapter{dog}
	fmt.Println("adapter dog")
	adapter.Fly()
	adapter.MakeSound()

}
