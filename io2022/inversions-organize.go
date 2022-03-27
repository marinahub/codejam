package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var t int
	io := bufio.NewReader(os.Stdin)
	//io := bufio.NewReader(file)
	fmt.Fscan(io, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(io, &n)
		arr := make([]string, 2*n)
		for j := 0; j < 2*n; j++ {
			fmt.Fscan(io, &arr[j])
		}
		//fmt.Printf("%v", arr)

		res := [2][2]int{}
		for i := 0; i < 2*n; i++ {
			for j := 0; j < 2*n; j++ {
				if arr[i][j] == 'I' {
					res[i/n][j/n]++
				}
			}
		}

		fmt.Printf("Case #%d: %d\n", i+1, Abs(res[0][0]-res[1][1])+Abs(res[0][1]-res[1][0]))
	}

}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
