package main

import (
	"DualArmControl"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const FileName = "RightArm"

func main() {
	const (
		protocol  = "tcp"
		RightAddr = ":21"
		LeftAddr  = ":22"
	)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Select arm :")
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	text = text[:1]
	switch text {
	case "r":
		fmt.Println(DualArmControl.NewServerRun(protocol, RightAddr, RightArmHandler))
	case "l":
		fmt.Println(DualArmControl.NewServerRun(protocol, LeftAddr, LeftArmHandler))
	}

}
