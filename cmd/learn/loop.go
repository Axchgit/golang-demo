package main

import "fmt"

func main() {

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// a := 10

	// for a > 0 {
	// 	a--
	// 	fmt.Println(a)
	// }

	// 九九乘法表
	// for m := 1; m <= 9; m++ {
	// 	for n := 1; n <= m; n++ {
	// 		fmt.Printf(" %d * %d = %d  ", m, n, m*n)
	// 	}
	// 	fmt.Printf("\n")
	// }
	// fmt.Printf("\n")
	// for g := 9; g >= 1; g-- {
	// 	for j := g; j >= 1; j-- {
	// 		fmt.Printf(" %d * %d = %d ", g, j, g*j)
	// 	}
	// 	fmt.Printf("\n")
	// }

	// e := 10
	// for l := range 10 {
	// 	fmt.Println(l)
	// }

	// for le := range e {
	// 	fmt.Println(le)
	// }

	// const ce = 10
	// for rce := range ce {
	// 	if rce > 7 {
	// 		break
	// 	}
	// 	fmt.Println(rce)
	// }

	// for aa := 1; aa < 9; aa++ {
	// 	for ab := 1; ab < 9; ab++ {
	// 		if ab > aa {
	// 			break
	// 		}
	// 		fmt.Println(aa, ab)
	// 	}
	// }

out:
	for ad := 1; ad <= 9; ad++ {
		for ac := 1; ac <= 9; ac++ {
			if ac > ad {
				// goto out 才会重新执行代码块，陷入死循环
				continue out // continue会从for的步进开始执行，不会重新执行ad赋值
			}
			fmt.Println(ad, ac)
		}
	}

}
