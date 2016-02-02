package main

import "fmt"

func main() {
	fmt.Println(vlen()) // Testing the response if no argument is not passed.
	fmt.Println(max(1, 4, 0, 4, 5, 3, 100, 23, 59))
}

// ...Type parameter makes this a variadic function. Meaning 0 to N number of arguments can be used for that parameter.
// The parameter is a slice when used.
func vlen(a ...int) int {
	return len(a)
}

// Does a max of 0 make sense here if a argument is not passed?
func max(a ...int) int {
	var max int
	for _, v := range a {
		if v > max {
			max = v
		}
	}
	return max
}

// Compiling with typeMix uncommented will result in `can only use ... as final argument in list`
// func typeMix(a ...int, b string) int{
//	return 0
// }