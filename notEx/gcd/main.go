package main

import (
	"fmt"
	"os"
	"strconv"
)

// func main() {
// 	if len(os.Args) == 3 {
// 		var res int
// 		num1, _ := strconv.Atoi(os.Args[1])
// 		num2, _ := strconv.Atoi(os.Args[2])
// 		for i := num2; i > 0; i-- {
// 			if num2%i == 0 && num1%i == 0 {
// 				res = i
// 				break
// 			}
// 		}
// 		fmt.Println(res)
// 		return
// 	}
// }

// func Itoa(num int) {
// 	if num == 0 {
// 		return
// 	}
// 	Itoa(num/10)
// 	z01.PrintRune(rune(num%10)+'0')
// 	z01.PrintRune('\n')
// }

func main() {
	if len(os.Args)==3 {
	var res int
	num1, _:=strconv.Atoi(os.Args[1])
	num2, _:=strconv.Atoi(os.Args[2])
	for i:=num2; i>=0; i-- {
		if num2%i == 0 && num1%i == 0 {
			res =i
			break
		}
	}
	fmt.Println(res)
	}
}
