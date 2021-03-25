package mygotools

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var hostOS = runtime.GOOS

func GetUserInput(printText string) string {
	println(printText)
	reader := bufio.NewReader(os.Stdin)
	userInput, _ := reader.ReadString('\n')
	if hostOS == "windows" {
		userInput = strings.TrimSuffix(userInput, "\r\n")
	} else {
		userInput = strings.TrimSuffix(userInput, "\n")
	}
	return userInput
}

func ClearCLI() {
	var clear *exec.Cmd
	if hostOS == "windows" {
		clear = exec.Command("cmd", "/c", "cls")
	}
	if hostOS == "linux" {
		clear = exec.Command("clear")
	}
	clear.Stdout = os.Stdout
	clear.Run()
}
