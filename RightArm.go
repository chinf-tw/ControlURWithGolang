package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const RightFileName = "RightArm"

func RightArmHandler(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Status :")
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = text[:1]
		switch text {
		case "f":
			freeDrive(conn, RightFileName)
		case "m":
			move(conn, RightFileName)
		default:
			fmt.Println("有問題喔... 你給的Status是 : ", []byte(text))
			fmt.Println([]byte("f"))
		}

	}

}
