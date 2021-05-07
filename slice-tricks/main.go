package main

import "fmt"

func main() {
	// fmt.Println("appendSlice", appendSlice())
	// fmt.Println("copySliceV1:", copySliceV1())
	// fmt.Println("copySliceV2:", copySliceV2())
	// fmt.Println("copySliceV3:", copySliceV3())
	// fmt.Println("cutSlice:", cutSlice())
	// fmt.Println("deleteIndexV1:", deleteIndexV1())
	// fmt.Println("deleteIndexV2:", deleteIndexV2())
	// fmt.Println("deleteNoOrder:", deleteNoOrder())
	// fmt.Println("filterInPlace:", filterInPlace())
	// fmt.Println("filterNewSlice:", filterNewSlice())
	// fmt.Println("insert:", insert())
	// fmt.Println("insertSlice:", insertSlice())
	// fmt.Println("push:", push())
	// v, a := pop()
	// fmt.Printf("pop: %d, %v\n", v, a)
	// v, a := shift()
	// fmt.Printf("shift: %d, %v\n", v, a)
	fmt.Println("unshift:", unshift())

}

func appendSlice() []int {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 8, 10}
	return append(a, b...)
}

func copySliceV1() []int {
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, len(a))
	fmt.Printf("1. main  -- i  %T: &i=%p i=%v\n", a, &a, a)
	fmt.Printf("1. main  -- i  %T: &i=%p i=%v\n", b, &b, b)
	copy(b, a)
	fmt.Printf("1. main  -- i  %T: &i=%p i=%v\n", b, &b, b)
	return b
}

func copySliceV2() []int {
	a := []int{1, 2, 3, 4, 5}
	return append(a[:0:0], a...)
}

func copySliceV3() []int {
	a := []int{1, 2, 3, 4, 5}
	return append([]int(nil), a...)
}

func cutSlice() []int {
	a := []int{1, 2, 3, 4, 5}
	i := 1
	j := 4
	return append(a[:i], a[j:]...)
}

func deleteIndexV1() []int {
	a := []int{1, 2, 3, 4, 5}
	i := 1
	return append(a[:i], a[i+1:]...)

}
func deleteIndexV2() []int {
	a := []int{1, 2, 3, 4, 5}
	i := 1
	return a[:i+copy(a[i:], a[i+1:])]
}

func deleteNoOrder() []int {
	a := []int{1, 2, 3, 4, 5}
	i := 1
	a[i] = a[len(a)-1]
	return a[:len(a)-1]
}

func filterInPlace() []int {
	a := []int{1, 2, 3, 4, 5}
	filterFunc := func(x int) bool {
		return x > 2
	}
	n := 0
	for _, x := range a {
		if filterFunc(x) {
			a[n] = x
			n++
		}
	}
	return a[:n]

}

func filterNewSlice() []int {
	a := []int{1, 2, 3, 4, 5}
	filterFunc := func(x int) bool {
		return x <= 2
	}
	var b []int

	for _, x := range a {
		if filterFunc(x) {
			b = append(b, x)
		}

	}
	return b

}

func insert() []int {
	a := []int{1, 2, 3, 4, 5}
	v := 10
	i := 1
	return append(a[:i], append([]int{v}, a[i:]...)...)
}

func insertSlice() []int {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8}
	i := 1
	return append(a[:i], append(b, a[i:]...)...)
}

func push() []int {
	a := []int{1, 2, 3, 4, 5}
	v := 6
	return append(a, v)
}

func pop() (int, []int) {
	a := []int{1, 2, 3, 4, 5}
	return a[len(a)-1], a[:len(a)-1]
}

func shift() (int, []int) {
	a := []int{1, 2, 3, 4, 5}
	return a[0], a[1:]
}

func unshift() []int {
	a := []int{1, 2, 3, 4, 5}
	v := 0
	return append([]int{v}, a...)
}
