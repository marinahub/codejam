package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &n)

	for i := 0; i < n; i++ {
		var leng int
		fmt.Fscan(io, &leng)
		arr := make([]int, leng)
		for j := 0; j < leng; j++ {
			fmt.Fscan(io, &arr[j])
		}

		//reverse
		for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
			arr[i], arr[j] = arr[j], arr[i]
		}

		res := make([]string, leng)
		A := make([]int, leng)
		B := make([]int, leng)
		
		for i := 0; i < leng/2; i++ {
			A[2*i] = 2*i + 1
			A[2*i+1] = 2 * i
		}

		//fmt.Println(A)
		for i := 0; i < leng/2; i++ {
			B[arr[2*i]-1] = arr[2*i+1]-1
			B[arr[2*i+1]-1] = arr[2*i]-1
		}
		//fmt.Println(B)

		for k, v := range res {
			if v == "" {
				queue := list.New()
				queue.PushBack(Pair{k, true})
				for queue.Len() > 0 {
					e := queue.Front()

					if res[e.Value.(Pair).a] != "" {
						queue.Remove(e)
						continue
					}
					d := "R"
					if e.Value.(Pair).b {
						d = "L"
					}
					res[e.Value.(Pair).a] = d
					queue.PushBack(Pair{A[e.Value.(Pair).a], !e.Value.(Pair).b})
					queue.PushBack(Pair{B[e.Value.(Pair).a], !e.Value.(Pair).b})

					queue.Remove(e)
				}
			}

		}
		resString := strings.Join(res,"")
		fmt.Printf("Case #%d: %s\n", i+1, resString)
	}
}

type Pair struct {
	a int
	b bool
}
