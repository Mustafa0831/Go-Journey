package main

import "fmt"

func main() {
	res := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	SortWordArr(res)
	fmt.Println(res)
}

// func SortWordArr(arr []string) {
// 	for i, _ :=range arr {
// 		for j:=i+1; j<len(arr); j++ {
// 			if arr[i]> arr[j] {
// 				arr[i], arr[j]= arr[j], arr[i]
// 			}
// 		}
// 	}
// }

func SortWordArr(arr []string) {
	for i, _:= range arr {
		for j:= i+1; j<len(arr); j++ {
			if arr[i]> arr[j] {
				arr[i], arr[j]= arr[j], arr[i]
			}
		}
	}
}