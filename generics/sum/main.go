package main

import "fmt"

func main() {
	i := 100
	f := 99.99

	fmt.Println(sum(i, i))
	fmt.Println(sum(f, f))
}

type number interface {
	int | float32 | float64
}

func sum[T number](val1, val2 T) T {
	return val1 + val2
}
