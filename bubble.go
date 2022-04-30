package main

import "fmt"

func main() {
	a := [6]int{1, 1, 1, 1, 1, 1}
	length := len(a)
	counter := 1
	loops := 0
	for counter != 0 {
		counter = 0
		loops++
		fmt.Println(loops)
		fmt.Println(a)
		for i := 0; i < length-1; i++ {
			if a[i] > a[i+1] {
				a[i+1], a[i] = a[i], a[i+1]
				counter++
			}
		}
	}
}
