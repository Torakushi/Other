package exemple

import (
	"fmt"
	"unsafe"
)

func UnsafeTest() {
	l := []task{
		{Name: "toto", id: 1, age: 1},
		{Name: "toto2", id: 2, age: 2},
		{Name: "toto3", id: 3, age: 3},
	}

	fmt.Printf("%T\n", l)
	fmt.Printf("%T\n", l[0])

	l2 := *(*[]task_archive)(unsafe.Pointer(&l))

	fmt.Printf("%T\n", l2)
	fmt.Printf("%T\n", l2[0])

}

type task struct {
	Name string
	id   int
	age  int
}

type task_archive struct {
	Name string
	id   int
	age  int
}
