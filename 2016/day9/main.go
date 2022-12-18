package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part1() {

	if input, err := ioutil.ReadFile("input.txt"); err == nil {
		start := 0
		flag := 1 // 1 = normal, 2 = decompressing, 3 = parsing number
		inputCopy := make([]byte, 0)

		length := 0
		repeat := 0

		for i := 0; i < len(input); i++ {
			if flag == 1 && string(input[i]) == "(" {
				start = i + 1
				flag = 3
				continue
			} else if flag == 3 && string(input[i]) == ")" {
				str := string(input[start:i])
				spl := strings.Split(str, "x")
				fmt.Sscanf(spl[0], "%d", &length)
				fmt.Sscanf(spl[1], "%d", &repeat)
				flag = 2
				continue
			}
			if flag == 2 {
				repeated := strings.Repeat(string(input[i:i+length]), repeat)
				inputCopy = append(inputCopy, []byte(repeated)...)
				flag = 1
				i += length - 1
				continue
			}
			if flag == 3 {
				continue
			}

			inputCopy = append(inputCopy, input[i])

		}

		fmt.Println(len(inputCopy))
	}
}

func decompress(str string) {
	count := 0
	start := 0
	flag := 0 // 0 = normal, 1 = decompressing, 2 = parsing number
	length := 0
	repeat := 0
	for i := 0; i < len(str); i++ {

		if flag == 0 && string(str[i]) == "(" {
			start = i + 1
			flag = 2
			continue
		}
		if flag == 0 {
			count++
			continue
		}
		if flag == 2 && string(str[i]) == ")" {
			str := string(str[start:i])
			spl := strings.Split(str, "x")
			fmt.Sscanf(spl[0], "%d", &length)
			fmt.Sscanf(spl[1], "%d", &repeat)
			flag = 1
			continue
		}
		if flag == 1 {
			decompress(string(str[i : i+length]))
			flag = 0
			i += length - 1
			continue
		}
	}
}

func part2() {
	input, _ := ioutil.ReadFile("input.txt")
	start := 0
	flag := 1 // 1 = normal, 2 = decompressing, 3 = parsing number
	inputCopy := make([]byte, 0)

	length := 0
	repeat := 0
	for i, c := range input {
		if flag == 1 && string(input[i]) == "(" {
			start = i + 1
			flag = 3
			continue
		}
		if flag == 3 && string(input[i]) == ")" {
			str := string(input[start:i])
			spl := strings.Split(str, "x")
			fmt.Sscanf(spl[0], "%d", &length)
			fmt.Sscanf(spl[1], "%d", &repeat)
			flag = 2
			continue
		}

		if flag == 2 {
			str := string(input[i : i+length])
			repeated := strings.Repeat(str, repeat)
			inputCopy = append(inputCopy, []byte(repeated)...)
			flag = 1
			i += length - 1
			continue
		}
		if flag == 3 {
			continue
		}
	}
}

func main() {
	part2()
}
