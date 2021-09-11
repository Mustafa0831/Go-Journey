package main

import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	if len(os.Args)==2 {
		num, _ :=strconv.Atoi(os.Args[1])
		div:=2
		for div<=num {
			if num%div==0 {
				fmt.Print(div)
				if num==div {
					fmt.Println()
					return
				}
				fmt.Print("*")
				num /=div
				div =1
			}
			div++
		}
	}
}