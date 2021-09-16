package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// defer z01.PrintRune(10)
// if len(os.Args) == 2 {
// 	arg := os.Args[1]
// 	var newstr string
// 	for i, r := range arg {
// 		if r != ' ' {
// 			newstr += string(r)
// 		} else if i != 0 && arg[i-1] != ' ' && r == ' ' {
// 			newstr += string(r)
// 		}
// 	}
// 	var firstword string
// 	flag := true
// 	count := 0
// 	for _, r := range newstr {
// 		if r == ' ' {
// 			flag = false
// 		}
// 		if flag == true {
// 			firstword += string(r)
// 			count++
// 		}
// 	}
// //Printing whole string except firstword
// for i := count + 1; i < len(newstr); i++ {
// 	z01.PrintRune(rune(newstr[i]))
// }
// //Printing space before firstword
// if arg[len(arg)-1] != ' ' {
// 	z01.PrintRune(' ')
// }
//  //Printing firstword
// for _, r := range firstword {
// 	z01.PrintRune(r)
// }
// 	}
// }
func main() {
	defer z01.PrintRune(10)
	if len(os.Args)==2 {
		var newstr string
		arg := os.Args[1]
		for i, r:=range os.Args[1] {
			if r != ' ' {
				newstr += string(r)
			}else if i != 0 && arg[i-1] != ' ' && r == ' ' {
				newstr +=string(r)
			}
		}
		firstword := ""
		count:=0
		flag :=false
		for _,r:=range newstr {
			if r == ' ' {
				flag = true
			}
			if flag == false {
				count++
				firstword +=string(r)
			}
		}
		for i:=count+1; i<len(newstr); i++ {
			z01.PrintRune(rune(newstr[i]))
		}
		if arg[len(arg)-1] != ' ' {
			z01.PrintRune(' ')
		}
		for _, r:=range firstword {
			z01.PrintRune(r)
		}
	}
}
