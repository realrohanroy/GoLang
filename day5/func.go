package main
import "fmt"

func avg(num []int)(int,int,error){
	sum:=0
	n := len(num)
	if n == 0{
		return 0,0,fmt.Errorf("empty slice")
	}
	for _,i:=range num{
		sum=sum+i
	}
	average := sum/n
	return sum, average, nil
}

func avg2(nums ...int) (int,error) {
	sum:=0
	
	for _,n:=range nums{
		sum+=n
	}
	return sum/len(nums),nil
}

func main(){
	// num:= []int{}
	// sum,avg,err:=avg(num)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("Sum:",sum,"Average:",avg)

	// avg,err:=avg2(0)
	// if err != nil {
	// 	fmt.Println("Error:",err)
	// 	return
	// }
	// fmt.Println("AVG:",avg)

	isEven:= func(n int) bool {
		if n %2 ==0 {
			return true
		}else {
			return false
		}
	}
	fmt.Println(isEven(4))



	
}