package student

//import "fmt"

func Len(str string) int {
	var len int
	for range str {
		len++
	}
	return len
}

func Power(nb int, power int) int {
	if power > 0 {
		return nb * Power(nb, power-1)
	}else if power == 0 {
		return 1
	}else {
		return 0
	}
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	lf, lt, ln := Len(baseFrom), Len(baseTo), Len(nbr)
	var dec = 0
	var nbr1, res, res1 string 
	for i := ln - 1; i >= 0; i-- {
		nbr1 += string(nbr[i])
	}
	for i, j := range nbr1 {
		for a, b := range baseFrom {
			if j == b {
				dec += a * Power(lf, i)
			}
		}
	}
	for i := dec; i > 0; i = i/lt {
		res += string(baseTo[i%lt])
	}
	for i := Len(res) - 1; i >= 0; i-- {
		res1 += string(res[i])
	}
	return res1
}