package main

import (
	"DualArmControl"
	"UR3Handler"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	const (
		protocol = "tcp"
		addr     = ":21"
	)

	fmt.Println(DualArmControl.NewServerRun(protocol, addr, connectHandler))
}

func connectHandler(conn net.Conn) {
	defer conn.Close()
	pose := make(map[string][]float64)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Status :")
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = text[:1]
		switch text {
		case "f":
			dataf := make([]byte, 1024)
			conn.Write([]byte{1})

			datafLen, err := conn.Read(dataf)
			if err != nil {
				log.Println(err)
				return
			}
			d := dataf[:datafLen]
			if string(d) == "freedrive_mode" {
				conn.Write([]byte{2})

				if err != nil {
					log.Println(err)
					return
				}

				fmt.Println("請輸入動作名稱 :")
				fmt.Print("=>")
				t, _ := reader.ReadString('\n')
				t = strings.Replace(t, "\n", "", -1)
				conn.Write([]byte{3})
				datafLen, err = conn.Read(dataf)
				log.Println(string(dataf[:datafLen]))
				p, err := UR3Handler.PoseTypeToFloatList(string(dataf[:datafLen]))
				if err != nil {
					log.Println(err)
				}
				pose[t] = p
				fmt.Println("完成!!")

			} else {
				log.Println("完蛋了...")
			}

		case "m":
			datam := make([]byte, 1024)
			conn.Write([]byte{11})
			datamLen, err := conn.Read(datam)
			if err != nil {
				log.Println(err)
				return
			}
			d := datam[:datamLen]
			if string(d) == "move" {
				for k, v := range pose {
					poseUr := UR3Handler.FloatListToUR3Float(v)
					_, err := fmt.Printf("動作 %s : %s\n", k, poseUr)
					if err != nil {
						log.Println(err)
					}
				}
				fmt.Println("請輸入動作名稱 :")
				fmt.Print("=>")
				t, _ := reader.ReadString('\n')
				t = strings.Replace(t, "\n", "", -1)
				ur3Pose := UR3Handler.FloatListToUR3Float(pose[t])
				conn.Write([]byte(ur3Pose))
				fmt.Println(ur3Pose)
			}
		default:
			fmt.Println("有問題喔... 你給的Status是 : ", []byte(text))
			fmt.Println([]byte("f"))
		}

	}

}
