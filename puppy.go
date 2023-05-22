package puppy

import (
	"github.com/JohnnyBinder/dog"
)

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof! Woof! Woof!"
}

func BigBark() string {
	return dog.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog.WhenGrownUp(Barks())
}

func SuperBigBark() string {
	return dog.WhenBigAndGrownUp(Barks())
}