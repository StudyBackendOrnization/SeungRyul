package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(" ----------------------------")
	fmt.Println("|           계산기           |")
	fmt.Println(" ----------------------------")

	var expression string

	fmt.Println("연산을 입력해주세요~~")
	fmt.Scanln(&expression)

	temp := calcMulAndDiv(strings.Split(expression, ""), 0)
	fmt.Println("first result : ", temp)
	temp2 := clacSumAndDif(temp, 0)
	fmt.Println("second result : ", temp2)
}

func calcMulAndDiv(expr []string, i int) []string {
	var exprSlice []string
	var num float32
	var exprSlice2 []string
	temp := 0
	exprSlice = expr
	fmt.Println(exprSlice)
	fmt.Println(len(exprSlice)-2-1, "testestestestest", i)
	if exprSlice[i] == "*" || exprSlice[i] == "/" {
		if exprSlice[i] == "*" {
			num = mul(convertStringToFloat(exprSlice[i-1]), convertStringToFloat(exprSlice[i+1]))
			temp = 1
		}
		if exprSlice[i] == "/" {
			num = div(convertStringToFloat(exprSlice[i-1]), convertStringToFloat(exprSlice[i+1]))
			temp = 1
		}

		exprSlice2 = append(exprSlice2, exprSlice[:i-1]...)
		exprSlice2 = append(exprSlice2, convertFloatToString(num))
		if len(exprSlice)-2-1 >= i {
			exprSlice2 = append(exprSlice2, exprSlice[i+2:]...)
		}
		fmt.Println("zzzzz", exprSlice[i])
		fmt.Println(exprSlice[i-1], exprSlice[i], exprSlice[i+1])
	}

	if temp == 0 {
		exprSlice2 = exprSlice
	}

	fmt.Println("len(exprSlice2) : ", len(exprSlice2), ", exprSlice2 : ", exprSlice2, ", i : ", i, ", temp : ", temp)
	fmt.Println()
	if len(exprSlice2) == 1 || len(exprSlice2)-1 == i {
		fmt.Println("1------------------------------------------------------------1")
		return exprSlice2
	}
	// charList := strings.Join(exprSlice2, "")
	if temp == 1 {
		return calcMulAndDiv(exprSlice2, i-1)
	}
	return calcMulAndDiv(exprSlice2, i+1)
}

func clacSumAndDif(expr []string, i int) []string {
	var exprSlice []string
	var num float32
	var exprSlice2 []string
	temp := 0
	exprSlice = expr
	fmt.Println(exprSlice)
	fmt.Println(exprSlice[i])
	if exprSlice[i] == "+" || exprSlice[i] == "-" {
		if exprSlice[i] == "+" {
			num = sum(convertStringToFloat(exprSlice[i-1]), convertStringToFloat(exprSlice[i+1]))
			temp = 1
		}
		if exprSlice[i] == "-" {
			num = dif(convertStringToFloat(exprSlice[i-1]), convertStringToFloat(exprSlice[i+1]))
			temp = 1
		}

		exprSlice2 = append(exprSlice2, exprSlice[:i-1]...)
		exprSlice2 = append(exprSlice2, convertFloatToString(num))
		if len(exprSlice)-2-1 >= i {
			exprSlice2 = append(exprSlice2, exprSlice[i+2:]...)
		}
		fmt.Println("zzzzz", exprSlice[i])
		fmt.Println(exprSlice[i-1], exprSlice[i], exprSlice[i+1])
	}

	if temp == 0 {
		exprSlice2 = exprSlice
	}

	fmt.Println("len(exprSlice2) : ", len(exprSlice2), ", exprSlice2 : ", exprSlice2, ", i : ", i, ", temp : ", temp)
	fmt.Println()
	if len(exprSlice2) == 1 || len(exprSlice2)-1 == i {
		fmt.Println("2------------------------------------------------------------2")
		return exprSlice2
	}
	// charList := strings.Join(exprSlice2, "")
	if temp == 1 {
		return clacSumAndDif(exprSlice2, i-1)
	}
	return clacSumAndDif(exprSlice2, i+1)
}

func convertStringToFloat(char string) float64 {
	if num, err := strconv.ParseFloat(char, 64); err == nil {
		return num
	}
	return 0
}

func convertFloatToString(num float32) string {
	return fmt.Sprintf("%f", num)
}

func sum(num1 float64, num2 float64) float32 {
	result := num1 + num2

	return round(result)
}

func dif(num1 float64, num2 float64) float32 {
	result := num1 - num2
	return round(result)
}

func mul(num1 float64, num2 float64) float32 {
	result := num1 * num2
	fmt.Println("::::::::::::::::::::::", num1, num2, result)
	return round(result)
}

func div(num1 float64, num2 float64) float32 {
	result := num1 / num2
	return round(result)
}

func round(num float64) float32 {
	return float32(int((1000000*num)+5)) / 1000000
}
