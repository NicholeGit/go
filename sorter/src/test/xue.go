package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
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

type Bus struct {
	Base
}

func (foo *Bus) Bar() {

	fmt.Println("Foo Bar")
	foo.Base.Bar()
}

func Count(ch chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	//fmt.Println("Counting")
	ch <- 1

	fmt.Println("Counting")
	wg.Done()
}

func tChannl() {
	chs := make([]chan int, 10)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i], &wg)
	}
	for _, ch := range chs {
		<-ch
	}

	wg.Wait()
}

type Element interface{}
type List []Element

type Person struct {
	name string
	age  int
}

//定义了 String 方法，实现了 fmt.Stringer
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + "years)"
}

func main() {
	num := runtime.NumCPU()
	fmt.Println("NumCPU = ", num)
	runtime.GOMAXPROCS(num)

	list := make(List, 3)
	list[0] = 1       // an int
	list[1] = "Hello" // a string
	list[2] = Person{"Dennis", 70}
	for index, element := range list {
		if value, ok := element.(int); ok {
			fmt.Printf("list[%d] is an int and its value is %d\n", index,
				value)
		} else if value, ok := element.(string); ok {
			fmt.Printf("list[%d] is a string and its value is %s\n", index,
				value)
		} else if value, ok := element.(Person); ok {
			fmt.Printf("list[%d] is a Person and its value is %s\n", index,
				value)
		} else {
			fmt.Println("list[%d] is of a different type", index)
		}
	}
}
