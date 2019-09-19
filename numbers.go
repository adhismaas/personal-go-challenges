package main

import (
	"fmt"
	"strings"

	"bufio"
	"os"
)

func main(){
	var numbers, zeroes string
	var s []string

	reader := bufio.NewReader(os.Stdin)
    numbers, _ = reader.ReadString('\n')
    numbers = strings.TrimSuffix(numbers, "\n")

    length := len(numbers)
    length -= 2
    numzero := length

	for _, r := range numbers{
		zeroes = ""
		for n:=0; n<numzero; n++{
			zeroes += "0"
		}
		numzero -= 1

		s = append(s, string(r) + zeroes)
	}

	for i := 0; i <= length; i++{
		fmt.Println(s[i])
	}
}