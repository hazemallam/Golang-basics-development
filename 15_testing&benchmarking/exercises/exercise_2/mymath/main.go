//package mymath provides basic math operation
package mymath

import "sort"

//CenterdAvg returns the average of list after removing the smallest and largest values
func CenterdAvg(xi []int) float64{
	sort.Ints(xi)
	xi = xi[1:(len(xi) - 1)]
	sum := 0
	for _, v := range xi {
		sum += v
	}
	avg := float64(sum) / float64(len(xi))
	return avg
}
