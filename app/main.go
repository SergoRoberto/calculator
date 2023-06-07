package main

import (
	"app/calculate"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	fmt.Println(calculator("12332432 * 124543123, 2 + 2, 3 - 3, 4 * 3, 12 / 3"))

}

func calculator(str string) string {

	exps := strings.Split(str, ", ")
	res := make(chan string)
	for _, exp := range exps {
		sym := strings.Split(exp, " ")
		x, err := strconv.Atoi(sym[0])
		if err != nil {
			return "error"
		}
		y, err := strconv.Atoi(sym[2])
		if err != nil {
			return "error"
		}

		switch sym[1] {
		case "+":
			go calculate.Add(res, x, y)
		case "-":
			go calculate.Subtract(res, x, y)
		case "*":
			go calculate.Multiply(res, x, y)
		case "/":
			go calculate.Divide(res, x, y)
		}
	}
	result := ""
	for i := range exps {
		result += <-res
		if i != len(exps)-1 {
			result += ", "
		}
	}

	return result
}
