package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Введите ширину и длину доски: ")
	fmt.Scan(&x, &y)
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			switch {
			case (i%2 != 0 && j%2 != 0) || (i%2 == 0 && j%2 == 0):
				fmt.Print("#")
			case (i%2 != 0 && j%2 == 0) || (i%2 == 0 && j%2 != 0):
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
