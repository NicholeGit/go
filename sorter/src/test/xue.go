package main

import (
	"fmt"
)

type Base struct {
	Name string
}

func (base *Base) Foo() {
	fmt.Println("base foo")
}

func (base *Base) Bar() {
	fmt.Println("base Bar")
}

type Foo struct {
	Base
}

func (foo *Foo) Bar() {
	foo.Base.Bar()
	fmt.Println("Foo Bar")
}

func main() {
	var f = new(Foo)

}
