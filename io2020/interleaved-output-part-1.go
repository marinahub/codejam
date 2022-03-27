package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int;

	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)

	for u := 0; u < n; u++ {
		var input string
		fmt.Fscan(io, &input)

		i:=0
		I:=0
		result :=0

		for c := 0; c < len(input); c++ {
			if (input[c]=='i'){
				i++
			}
			if (input[c]=='I'){
				I++
			}
			if (input[c]=='o'){
				if (i>0){
					i--
				}else{
					I--
				}
			}
			if (input[c]=='O'){
				if (I>0){
					I--
					result++
				}else{
					i--
				}
			}
		}
		fmt.Printf("Case #%d: %d\n",u+1,result)

	}
}
