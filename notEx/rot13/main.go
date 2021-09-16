package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	if len(os.Args)==2 {
// 		for _, r:= range os.Args[1] {
// 			if r>= 65 && r<= 77|| r>= 97 && r<= 109 {
// 				z01.PrintRune(r +13)
// 			}else if r >=78 && r <= 90|| r>= 110 && r<= 122 {
// 				z01.PrintRune(r -13)
// 			}else {
// 				z01.PrintRune(r)
// 			}
// 		}
// 		if len(os.Args[1]) != 0 {
// 			z01.PrintRune(10)
// 		}
// 	}
// }

func main() {
	if len(os.Args)==2 {
		for _, r:=range os.Args[1] {
			if r >= 'A' && r <= 'N' || r >= 'a' && r <='n' {
				z01.PrintRune(r+13)
			}else if r >='M' && r <= 'Z' || r>= 'm' && r<= 'z' {
				z01.PrintRune(r-13)
			}else {
				z01.PrintRune(r)
			}
		}
		z01.PrintRune(10)
	}
}
