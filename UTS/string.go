package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var (
		msg, msg2 string
	)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		msg = scanner.Text()
	}

	scanner2 := bufio.NewScanner(os.Stdin)
	if scanner2.Scan() {
		msg2 = scanner2.Text()
	}

	fmt.Println("Input was: ", msg, msg2)
}
