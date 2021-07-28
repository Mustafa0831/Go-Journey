package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)	

func PrintString(str string) {
	for _, j := range str {
		z01.PrintRune(j)
	}
}

func main() {
	lenargs := 0
	for range os.Args {
		lenargs++
	}
	if lenargs > 1 {
		for i := 1; i < lenargs; i++ {
			file, err := ioutil.ReadFile(os.Args[i])
			if err != nil {
				msg := "cat: " + os.Args[i] + ": No such file or direcory"
				PrintString(msg)
				z01.PrintRune('\n')
				continue
			}
			PrintString(string(file))
		}
	}
	if lenargs == 1 {
		io.Copy(os.Stdin, os.Stdout)
	}
}