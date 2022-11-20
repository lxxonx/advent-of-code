package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	input := "reyedfim"
	start := time.Now()
	output := make([]string, 8)
	found := 0

	for i := 0; found < 8; i++ {
		sum := md5.Sum([]byte(input + strconv.Itoa(i)))

		if sum[0] != 0 {
			continue
		}

		hashStr := fmt.Sprintf("%x", sum)
		if hashStr[:5] == "00000" {
			fmt.Print()
			pos := string(hashStr[5])
			posInt, err := strconv.Atoi(pos)

			if posInt < 8 && err == nil {
				if output[posInt] == "" {
					output[posInt] = string(hashStr[6])
					fmt.Println(output)
					found++
				}
			}
		}
	}

	end := time.Now()
	fmt.Printf("duration: %s \n", end.Sub(start))
	fmt.Println(strings.Join(output, ""))
}
