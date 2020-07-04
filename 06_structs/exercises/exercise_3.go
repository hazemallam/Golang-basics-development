package main
import "fmt"
type Vehicle struct{
	doors  int
	color string

}
type Truck struct{
	Vehicle
	fourWheel bool

}
type Sedan struct{
	Vehicle
	luxery bool
}

func main(){
	truck := Truck{
		Vehicle : Vehicle{
			doors : 2,
			color : "red",
		},
		fourWheel : true,

	}

	sedan := Sedan{
		Vehicle : Vehicle{
			doors : 4,
			color : "yellow",

		},
		luxery : true,

	}
	fmt.Println(truck)
	fmt.Println(sedan)

	fmt.Println(truck.Vehicle)

	//create and use an Anonymous struct

	s := struct{
		first string
		friends map[string]int
		drinks []string
	}{
		first : "James Bond",
		friends : map[string]int{
			"Miss moneypenny" : 444,
			"name" : 500,
			"last" : 700,
		},
		drinks : []string{
			"water",
			"tea",
			"juce",
		},
	}
	fmt.Println(s)
}