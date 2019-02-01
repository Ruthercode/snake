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
	x_tail       = make([]int, 1)
	y_tail       = make([]int, 1)
	gps_x        = make([]int, 1)
	gps_y        = make([]int, 1)
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

	gps_x = append(gps_x, xcor)
	gps_y = append(gps_y, ycor)

	if xcor == foodX && ycor == foodY {
		foodAvbl = 0
		if x_tail[0] == 0 {
			x_tail[0] = xcor
			y_tail[0] = ycor
			if dir == "w" {
				y_tail[0]++
			} else if dir == "s" {
				y_tail[0]--
			} else if dir == "a" {
				x_tail[0]++
			} else if dir == "d" {
				x_tail[0]--
			}
		} else {
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
		}
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
		tail_move()
		move()
		tools.Clear()
		draw()
		time.Sleep(time.Millisecond * 100)
	}
}
