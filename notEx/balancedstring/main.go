package main

import (
	"fmt"
	"os"
)

// func main() {
// 	if len(os.Args) == 2 {
// 		var res, count int
// 		for _, r := range os.Args[1] {
// 			if r == 'C' {
// 				count++
// 			} else {
// 				count--
// 			}
// 			if count == 0 {
// 				res++
// 			}
// 		}
// 		fmt.Println(res)
// 		return
// 	}
// }
func main() {
	if len(os.Args) == 2 {
		var res, count int
		for _, r:=range os.Args[1] {
			if r== 'C' {
				count++
			}else {
				count--
			}
			if count==0 {
				res++
			}
		}
		fmt.Println(res)
	}
}
