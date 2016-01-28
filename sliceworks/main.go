package main

import "fmt"

func fillslice(v []int) []int {
	for i := 0; i < 30; i++ { // purposely larger than the capacity size of slice to see the behavior.
		fmt.Println(i, " len(v)=", len(v), " cap(v)=", cap(v))
		if i < len(v) {
			v[i] = i // reassign initialized slice element by index.
		} else {
			v = append(v, i) // Anything beyond the initialized element then must be appended.
		}
	}
	return v
}

func main() {
	m := make([]int, 5, 25) // Build Slice with make(), args: Type, initial size, array capacity
	//m := []int{0,0,0,0,0} ~= make([]int, 5) args: Type, initial size/ array capacity
	fmt.Println("cap: ", cap(m), "len: ", len(m))
	for k, v := range m {   // initial slice size 5
		fmt.Println(k, " - ", v) // print idx=k 0 - 5 value of 0
	}
	m = fillslice(m)
	for k, v := range m {
		fmt.Println(k, " - ", v) // Filled slice
	}
	fmt.Println("cap: ", cap(m), "len: ", len(m))
}
