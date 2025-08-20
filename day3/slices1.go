package main 

import "fmt"

func main() {
	s := make ([]int,0,3)
	fmt.Println("start:", s, "len:",len(s),"cap:",cap(s))

	for i:=1;i<=13;i++{
		s =append(s,i)
        fmt.Println("after append", i, "=>", s, "len:", len(s), "cap:", cap(s))
	}	
}	