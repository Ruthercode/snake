package main

import (
	"fmt"
	"time"
)

var ycor int
var xcor int

func draw() {
	for i := 0; i < 29; i++ {
		fmt.Print("#")
	}
	fmt.Println("#")
	for i := 0; i < 28; i++ {
		fmt.Print("#")
		for j := 0; j < 28; j++ {
			if j == xcor && i == ycor {
				fmt.Print("O")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("#")
	}

	for i := 0; i < 30; i++ {
		fmt.Print("#")
	}
}
func main() {
	ycor = 15
	xcor = 15
	for true {
		draw()
		time.Sleep(time.Second * 2)
	}
}
