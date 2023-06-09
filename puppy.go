package puppy

import "github.com/sercankavdir/dog"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof!"
}

func BigBark() {
	return dog.WhenGrowUp(Bark())
}

func BigBarks() {
	return dog.WhenGrowUp(Barks())
}
