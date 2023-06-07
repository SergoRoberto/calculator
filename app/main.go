package main

import (
	"app/calculate"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	calculator("2 + 2, 3 - 3, 4 * 3, 12 / 3")

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
			go calculate.Add(x, y)
		case "-":
			go calculate.Subtract(x, y)
		case "*":
			go calculate.Multiply(x, y)
		case "/":
			go calculate.Divide(x, y)
		}
		fmt.Println(sym[1])
	}

	return <-res
}
