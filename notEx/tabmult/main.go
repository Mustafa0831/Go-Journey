package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args)==2 {
		var res string
		num,_ := strconv.Atoi(os.Args[1])
		for i:=1; i<=9; i++ {
			res+=Itoa(i) + " x " + Itoa(num) + " = " + Itoa(num*i) + "\n" 
		}
		for _,r:=range res {
			z01.PrintRune(r)
		}
	}
}

func Itoa(num int)(res string) {
	for num >0 {
		res = string(num%10+'0') +res
		num /=10
	}
	return
}