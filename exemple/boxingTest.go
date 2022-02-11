package exemple

import (
	"fmt"
)

func BoxingTest() {
	t := &taskb{}
	l := (*task2)(t)
	fmt.Println(l)
}

type taskb struct {
	Name string
	age  int
}

type task2 struct {
	Name string
	age  int
}
