package main

import (
	"fmt"
	"os"
	"os/exec"
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

func clear() {
	bash := exec.Command("clear")
	bash.Stdout = os.Stdout
	bash.Run()
}
func main() {
	ycor = 15
	xcor = 15
	for true {
		clear()
		draw()
		time.Sleep(time.Second * 2)
	}
}
