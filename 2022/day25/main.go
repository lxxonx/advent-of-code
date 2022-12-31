package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func LargestDigit(num, exp int) int {
	pow := int(math.Pow(5, float64(exp)))
	if num < pow*2 {
		return exp
	}

	return LargestDigit(num, exp+1)
}

func SNAFU(num int, exp int, res string) (int, string) {
	thisPow := int(math.Pow(5, float64(exp)))
	nextPow := int(math.Pow(5, float64(exp+1)))
	mod := num % nextPow

	if num == 0 {
		return 0, res
	}

	fmt.Println(num, mod, thisPow, res)

	if mod == 4*thisPow {
		res = "-" + res
		num += thisPow
	} else if mod == 3*thisPow {
		res = "=" + res
		num += thisPow * 2
	} else if mod == 2*thisPow {
		res = "2" + res
		num -= thisPow * 2
	} else if mod == thisPow {
		res = "1" + res
		num -= thisPow
	} else if mod == 0 {
		val := mod / thisPow
		res = strconv.Itoa(val) + res
		num -= thisPow * val
	}

	return SNAFU(num, exp+1, res)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(input), "\n")

	var sum int

	sum = 0

	for _, line := range lines {
		max := len(line) - 1

		chars := strings.Split(line, "")

		for i, char := range chars {
			switch char {
			case "1":
				sum += int(math.Pow(5, float64(max-i)))
				break
			case "2":
				sum += int(math.Pow(5, float64(max-i)) * 2)
				break
			case "=":
				sum += int(math.Pow(5, float64(max-i)) * -2)
				break
			case "-":
				sum += int(math.Pow(5, float64(max-i)) * -1)
				break
			}
		}
	}

	_, str := SNAFU(sum, 0, "")

	fmt.Println(str)
}
