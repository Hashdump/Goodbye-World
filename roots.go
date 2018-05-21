package main

/*
 * go run roots.go
 */

import (
	"fmt"
)

var (
	d = 72
	e = []int{0, 29, 7, 0, 3, -79, 55, 24, 3, -6, -8, -67}
)

type C struct {
	x byte
}

func (c C) s() string {
	return string([]byte{c.x})
}

func c(a ...byte) string {
	return C{byte(len(a))}.s()
}

func a(x int, y ...func(int) string) (z string) {
	for i, j := range y {
		z += j(e[i] + x)
		x = e[i] + x
	}
	return z
}

func t(x int) string {
	return c(make([]byte, x)...)
}

func i(x int, z func(int) string) []func(int) string {
	a := make([]func(int) string, 0)
	for y := 0; y < x; y++ {
		a = append(a, z)
	}
	return a
}

func main() {
	fmt.Println(a(len(e)*len(a(41, i(6, t)...)), i(len(e), t)...))
}
