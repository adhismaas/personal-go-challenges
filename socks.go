package main

import (
	"fmt"
	"strings"

	"bufio"
	"os"
)

func main(){
	var socks, ready int
	var input string
	pairs := make(map[string]int)

	reader := bufio.NewReader(os.Stdin)
	_, _ = fmt.Scanf("%d\r", &socks)
    input, _ = reader.ReadString('\n')
    input = strings.TrimSuffix(input, "\n")

	bucket := strings.Split(input, " ")

    for i := 0; i < socks; i++{
		_, exist := pairs[bucket[i]]
		if exist{
			pairs[bucket[i]] += 1
		} else{
			pairs[bucket[i]] = 1
		}
	}

	for _, element := range pairs{
		ready += int(element/2)
	}

	fmt.Println(int(ready))
	
}