package main

import (
	"fmt"
)

func main() {
	for index := 0; index < 5; index++ {
		fmt.Println(index%2 == 0)
	}
}
