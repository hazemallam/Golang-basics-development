// package mymath provides math solutions
package mymath

// Sum adds unlimited number of values of type int
func Sum(n ...int) int {
	total := 0
	for _, v := range n {
		total += v
	}
	return total
}
