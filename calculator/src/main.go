package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(" ----------------------------")
	fmt.Println("|           계산기           |")
	fmt.Println(" ----------------------------")
	var calculation string
	fmt.Println("연산을 입력해주세요~~")
	fmt.Scanln(&calculation)
	// calcSlice := make([]string, 20)
	temp := strings.Split(calculation, "")

	// calcSlice = append(calcSlice, temp)
	fmt.Println(temp[3])
}
