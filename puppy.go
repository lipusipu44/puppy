package puppy

import (
	"github.com/lipusipu44/dog"
)

func Bark() string {
	return "Woof!!!"
}

func Barks() string {
	return "Woof!!! Woof!!! Woof!!!"
}

func BigBark() string {
	return dog.ConvertToAdult(Bark())
}

func BigBarks() string {
	return dog.ConvertToAdult(Barks())
}
