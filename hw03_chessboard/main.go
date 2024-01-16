package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Введите ширину и длину доски: ")
	fmt.Scan(&x, &y)

	for i := 1; i <= y; i++ {
		if i%2 != 0 {
			for j := 1; j <= x; j++ {
				if j%2 != 0 {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		} else {
			for j := 1; j <= x; j++ {
				if j%2 != 0 {
					fmt.Print(" ")
				} else {
					fmt.Print("#")
				}
			}
			fmt.Println()
		}
	}
}
