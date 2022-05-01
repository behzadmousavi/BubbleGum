package main

import "fmt"

func main() {
	a := [6]int{5, 1, 3, 6, 2, 4}
	length := len(a)
	//counter := 0
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
			fmt.Println(a)
		}
	}

}
