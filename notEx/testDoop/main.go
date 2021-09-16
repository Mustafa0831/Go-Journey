package main

import (
	"fmt"
	"os"
)

// func main() {
// 	if len(os.Args) == 4 {
// 		var res int
// 		num1 := Atoi(os.Args[1])
// 		num2 := Atoi(os.Args[3])
// 		if num1 == 0 {
// 			fmt.Print()
// 			return
// 		}
// 		if os.Args[2] != "+" && os.Args[2] != "-" && os.Args[2] != "*" && os.Args[2] != "/" && os.Args[2] != "%" {
// 			fmt.Print()
// 			return
// 		}
// 		if os.Args[2] == "+" {
// 			res = num1 + num2
// 			if num1 > 0 && num2 > 0 && res < 0 {
// 				fmt.Print()
// 				return
// 			}
// 		}
// 		if os.Args[2] == "-" {
// 			res = num1 - num2
// 			if num1 < 0 && num2 < 0 && res < 0 {
// 				fmt.Print()
// 				return
// 			}
// 		}
// 		if os.Args[2] == "*" {
// 			res = num1 * num2
// 			if res/num1 != num2 {
// 				fmt.Print()
// 				return
// 			}
// 		}
// 		if os.Args[2] == "/" {
// 			if num2 == 0 {
// 				fmt.Println("nO DIVISION  by 0")
// 				return
// 			} else {
// 				res = num1 / num2
// 			}
// 		}
// 		if os.Args[2] == "%" {
// 			if num2 == 0 {
// 				fmt.Println("nO modulo  by 0")
// 				return
// 			} else {
// 				res = num1 % num2
// 			}
// 		}
// 		fmt.Println(res)
// 	}
// }

// func OverFlow(str string) bool {
// 	flag := false
// 	if len(str) >= 20 && str[0] == '-' {
// 		flag = true
// 		str = str[1:]
// 	}
// 	if len(str) > 19 {
// 		return false
// 	} else if len(str) == 19 {
// 		if flag && str > "9223372036854775808" || !flag && str > "9223372036854775807" {
// 			return false
// 		}
// 	}
// 	return true
// }

// func Atoi(str string) int {
// 	if OverFlow(str) {
// 		res, sign := 0, 1
// 		for _, r := range str {
// 			count := 0
// 			if r == '-' {
// 				sign = -1
// 			}
// 			for i := '0'; i < r; i++ {
// 				count++
// 			}
// 			res = res*10 + count
// 			for _, r := range []rune(str) {
// 				r -= '0'
// 				if r > '9' {
// 					return 0
// 				}
// 			}
// 		}
// 		return res * sign
// 	}
// 	return 0
// }

func main() {
	if len(os.Args) == 4 {
		res := 0
		num1 := Atoi(os.Args[1])
		num2 := Atoi(os.Args[3])
		switch os.Args[2] {
		case "+":
			res = num1 + num2
			if res < 0 && num1 > 0 && num2 > 0 {
				return
			}
		case "-":
			res = num1 - num2
			if res < 0 && num1 < 0 && num2 < 0 {
				return
			}

		case "*":
			res = num1 * num2
			if res/num1 != num2 {
				return
			}
		case "/":
			if num2 ==0 {
				fmt.Println("")
			}
		case "%":
			res = num1 + num2
			if res < 0 && num1 > 0 && num2 > 0 {
				return
			}
		}
	}
}
func OverFlow(str string) bool {
	flag := false
	if len(str) >= 20 && str[0] == '-' {
		str = str[1:]
		flag = true
	}
	if len(str) > 19 {
		return false
	} else if len(os.Args) == 19 {
		if flag && str > "8" || !flag && str > "7" {
			return false
		}
	}
	return true
}

func Atoi(str string) int {
	if OverFlow(str) {
		res, sign := 0, 1
		for _, r := range str {
			count := 0
			if r == '-' {
				sign = -1
			}
			for i := '0'; i < r; i++ {
				count++
			}
			res = res*10 + count
		}
		for _, r := range []rune(str) {
			r -= '0'
			if r > '9' {
				return 0
			}
		}
		return res * sign
	}
	return 0
}
