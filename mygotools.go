package mygotools

import (
	"bufio"
	"os"
	"runtime"
	"strings"
)

func GetUserInput(printText string) string {
	println(printText)
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	if runtime.GOOS == "windows" {
		userInput = strings.TrimSuffix(userInput, "\r\n")
	} else {
		userInput = strings.TrimSuffix(userInput, "\n")
	}
	return userInput
}
