package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	if len(os.Args)==3 {
// 		arg1:=os.Args[1]
// 		arg2:=os.Args[2]
// 		i:=0
// 		for j:=range arg2 {
// 			if arg1[i]==arg2[j] {
// 				i++
// 				if len(arg1)==i {
// 					z01.PrintRune('1')
// 					z01.PrintRune(10)
// 					return
// 				}
// 			}
// 		}
// 		z01.PrintRune('0')
// 		z01.PrintRune(10)
// 	}
// }

func main() {
	if len(os.Args) == 3 {
		arg1 :=os.Args[1]
		arg2 :=os.Args[2]
		i:=0
		for j:=range arg2 {
			if arg1[i]==arg2[j] {
				i++
				if i == len(arg1) {
					z01.PrintRune('1')
					z01.PrintRune(10)
					return
				}
			}
		}
		z01.PrintRune('0')
		z01.PrintRune(10)
	}
}
