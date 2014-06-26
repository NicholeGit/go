package main

import "fmt"
import "mymath"

// PersonInfo是一个包含个人详细信息的类型
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)
	// 往这个map里插入几条数据
	personDB["12345"] = PersonInfo{"12345", "Tom", "Room 203,..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 101,..."}
	// 从这个map查找键为"1234"的信息
	person, ok := personDB["1234"]
	delete(personDB, "1234")

	// ok是一个返回的bool型，返回true表示找到了对应的数据
	if ok {
		fmt.Println("Found person", person.Name, "with ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}
	fmt.Println("Found person", example(0))
	c, _ := mymath.Add(1, 2)
	fmt.Println("mymath add", c)
	test()
}

func example(x int) int {
	if x == 0 {
		return 5
	} else {
		return x
	}
}

func test() {
	var j int = 5
	a := func() func() {
		var i int = 10
		return func() {
			fmt.Printf("i, j: %d, %d\n", i, j)
		}
	}()
	a()
	j *= 2
	a()
}
