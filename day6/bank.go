package main
import "fmt"

type BankAccount struct{
	owner string
	balance float64
}

func(b *BankAccount) deposite (amt float64){
	b.balance+=amt
}
func(b *BankAccount) withdraw (amt float64){
	if b.balance <amt {
		fmt.Println("Insufficient balance")
	}else{
		b.balance-=amt
	}
}
func(b BankAccount) checkbalance(){
	fmt.Println("Balance:",b.balance)
}

func main(){
	b1:=BankAccount{owner:"Rohan",balance: 100000}
	fmt.Println(b1.owner)
	b1.deposite(5000)
	b1.checkbalance()
	b1.withdraw(10000)
	b1.checkbalance()
	b1.withdraw(100000)
}
