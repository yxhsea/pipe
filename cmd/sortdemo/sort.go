package main

import (
	"sort"
	"fmt"
)

func main()  {
	num := []int{1,5,2,4,2,7,6,8}
	sort.Ints(num)
	for _, v := range num {
		fmt.Println(v)
	}
}
