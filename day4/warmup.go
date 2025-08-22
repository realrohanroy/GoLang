package main

import "fmt"
import "strings"


func vowels(s string){
	i :=0
		for _, v := range s {
			if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u'{
				i++
			}
		}
		fmt.Println(i)
}

func reverse(s string)  {
	runes := []rune(s)
	right:= len(s)-1
	left:=0
	for right>left{
		swap:= runes[left] 
		runes[left] = (runes[right])
		runes[right] = swap

		right--
		left++
	}
	fmt.Println(string(runes))
}

func join(word []string)  {
	fmt.Println(strings.Join(word," "))
}

func main(){
	// even := make ([]int ,0,7)
	// nums :=  []int{1,2,3,4,5,6,7}
	// for i := 0; i < len(nums) ; i++ {
	// 	if nums[i] %2==0 {
	// 		fmt.Print(nums[i])
	// 	}
	// }
	// fmt.Println(even)

	// s:="hello"
	// vowels(s)
	// reverse(s)

	word := []string{"go", "is","cool"}
	join(word)
}