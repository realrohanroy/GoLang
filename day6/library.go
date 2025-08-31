package main
import "fmt"

type Book struct {
	title string
	author string
	copies int
}

func(b *Book) Borrow(){
	if b.copies>0 {
		b.copies-=1
	}else{
		fmt.Println("No copies avalable")
	}
}

func(b *Book) Return() {
	b.copies+=1
}

func(b Book) Details(){
	fmt.Println(b.title,"by",b.author,"(",b.copies,")")
}

func main(){
	b1 := Book{title:"Kafka on the shore",author:"Murakami",copies:50}
	b1.Borrow()
	b1.Details()
	
}