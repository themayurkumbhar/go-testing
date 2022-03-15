package messages

import (
	"fmt"
)

//public method to test within pkg
func Greet(name string) string {
	return fmt.Sprintf("Hello, %v!\n", name)
}

//private method to pkg
func depart(name string) string {
	return fmt.Sprintf("Good bye %v.\n", name)
}

func nonCoveredMethod() {
	fmt.Print("Hello!!")
}
