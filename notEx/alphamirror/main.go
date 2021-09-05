package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	if len(os.Args) == 2 {
// 		for _, r := range os.Args[1] {
// 			if r >= 65 && r <= 90 {
// 				z01.PrintRune(90 - (r - 65))
// 			} else if r >= 97 && r <= 122 {
// 				z01.PrintRune(122 - (r - 97))
// 			} else {
// 				z01.PrintRune(r)
// 			}
// 		}
// 		z01.PrintRune(10)
// 	}
// }

func main() {
	if len(os.Args)==2 {
		for _, r:=range os.Args[1] {
			if r >= 65 && r <= 90 {
				z01.PrintRune(90 - (r-65))
			}else if r >= 97 && r <= 122 {
				z01.PrintRune(122 - (r-97))
			}else {
				z01.PrintRune(r)
			}
		}
		z01.PrintRune(10)
	}
}


