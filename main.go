package main

import (
	"bufio"
	"fmt"
	"os"
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
	score    int = 0
)

func draw() {
	for i := 0; i < 30; i++ {
		fmt.Print("#")
	}
	fmt.Println("#")

	for i := 1; i < 30; i++ {
		fmt.Print("#")
		for j := 1; j < 30; j++ {
			if j == xcor && i == ycor {
				fmt.Print("O")
			} else if i == foodY && j == foodX {
				fmt.Print("F")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("#")
	}

	for i := 0; i < 30; i++ {
		fmt.Print("#")
	}
	fmt.Println("#")

	fmt.Print("You score is ", score)
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

	if xcor == foodX && ycor == foodY {
		foodAvbl = 0
		score++
	}

	if xcor == 30 {
		xcor = 1
	}
	if xcor == 0 {
		xcor = 29
	}
	if ycor == 30 {
		ycor = 1
	}
	if ycor == 0 {
		ycor = 29
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
		tools.Clear()
		draw()
		time.Sleep(time.Millisecond * 100)
	}
}
