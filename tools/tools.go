package tools

import (
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func Clear() {
	bash := exec.Command("clear")
	bash.Stdout = os.Stdout
	bash.Run()
}
