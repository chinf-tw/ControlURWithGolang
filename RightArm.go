package main

import (
	"UR3Handler"
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

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
			data, err := ioutil.ReadFile("./" + FileName + ".json")
			check(err)
			var poses Poses
			check(json.Unmarshal(data, &poses))
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
				poses.AddPose(Pose{PoseName: t, PoseData: p})
				fmt.Println("完成!!")
				poses.AddPosesToJsonFile(FileName)
			} else {
				log.Println("完蛋了...")
			}

		case "m":

			data, err := ioutil.ReadFile("./" + FileName + ".json")
			check(err)
			var poses Poses
			check(json.Unmarshal(data, &poses))
			datam := make([]byte, 1024)
			conn.Write([]byte{11})
			datamLen, err := conn.Read(datam)
			if err != nil {
				log.Println(err)
				return
			}
			d := datam[:datamLen]
			if string(d) == "move" {
				for i, v := range poses {
					poseUr := UR3Handler.FloatListToUR3Float(v.PoseData)
					aaa := fmt.Sprintf("%d. Action %s : %s\n", i, v.PoseName, poseUr)
					fmt.Println(aaa)
					if err != nil {
						log.Println(err)
					}
				}
				fmt.Println("請輸入動作編號 :")
				fmt.Print("=>")
				t, _ := reader.ReadString('\n')
				t = strings.Replace(t, "\n", "", -1)
				t = t[:len(t)-1]
				index, err := strconv.Atoi(t)
				check(err)
				ur3Pose := UR3Handler.FloatListToUR3Float(poses[index].PoseData)
				conn.Write([]byte(ur3Pose))
				fmt.Println(ur3Pose)
			}
		default:
			fmt.Println("有問題喔... 你給的Status是 : ", []byte(text))
			fmt.Println([]byte("f"))
		}

	}

}
