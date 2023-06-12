package puppy

import (
	"fmt"

	"github.com/sercankavdir/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrowUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrowUp(Barks())
}

func From11() {
	fmt.Println("I'm from version 1.1.0")
}
