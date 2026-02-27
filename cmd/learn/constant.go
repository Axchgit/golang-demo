package main

import (
	"fmt"
	"unsafe"
)

const a = 10

const (
	b = iota
	c
	d
)

const (
	e = 1 << iota
	f = 3 << iota
	g
	h
)

const (
	i = iota
	j
	k
	l = "hello"
	m = 100
	o
	n = iota
)

const (
	ab = "abc"
	ac = len(ab)
	ad = unsafe.Sizeof(ab)
	ae = unsafe.Sizeof(ac)
	af = unsafe.Sizeof(ad)
)

func main() {
	fmt.Println(a)

	fmt.Println(b, c, d)

	fmt.Println(e, f, g, h)
	fmt.Println(i, j, k, l, m, n, o)

	fmt.Println(ab, ac, ad, ae, af)
}
