package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return 1
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
