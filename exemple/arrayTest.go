package exemple

import (
	"fmt"
)

func ArrayTest() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a[2])
	fmt.Printf("A: len: %d, cap: %d\n", len(a), cap(a))

	b := a[1:3]
	fmt.Printf("b: %v\n", b)
	b[1] = 4
	fmt.Printf("b: %v\n", b)
	fmt.Printf("B: len: %d, cap: %d\n", len(b), cap(b))

	fmt.Printf("a: %v\n", a)

	a = append(a, 4)
	fmt.Printf("a: %v\n", a)

	b[1] = 5
	fmt.Printf("b: %v\n", b)

	fmt.Printf("a: %v\n", a)

}
