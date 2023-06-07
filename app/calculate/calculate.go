package calculate

import (
	"strconv"
)

func Add(res chan string, x int, y int) {
	res <- strconv.Itoa(x) + " + " + strconv.Itoa(y) + " = " + strconv.Itoa(x+y)
}

func Subtract(res chan string, x int, y int) {
	res <- strconv.Itoa(x) + " - " + strconv.Itoa(y) + " = " + strconv.Itoa(x-y)
}

func Multiply(res chan string, x int, y int) {
	res <- strconv.Itoa(x) + " * " + strconv.Itoa(y) + " = " + strconv.Itoa(x*y)
}

func Divide(res chan string, x int, y int) {
	res <- strconv.Itoa(x) + " / " + strconv.Itoa(y) + " = " + strconv.Itoa(x/y)
}
