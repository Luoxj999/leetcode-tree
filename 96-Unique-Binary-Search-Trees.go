package main

import (
	"fmt"
)

func main() {
	fmt.Println(numTrees(4))
}
// n个结点构造多少种树
// http://c.biancheng.net/view/3402.html
func numTrees(n int) int {
	s := make([]int, 0)
	s = append(s, 1) // 0
	s = append(s, 1) // 1
	s = append(s, 2) // 2
	s = append(s, 5) // 3
	for i := 4; i <= n; i++ {
		temp := i - 1
		total := 0
		for j := 0; j <= temp; j++ {
			total += s[j] * s[temp - j]
		}
		s = append(s, total)
	}
	return s[n]
}