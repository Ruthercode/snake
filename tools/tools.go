package tools

import (
	"math/rand"
	"os"
	"os/exec"
	"time"
	"github.com/eiannone/keyboard"
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

func GetKey(tm time.Duration) (ch rune, err error) {
	if err = keyboard.Open(); err != nil {
		return
	}
	defer keyboard.Close()

	var (
		chChan  = make(chan rune, 1)
		errChan = make(chan error, 1)

		timer = time.NewTimer(tm)
	)
	defer timer.Stop()

	go func(chChan chan<- rune, errChan chan<- error) {
		ch, _, err := keyboard.GetSingleKey()
		if err != nil {
			errChan <- err
			return
		}
		chChan <- ch
	}(chChan, errChan)

	select {
	case <-timer.C:
		return rune(0), err
	case ch = <-chChan:
	case err = <-errChan:
	}

	return
}