package main

import (
	"bufio"
	"fmt"
	"os"
)

func input(str string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(str)
	if scanner.Scan() {
		inputStr := scanner.Text()
		return inputStr
	}
	fmt.Println("Failed to get input.")
	return ""
}
