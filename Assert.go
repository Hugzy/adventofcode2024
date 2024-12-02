package main

import "fmt"

var (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

func Assert(guess int, answer int) {
	if guess == answer {
		fmt.Println(Green, "guess of", guess, "is correct!", Reset)
		return
	}

	fmt.Println(Red, "guess is not correct expected", answer, "got", guess, Reset)
}

func PrintAnswer(answer any) {
	fmt.Println(Yellow, "Answer is", answer, Reset)
}
