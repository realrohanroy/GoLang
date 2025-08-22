package main 

import "fmt"

// func greet(name string) string {
// 	return "Welcome," + name+ " to day 2 of GO!"
// }

// func getDetails(name string)(string,int){
// 	message := "Welcome," + name + "!"
// 	length := len(name)
// 	return message,length
// }

func main(){
	// var name string
	// fmt.Print("enter ur name: ")
	// fmt.Scanln(&name)

	// // fmt.Println(greet(name))
	// msg, nameLength := 	getDetails(name)
	// fmt.Println(msg)
	// fmt.Println("your name has", nameLength, " characters.")

//-----Arrays-----//
	// var numbers [5]int

	// numbers[0] = 10
	// numbers[1] = 20
	// numbers[2] = 30

	// fmt.Println("Array:",numbers)
	// fmt.Println("first el", numbers[0])

//-----Slices-----//
	// fruits := []string{"Apple","Banana","Mango"}

	// fmt.Println("Fruits: ", fruits)

	// //add an element
	// fruits = append(fruits, "Orange")
	// fmt.Println("After append", fruits)

	// //slice part of it
	// subSlice := fruits[0:3] //1 to 2{3-1}
	// fmt.Println(subSlice)

	// numbers := []int{10,20,30,40,50}
	// fmt.Println(numbers)
	// fmt.Println("len: ", len(numbers), ", cap: ", cap(numbers))

	// fmt.Println(numbers[1:4])//references the same array
	// fmt.Println(numbers[:3])

	// numbers = append(numbers,60,70,80)
	// fmt.Println(numbers)

	// copySlice := make([]int,len(numbers)) //creates copySlice of int type and len same as numbers
	// copy(copySlice,numbers) //copies numbers to copySlice

	// //update
	// numbers[2] = 90
	// fmt.Println(numbers)
	// //remove
	// index:=2
	// numbers = append(numbers[:index],numbers[index+1:]...)
	// //insert
	// value:=100
	// numbers = append(numbers[:index], append([]int {value},numbers[index:]...)...)

	cities := []string {"nagpur","pune","mumbai","chandrapur"}
	cities = append(cities, "nashik","madgao")
	fmt.Println(cities)
	cities = append(cities[:2],cities[2+1:]...)
	fmt.Println(cities)
	cities[1] = "PCMC"
	fmt.Println(cities)
}