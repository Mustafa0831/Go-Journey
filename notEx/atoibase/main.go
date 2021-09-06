package main

import "fmt"

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

func CheckBase(base string)bool {
	var newstr string
	if len(base) <2 {
		return false
	}
	for _, r:=range base {
		if r == '+'|| r== '-' {
			return false
		}
		for _, v :=range newstr {
			if r==v {
				return false
			}
		}
	}
	return true
}

func AtoiBase(str, base string)(res int) {
	if !CheckBase(base) {
		return 0
	}else {
		for i :=range str {
			for j:=range base {
				if str[i]== base[j] {
					res = res*len(base) +j
				}
			}
		}
	}
	return
}