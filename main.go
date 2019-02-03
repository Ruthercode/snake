package main

import (
	"./tools"
	"fmt"
	"github.com/eiannone/keyboard"
	"time"
)

const n = 150 // time

// pajiloy gamedev
var (
	ycor     int
	xcor     int
	dir      string
	foodAvbl int = 0
	foodX    int
	foodY    int
	score    int  = 0
	x_tail        = make([]int, 1)
	y_tail        = make([]int, 1)
	gps_x         = make([]int, 1)
	gps_y         = make([]int, 1)
	gameOver bool = false
	costil   string
)

func tail_move() {
	j := 0
	if score != 0 {
		for i := score - 1; i >= 0; i-- {
			index := len(gps_x) - 1 - j
			x_tail[i] = gps_x[index]
			y_tail[i] = gps_y[index]
			j++
		}
	}
}

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
				solve := 0
				for k := 0; k < score; k++ {
					if i == y_tail[k] && j == x_tail[k] {
						fmt.Print("s")
						solve++
					}
				}
				if solve == 0 {
					fmt.Print(" ")
				}
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

func input() {
	ch, err := tools.GetKey(time.Millisecond * n)
	if err != nil {
		panic(err)
	}
	if ch != rune(0) && (string(ch) == "w" || string(ch) == "s" || string(ch) == "d" || string(ch) == "a") {
		dir = string(ch)
	}
}

func move() {

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

	if xcor == 30 {
		xcor = 1
	} else if xcor == 0 {
		xcor = 29
	} else if ycor == 30 {
		ycor = 1
	} else if ycor == 0 {
		ycor = 29
	}

	for i := 0; i < score; i++ {
		if xcor == x_tail[i] && ycor == y_tail[i] {
			gameOver = true
		}
	}

	gps_x = append(gps_x, xcor)
	gps_y = append(gps_y, ycor)

	if xcor == foodX && ycor == foodY {
		foodAvbl = 0
		x_tail = append(x_tail, xcor)
		y_tail = append(y_tail, ycor)
		if dir == "w" {
			y_tail[score]++
		} else if dir == "s" {
			y_tail[score]--
		} else if dir == "a" {
			x_tail[score]++
		} else if dir == "d" {
			x_tail[score]--
		}
		score++
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
	for !gameOver {
		food()
		tools.Clear()
		draw()
		input()
		tail_move()
		move()
	}
	fmt.Println(" ")
	fmt.Println("GAME OVER!")
	fmt.Println("Press any key to continue ")
	char, _, err := keyboard.GetSingleKey()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\r\n", char)
	tools.Clear()
}
