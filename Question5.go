package main

import "fmt"

func main() {
	//var i int
	var a string = "abcdef"
	j := len(a)

	for i := 0; i < j; i++ {
		if (a[i+1] - a[i]) == (a[j] - a[j-1]) {
			fmt.Println("Given string is equal difference ")
		} else {
			fmt.Println("Given string is unequal difference ")
		}
	}
}
