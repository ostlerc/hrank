package main

import "fmt"

func read(n int) []int {
	ret := make([]int, n)
	for i := range ret {
		_, err := fmt.Scan(&ret[i])
		if err != nil {
			panic(err)
		}
	}
	return ret
}

func sum(v []int) int {
	ret := 0
	for i := range v {
		ret += v[i]
	}
	return ret
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Println(sum(read(N)))
}
