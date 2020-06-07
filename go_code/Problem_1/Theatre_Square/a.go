package main

import "fmt"

func main() {
	var n, m, a int64
	fmt.Scanf("%d %d %d", &n, &m, &a)
	fmt.Printf("%d\n", ((n+a-1)/a)*((m+a-1)/a))
}
