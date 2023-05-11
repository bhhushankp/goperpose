package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var pl = fmt.Println
	var pf = fmt.Printf
	rStr := "abcdefg"
	pl("Rune Count :", utf8.RuneCountInString(rStr))
	for i, runeVal := range rStr {
		pf("%d : %#U : %c :\n", i, runeVal, runeVal)
	}
}
