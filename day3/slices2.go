package main

import "fmt"

func main () {
	nums:= []int{1,2,3,4,5}
	fmt.Println("nums:",nums)

	sub:= nums[1:4] //2,3,4
	// fmt.Println("sub:",sub)
//---changes the main slice---
	// sub[0] = 99
	// fmt.Println("after change in sub:",sub) 
	// fmt.Println("nums also changes: ",nums)

//--lets copy to a new slice--
	independent := make([]int,len(sub))
	copy(independent,sub)
	
	nums[1] = 99
	fmt.Println("nums:",nums)
	fmt.Println("sub is still linked: ", sub)
	fmt.Println("independent:",independent)
}