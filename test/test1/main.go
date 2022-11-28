package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var result string
	reader := bufio.NewReader(os.Stdin)
	s, _, _ := reader.ReadLine()
	strLen := len(s)
	for i := 0; i < strLen; i++ {
		ch := s[i]
		if ch >= 'A' && ch <= 'Z' {
			ch = 'A' + (ch-'A'+4)%26 - 1
		} else if ch >= 'a' && ch <= 'z' {
			ch = 'a' + (ch-'a'+4)%26 - 1
		}
		result += string(ch)
	}
	fmt.Printf("%s\n", result)

}
