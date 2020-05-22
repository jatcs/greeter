package main

import (
	greeter "github.com/jturne3428/greeter/packGreet"
)

func main() {
	//declare Greeter type (struct) in the package greeter
	g := greeter.Greeter{}

	//calling functions which use Greeter type as object
	g.NewGreeter()
	g.Say("Jessica")

	g.Say("Cat")
}
