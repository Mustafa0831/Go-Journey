package student

func Compare(b, a string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	var len int = 0
	var char string
	for range array {
		len++
	}
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if f(array[j], array[j+1]) == 1 {
				char = array[j]
				array[j] = array[j+1]
				array[j+1] =  char
			}
		}
	}
}