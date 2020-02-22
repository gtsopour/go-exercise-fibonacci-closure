package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610,
// 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025,
// 121393, 196418, 317811, 514229
func fibonacci() func() int {
	numbers := make(map[int]int)
	n := 0
	return func() int {
		if n == 0 {
			numbers[n] = 0
			n++
			return 0
		}
		if n == 1 {
			numbers[n] = 1
			n++
			return 1
		}
		number := numbers[n-1] + numbers[n-2]
		numbers[n] = number
		n++
		return number
	}}

func _fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 30; i++ {
		fmt.Println(f())
	}
}
