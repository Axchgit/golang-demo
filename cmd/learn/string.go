package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println('a')
	fmt.Print(`
		343243243
		ewrsedrfdsf
		fdsfd
	`)
	a := "abcdefg "
	fmt.Println(a[1])
	fmt.Println(string(a[1]))
	fmt.Println(a[:2])
	fmt.Println(string(a[:2]))
	fmt.Println(string(a))

	bytes := []byte(a)
	fmt.Println(bytes)
	fmt.Println(string(bytes))

	bytes = append(bytes, 109, 103, 32, 111)

	bytes[0] = 120
	fmt.Println(string(bytes))

	b := "一二三四五"

	fmt.Println(len(b))

	copB := make([]byte, 111)

	copy(copB, b)

	cB := strings.Clone(b)

	fmt.Println(string(copB))
	fmt.Println(string(cB))

	pB := b + "你好"
	fmt.Println(pB)
	pBb := []byte(pB)
	pBb = append(pBb, "nihaoa"...)
	fmt.Println(string(pBb))

	c := string(pBb)
	c += "13213232"
	fmt.Println(c)

	d := "123456"
	fmt.Println(d[1])
	db := []byte(d)
	fmt.Println(db[1])
	fmt.Println(d[1] == db[1])

	额 := "中国"
	e := []rune(额)
	fmt.Println(string(e[1]))

	builder := strings.Builder{}

	builder.WriteString("链接B")
	builder.WriteString("链接C")
	fmt.Println(builder.String())
}
