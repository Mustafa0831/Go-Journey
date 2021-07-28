package student

func AtoiBase(s string, base string) int {
	var res int
	if !CheckBase(base) {
		return 0
	}else {
		for i1:= range s {
			for i2 := range base {
				if s[i1] == base[i2] {
					res = res * len(base) + i2
				}
			}
		}
	}
	return res
}

func len(base string) int {
	length := 0
	for range base {
		length++
	}
	return length
}

func CheckBase(base string) bool {
	var newstr string
	if len(base) < 2 {
		 return false
	}
	for _, l1:= range base {
		if l1 == '-'|| l1 == '+' {
			return false
		}
		for _, l2:= range newstr {
			if l1 == l2 {
				return false
			}
		}
		newstr += string(l1)
	}
	return true
}