package main

import (
	"fmt"
	"strings"

	"bufio"
	"os"
)

func main(){
	var n, valleys, ds int
	var s string

	reader := bufio.NewReader(os.Stdin)
	_, _ = fmt.Scanf("%d\r", &n)
	if n < 2 {
		fmt.Println("no")
		os.Exit(1)
	}

    s, _ = reader.ReadString('\n')
    s = strings.TrimSuffix(s, "\n")
    s = strings.TrimSpace(s)


	// s = strings.Replace(s, "U", " ", -1)
	// fmt.Println(s)
	path := strings.Split(s, " ")

	// fmt.Println(path)

    for i := 0; i < n; i++{
    	if path[i] == "D"{
    		ds++
    		if ds == 2{
    			valleys++
    		}
    	}else if(path[i] == "U"){
    		ds = 0
    	}
    }

    fmt.Println(valleys)

	// for _, element := range pairs{
		
	// }

	
}