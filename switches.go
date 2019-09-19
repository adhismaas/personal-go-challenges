package main

import (
	"fmt"
	// "strings"

	// "bufio"
	// "os"
)

func main(){
	var switches[100]int
	var on int

	for i:=0; i<100; i++{
		for n:=0; n<len(switches); n++{
			if (n+1)%(i+1) == 0{
				if switches[n] == 0{
					switches[n] = 1
				} else{
					switches[n] = 0
				}
			}
		}

		for count:=0; count<len(switches); count++{
			if switches[count] == 1{
				on++
			}
		}

		fmt.Println("iteration ", i+1, "\t | switches on : ", on)
		on = 0
	}
}