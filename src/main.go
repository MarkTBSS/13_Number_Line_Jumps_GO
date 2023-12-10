/* func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	var result string
	result = "NO"
	// result := "NO"
	var y float64
	y = float64(x1-x2) / float64(v2-v1)
	fmt.Println(y)
	fmt.Println(float64(int(y)))
	if y > 0 && y == float64(int(y)) {
		result = "YES"
	}
	return result
} */

package main

import "fmt"

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	var result string
	result = "NO"
	// result := "NO"
	var ableToCatch bool
	ableToCatch = v2 < v1

	if ableToCatch && (x1-x2)%(v2-v1) == 0 {
		return "YES"
	}
	return result
}

func main() {
	fmt.Println(kangaroo(21, 6, 47, 3))
}
