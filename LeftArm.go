package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

// func main() {
// 	const (
// 		protocol = "tcp"
// 		addr     = ":22"
// 	)

// 	fmt.Println(DualArmControl.NewServerRun(protocol, addr, connectHandler))
// }
const LeftFileName = "LeftArm"

func LeftArmHandler(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Status :")
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = text[:len(text)-1]
		switch text {
		case "f":
			freeDrive(conn, LeftFileName)
		case "m":
			move(conn, LeftFileName)
		case "c":
			conn.Write([]byte{3})
		case "o":
			conn.Write([]byte{33})
		case "mJ":

			fmt.Println("請輸入轉動圈數 :")
			fmt.Print("=>")
			t, err := reader.ReadString('\n')
			t = strings.Replace(t, "\n", "", -1)
			t = t[:len(t)-1]
			if err != nil {
				log.Println(err)
			}
			i, err := strconv.Atoi(t)
			if err != nil {
				log.Println(err)
			}
			isOpen := i > 0
			i *= 2
			if !isOpen {
				i = 0 - i
			}
			dataf := make([]byte, 1024)

			clockwise := []byte("(3.14)")
			counterclockwise := []byte("(-3.14)")
			for index := 0; index < i; index++ {
				conn.Write([]byte{2})
				datafLen, err := conn.Read(dataf)
				if err != nil {
					log.Println(err)
					return
				}
				d := dataf[:datafLen]
				if string(d) == "moveJoint" {
					if err != nil {
						log.Println(err)
						return
					}
					if index%2 == 0 {

						if isOpen {
							conn.Write(clockwise)
						} else {
							conn.Write(counterclockwise)
						}
						conn.Write([]byte{3})

					} else {
						if isOpen {
							conn.Write(counterclockwise)
						} else {
							conn.Write(clockwise)
						}
						conn.Write([]byte{33})
					}

				} else {
					log.Println("完蛋了...")
				}
			}
			fmt.Printf("完成%s轉圈!!", t)

		default:
			fmt.Println("有問題喔... 你給的Status是 : ", []byte(text))
			fmt.Println([]byte("f"))
		}

	}

}
