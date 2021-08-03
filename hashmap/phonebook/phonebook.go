package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()

	stringN := scanner.Text()
	n, err := strconv.Atoi(stringN)
	if err != nil {
		fmt.Println("not a number")
		return
	}

	mapPhoneNumber := make(map[string]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		temp := scanner.Text()
		record := strings.Split(temp, " ")
		mapPhoneNumber[record[0]] = record[1]
	}

	input := "not empty"
	for scanner.Scan() {
		input = scanner.Text()
		if input == "" {
			break
		}
		if mapPhoneNumber[input] != "" {
			fmt.Println(input + "=" + mapPhoneNumber[input])
		} else {
			fmt.Println("Not found")
		}
	}
}
