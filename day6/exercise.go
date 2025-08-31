package main
import "fmt"

type Engine struct {
	horsepower int
	fueltype string
}
type Car struct {
	brand string
	model string
	engine Engine  //COMPOSITION in Go 
}

// func (c *Car) drive(km int){
// 	c.mileage+=km/2
// }
// func (c Car ) info(){
// 	fmt.Println("Brand:",c.brand,", Model:",c.model,", Mileage:",c.mileage)
// }

// func (c Car) info(){
// 	fmt.Println("Brand:",c.brand,"Model:",c.model,"Bhp:",c.engine.horsepower,"FuelType:",c.engine.fueltype)
// }

type Server struct {
	ip string
	location string
}
type Cloudinstance struct{
	provider string
	server Server
}

func main(){
	// c:= Car{brand:"Tata",model:"Nexon",mileage:15}
	// c.info()
	// c.drive(10)
	// c.info()


	// eng1:=Engine{horsepower:294,fueltype: "petrol" }
	// car1:=Car{brand:"BMW",model:"X3",engine: eng1}  //car is made of another struct engine
	// car2 := Car{brand: "Tesla", model: "Model S", engine: Engine{horsepower: 500, fueltype: "Electric"}}  //Inheritance-like behaviour	
	// car2.info()
	// car1.info()


	s1:=Server{ip:"192.168.1.0",location: "asia-south-1"}
	c1:=Cloudinstance{provider: "AWS",server:s1}
	fmt.Println("Provider:",c1.provider,", IP:",c1.server.ip,", Location:",c1.server.location)
}