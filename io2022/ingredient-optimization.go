package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	var t int
	io := bufio.NewReader(os.Stdin)
	//io := bufio.NewReader(file)
	fmt.Fscan(io, &t)

	for i := 0; i < t; i++ {

		var d, n, u int
		fmt.Fscan(io, &d, &n, &u)
		//fmt.Printf("%v %v %v\n", d, n, u)

		deliveries := make([][]int, d)

		for i := 0; i < d; i++ {
			deliveries[i] = make([]int, 3)
			fmt.Fscan(io, &deliveries[i][0], //time
				&deliveries[i][1], //quantity
				&deliveries[i][2]) //number of minutes those leaves can be stored and remain fresh
		}

		orders := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(io, &orders[i])
		}

		//fmt.Printf("%v\n", deliveries)
		//fmt.Printf("%v\n", orders)

		delivery := 0
		order := 0
		m := make(map[int]int) //time -> quantity
		for order < n {
			if delivery == d || orders[order] < deliveries[delivery][0] {
				keys := sortMap(m)
				count := u
				for _, k := range keys {

					if orders[order] >= k {
						delete(m, k) //spoiled
					} else {
						if count >= m[k] {
							count -= m[k]
							delete(m, k) //all gone
						} else {
							m[k] -= count
							count = 0
						}
					}
					if count == 0 {
						break
					}
				}

				if count != 0 {
					break
				}
				order++
			} else {
				expiration := deliveries[delivery][0] + deliveries[delivery][2]
				m[expiration] += deliveries[delivery][1]
				delivery++
			}
		}

		fmt.Printf("Case #%d: %d\n", i+1, order)
	}
}

func sortMap(m map[int]int) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}
