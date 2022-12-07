package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Queue []string

func (q *Queue) Push(v string) {
	*q = append(*q, v)
}

func (q *Queue) Pop() string {
	if len(*q) == 0 {
		return ""
	}
	res := (*q)[0]
	*q = (*q)[1:]
	return res
}

func (q *Queue) hasDuplications() bool {
	for i := 0; i < len(*q); i++ {
		for j := i + 1; j < len(*q); j++ {
			if (*q)[i] == (*q)[j] {
				return true
			}
		}
	}
	return false
}

func main() {
	start := time.Now()
	input, _ := ioutil.ReadFile("input.txt")

	chars := strings.Split(string(input), "")

	queue := Queue{}
	for i, char := range chars {
		if len(queue) < 14 {
			queue.Push(char)
			continue
		}
		queue.Pop()
		if queue.hasDuplications() {
			queue.Push(char)
			continue
		}
		str := strings.Join(queue, "")
		if strings.Contains(str, char) {
			queue.Push(char)
			continue
		}
		fmt.Println(str, i, char)
		break
	}

	end := time.Now()
	fmt.Println(end.Sub(start).String())
}
