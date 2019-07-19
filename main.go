package main

import (
	"fmt"
)

func main(){
	stepsIn([]int{3,4,2,7}, func(n int) {
		fmt.Println(n)
	})
}

func stepsIn(nums []int, cb func(int)) {
	for _, num := range nums {
		cb(num)
	}
}