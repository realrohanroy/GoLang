package main
import "fmt"

//Define struct
type Person struct {
	name string
	age int
	city string
}

func (p1 Person) greet(){
	fmt.Println("hello,my name is",p1.name)
}
func(p1 *Person) hadBirthday(){
	p1.age++
}

func main (){
	p1:=Person{"Rohan",21,"Pune"}
	fmt.Println(p1)

//Best practice:-
	p2:=Person{name:"Aditya",age:22,city:"PCMC"}
	fmt.Println("Name:",p2.name)
	fmt.Println("Age:",p2.age)
//Zero-value struct
	var p3 Person
	fmt.Println(p3)

	p1.age = 23 //updating
	fmt.Println(p1.age)

//Pointer in struct
	pPtr := &p1
	fmt.Println("Name:",pPtr.name)
	pPtr.name = "Rohan Roy"
	fmt.Println("Updated Name:",pPtr.name)

	//Functions/methods in struct
	p1.greet() 
	p1.hadBirthday()
	fmt.Println("After bday age:",p1.age)

}