package main

import "fmt"

func main() {
	fmt.Println(Itoa(-42))
}

// func Itoa(num int) (res string) {
// 	flag := false
// 	if num < 0 {
// 		num = -num
// 		flag = true
// 	}
// 	for num > 0 {
// 		res = string(num%10+'0') + res
// 		num /= 10
// 	}
// 	if flag {
// 		res = "-" + res
// 	}
// 	return
// }

func Itoa(num int)(res string) {
	flag :=false
	if num<0 {
		num = -num
		flag = true
	}
	for num >0 {
		res = string(num%10+'0') +res
		num /=10
	}
	if flag {
		res = "-" +res
	}
	return
}
