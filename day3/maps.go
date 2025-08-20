package main

import "fmt"

func main(){

	//----declaring a map--------
	//Method 1: make
	m:=make(map[string]int)
	m["Alice"] = 23
	m["Bob"] = 30

	//Method 2: map literal
	scores := map[string]int{
		"Go": 100,
		"Java":90,
		"Python":95,
	}

	// fmt.Println("m:",m)
	// fmt.Println("scores:",scores)

	//---accessing and checking keys------
	age:= m["Alice"]
	fmt.Println("Alice age:", age)
	//if key doesnt exits, returns 0 value
	fmt.Println("Eve age:",m["Eve"])

	//to check if a key exist
	val, ok:=m["Eve"] //Use this to distinguish between Zero value and Key Not Present 
	if ok{
		fmt.Println("Eve exist", val)
	} else{
		fmt.Println("key does not exist")
	}


	//-----UPDATING AND DELETING----
	m["Bob"]= 35 //update
	delete(m,"Alice") //delete key
	fmt.Println("After updating and deleting", m)


	//------Iterating Over a Map-------------
	for k, v:=range scores {
		fmt.Println("Lang:",k,"Scores:", v)
	}

	//-------Length--
	fmt.Println("Length of scores:",len(scores))
}