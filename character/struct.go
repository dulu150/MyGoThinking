// testarrayslice project main.go
package main

import (
	"fmt"
)

type abc struct {
	a int
}

func main() {
	b := &abc{a: 1}
	fmt.Println(b.a)

	var c abc
	c = abc{a: 2}
	fmt.Println(c.a)
}
