package main
import (
	"fmt"
	"sort"
)
type PPerson struct{
	Name string
	Age int
}
func (p PPerson) String() string{
	return fmt.Sprintf("%s : %d", p.Name, p.Age)
}

//byAge implements sort.Interface for []pperson based on the Age field
type ByAge []PPerson
type ByName []PPerson
// must have theis methods
func (a ByAge) Len() int {
	return len(a)
}
func (a ByAge) Swap(i, j int){
	 a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool{
	return a[i].Age < a[j].Age
} 

func (a ByName) Len() int {
	return len(a)
}
func (a ByName) Swap(i, j int){
	 a[i], a[j] = a[j], a[i]
}
func (a ByName) Less(i, j int) bool{
	return a[i].Name < a[j].Name
} 
func main(){
	people := []PPerson{
		{"Bob", 31},
		{"John", 42},
		{"michael", 17},
		{"Jenny", 26},
	}
	fmt.Println(people)
	sort.Sort(ByAge(people)) //sort.Sort(Data Interface{})
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}