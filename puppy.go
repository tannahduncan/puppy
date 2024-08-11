package puppy

import (
	"fmt"

	"github.com/tannahduncan/dog"
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
	return dog.WhenGrownUp(Bark())
}
func From1() {
	fmt.Println("I'm from version 1.0.0")
}
func From2() {
	fmt.Println("I'm from version 2.0.0")
}
func From3() {
	fmt.Println("I'm from version 3.0.0")
}
