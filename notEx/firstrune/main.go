package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(FirstRune("Hello!"))
	z01.PrintRune(FirstRune("Salut!"))
	z01.PrintRune(FirstRune("Ola!"))
	z01.PrintRune('\n')
}

func FirstRune(str string) rune{
	runes := []rune(str)
	return runes[0]
}
