package main

import "fmt"

func main() {

	nums := []int {2, 3, 4}
	fmt.Println(nums)
	nums = append(nums, 5)
	fmt.Println(nums)
}
