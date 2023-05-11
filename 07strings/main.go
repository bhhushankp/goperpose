package main

import (
	"fmt"
	"strings"
)

func main() {
	var pl = fmt.Println

	s1 := "My name is what"
	replacer := strings.NewReplacer("what", "tushar")
	s2 := replacer.Replace(s1)
	pl(s2)
	pl(len(s1))

	pl(strings.Contains(s1, "what"))
	pl(strings.Replace(s1, "a", "why", -1))
	pl("split :", strings.Split("a b c d", " "))
	pl("Uppercase :", strings.ToUpper((s1)))
	pl("Lowercase :", strings.ToLower(s1))
	pl("Prefix :", strings.HasPrefix(s1, "My"))
	pl("Sufix :", strings.HasSuffix(s1, "at"))

}
