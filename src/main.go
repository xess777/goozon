package main

import (
	"fmt"
	"spaces"
)

func main() {
	s := []rune(" On  my   home world")
	sPtr := spaces.RemoveExtraSpaces(&s)
	fmt.Println(string(*sPtr))
}
