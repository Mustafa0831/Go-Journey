package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	if len(os.Args)==3 {
// 		if os.Args[1]> os.Args[2] {
// 			z01.PrintRune('1')
// 		}else if os.Args[1]== os.Args[2] {
// 			z01.PrintRune('0')
// 		}else {
// 			for _, r:=range "-1"{
// 				z01.PrintRune(r)
// 			}
// 		}
// 		z01.PrintRune(10)
// 	}
// }
func main() {
	if len(os.Args)==3 {
		if os.Args[1] > os.Args[2]{
			z01.PrintRune('1')
		}else if os.Args[1]==os.Args[2] {
			z01.PrintRune('0')
		}else {
			for _, r:=range "-1" {
				z01.PrintRune(r)
			}
		}
		z01.PrintRune(10)
 	}
}