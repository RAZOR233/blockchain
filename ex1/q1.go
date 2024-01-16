package main

import "fmt"

func gongbei(a, b int) int {
	var i int
	for i = 1; ; i++ {
		if (a*i)%b == 0 {
			break
		}
	}
	return a * i
}
func main() {
	var a, b, c, i int
	fmt.Scanf("%d%d%d", &a, &b, &c)
	i = gongbei(a, b)
	fmt.Println(gongbei(i, c))
}
