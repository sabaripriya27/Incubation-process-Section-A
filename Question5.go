package main

import "fmt"

func main() {
	//var i int
	var a string = "abcdef"
	n:= len(a)

	for i := 0; i < n; i++ {
		if (a[i+1] - a[i]) == (a[n] - a[n-1]) {
			fmt.Println("Given string is equal difference ")
		} else {
			fmt.Println("Given string is unequal difference ")
		}
	}
}
}
