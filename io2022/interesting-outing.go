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
		var v int
		fmt.Fscan(io, &v)
		edges := make([][]int, v-1)
		sum := 0
		for j := 0; j < v-1; j++ {
			edges[j] = make([]int, 3)
			fmt.Fscan(io, &edges[j][0],
				&edges[j][1],
				&edges[j][2])

			sum += edges[j][2]
		}

		sum *= 2

		m := make(map[int][]int)
		w := make(map[int][]int)
		for k := 0; k < len(edges); k++ {
			m[edges[k][0]] = append(m[edges[k][0]], edges[k][1])
			w[edges[k][0]] = append(w[edges[k][0]], edges[k][2])

			m[edges[k][1]] = append(m[edges[k][1]], edges[k][0])
			w[edges[k][1]] = append(w[edges[k][1]], edges[k][2])
		}

		i1, i2 := helper(m, w, -1, 1)

		//fmt.Printf("%v\n", m)
		//fmt.Printf("%v\n", w)

		fmt.Printf("Case #%d: %d\n", i+1, sum-max(i1, i2))
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func helper(m map[int][]int, w map[int][]int, parent int, self int) (int, int) {
	maxLength := 0
	maxLengthVert := 0
	candidates := make([]int, 0)
	//fmt.Printf("%v\n", m[self])
	for idx, val := range m[self] {
		if val == parent {
			continue
		} else {
			i, i2 := helper(m, w, self, val)
			candidates = append(candidates, w[self][idx]+i2)
			if maxLength < i {
				maxLength = i
			}
		}
	}

	sort.Ints(candidates)
	//fmt.Printf("%v\n", candidates)
	l := len(candidates)
	if l >= 2 {
		res := candidates[l-1] + candidates[l-2]
		if maxLength < res {
			maxLength = res
		}
	}

	if l >= 1 {
		maxLengthVert = candidates[l-1]
	}

	return maxLength, maxLengthVert
}
