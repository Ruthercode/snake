package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"time"

	"./tools"
)

var (
	ycor     int
	xcor     int
	dir      string
	foodAvbl int = 0
	foodX    int
	foodY    int
)

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
			} else if i == foodY && j == foodX {
				fmt.Print("F")
				foodAvbl++
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

// not real-time input
func move() {
	r := bufio.NewReader(os.Stdin)
	c, err := r.ReadByte()
	if err != nil {
		panic(err)
	}
	dir = string(c)
	switch dir {
	case "w":
		ycor -= 1
	case "s":
		ycor += 1
	case "a":
		xcor -= 1
	case "d":
		xcor += 1
	}

	if xcor == 29 {
		xcor = 1
	} else if xcor == 0 {
		xcor = 28
	}
	if ycor == 29 {
		ycor = 1
	} else if ycor == 0 {
		ycor = 28
	}
}
func food() {
	if foodAvbl == 0 {
		foodX = tools.RandInt(1, 29)
		foodY = tools.RandInt(1, 29)
		foodAvbl++
	}
}
func main() {
	ycor = 15
	xcor = 15
	for true {
		food()
		move()
		clear()
		draw()
		time.Sleep(time.Second * 1)
	}
}
