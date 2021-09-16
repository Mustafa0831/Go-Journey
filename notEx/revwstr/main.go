package main

import (
	"os"

	"github.com/01-edu/z01"
)

// func main() {
// 	defer z01.PrintRune('\n')
// 	if len(os.Args)==2 {
// 		var res string
// 		arg:=os.Args[1]
// 		for i:=len(arg)-1; i>=0; i-- {
// 			if arg[i]==' ' {
// 				res +=arg[i+1:] + " "
// 				arg = arg[:i]
// 			}
// 		}
// 		res+= arg
// 		for _,r :=range res{
// 			z01.PrintRune(r)
// 		}
// 	}
// }
func main() {
	if len(os.Args)==2 {
		  arg :=os.Args[1]
		//   arr := []rune(arg)
		  res :=""
		  for i:=len(arg)-1; i>=0; i--{
			  if arg[i]==' ' {
				  res +=arg[i+1:] + " "
				  arg = arg[:i]
			  }
		  }
		  res +=arg
		  for _,r:=range res {
			  z01.PrintRune(r)
		  }
		  z01.PrintRune(10)
	}
}