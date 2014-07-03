package main

import (
	"fmt"
	"runtime"
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

func main() {
	num := runtime.NumCPU()
	fmt.Println("NumCPU = ", num)
	runtime.GOMAXPROCS(num)

	tSelect()

}
