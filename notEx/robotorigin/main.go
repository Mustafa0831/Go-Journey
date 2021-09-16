package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main(){
	defer z01.PrintRune(10)
	if len(os.Args[1])==2 {
		res := "false"
		var way int
		for _, r:= range os.Args[1] {
			if r == 'U' || r == 'R' {
				way++
			}
			if r == 'D' || r == 'L' {
				way--
			}
		}
		if way == 0 {
			res = "true"
		}
		for _, r:=range res{
			z01.PrintRune(r)
		}
	}
}