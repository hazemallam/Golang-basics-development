package main
import "fmt"
func main() {
	m := map[string]int{
		"a": 20,
		"b": 30,
	}
	fmt.Println(m)

	//printing the value using key
	fmt.Println(m["b"])

	//if you entered a key that is not existing it will print 0 like the following example
	fmt.Println(m["c"])

	//to avoid this situation use this like the follwoing example
	v, ok := m["c"]
	fmt.Println(v, ok)

	//so we can use this in the following context
	if v, ok := m["c"]; ok{
		fmt.Println("HTIS IS THE IF PRINT", v)
	}

	if v, ok := m["a"]; ok{
		fmt.Println("HTIS IS THE IF PRINT", v)
	}

	//addin key and value to an existing map
	 m["d"] = 80

	 //printing keys and values in map

	 for key, value := range m {
		 fmt.Println(key, value)
	 }

	 //deleting key from a map
	 delete(m, "d")
	 fmt.Println(m)

	 //if you tried to delete a key that does not exist there is no error appear like the following example
	 delete(m, "n")
	 fmt.Println(m)

	 //to avoid this error or situation do like the following example
	 if value, ok := m["b"]; ok {
		fmt.Println("value : ", value)
		delete(m, "b")
	 }
	 fmt.Println(m)


}
