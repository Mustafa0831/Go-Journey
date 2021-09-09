package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 3 {
		mp := map[rune]int{}
		for _, r := range os.Args[1] {
			for _, v := range os.Args[2] {
				if r == v && mp[r] < 1 {
					mp[r]++
					z01.PrintRune(r)
				}
			}
		}
		z01.PrintRune(10)
		return
	}
}
