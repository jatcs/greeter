package greeter

import (
	"fmt"
)

type Greeter struct {
}

func NewGreeter() *Greeter {

	return &Greeter{}
}

func (g*Greeter)Say(name string) {
	fmt.Printf("Hi %s!\n", name)
}

