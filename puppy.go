package puppy

import (
	"github.com/rezky1313/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof!! Woof!"
}

func BigBark() string {
	return dog.Dog(Bark())
}

func BigBarks() string {
	return dog.Dog(Barks())
}
