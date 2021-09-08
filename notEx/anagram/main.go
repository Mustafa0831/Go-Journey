package main

import "fmt"

func main() {
	test1 := IsAnagram("listen", "silent")
	fmt.Println(test1)

	test2 := IsAnagram("alem", "school")
	fmt.Println(test2)

	test3 := IsAnagram("neat", "a net")
	fmt.Println(test3)

	test4 := IsAnagram("anna madrigal", "a man and a girl")
	fmt.Println(test4)
}

// func IsAnagram(str1, str2 string) bool {
// 	res := rune(0)
// 	for _, r := range str1 + str2 {
// 		if r == ' ' {
// 			continue
// 		}
// 		res ^= rune(r)
// 	}
// 	return res == 0
// }
func IsAnagram(str1, str2 string) bool {
	res :=rune(0)
	for _, r:=range str1+str2 {
		if r == ' ' {
			continue
		}
		res ^= rune(r)
	}
	return res ==0
}