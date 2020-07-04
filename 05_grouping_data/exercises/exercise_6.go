package main
import "fmt"
func main(){
	m := map[string] []string{
		"bond_james": []string{"Shaken not stirred", "Martins", "Women"},
		"moneypenny_mess": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr": []string{"Ice cream", "Sunsets"},
	}
	m["no_name"] = []string {"Tennis", "Football", "Swimming"}
	for k, record := range m{
		fmt.Printf("This the record for %v : \n", k)
		for i, v := range record {
			fmt.Println("\t", i, v)
		}
	}
	fmt.Println("deleting record")
	if _, ok := m["no_name"]; ok{
		delete(m, "no_name")
		fmt.Println("record deleted")

	}
	for k, record := range m{
		fmt.Printf("This the record for %v : \n", k)
		for i, v := range record {
			fmt.Println("\t", i, v)
		}
	}
}