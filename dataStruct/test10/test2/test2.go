package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
假设表达式中允许包含两种括号：圆括号和方括号，其嵌套的顺序随意，如 ( [ ] ( ) ) 或 [ ( [ ] [ ] ) ] 等为正确的匹配，[ ( ] ) 或 ( [ ] ( ) 或 ( ( ) ) ) 均为错误的匹配。

现在的问题是，要求检验一个给定表达式中的括弧是否正确匹配？

输入一个只包含圆括号和方括号的字符串，判断字符串中的括号是否匹配，匹配就输出OK ，不匹配就输出 Wrong。
*/

type stack struct {
	data []byte
	size int
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("input failed")
	}
	s := getStack()
	var p byte
	for i := 0; i < len(input)-2; i++ {
		ch := input[i]
		if s.size == 0 {
			s.push(ch)
			continue
		}
		p = s.peek()
		if ch == ')' {
			if p == '(' {
				s.pop()
				continue
			} else {
				fmt.Println("Wrong")
				return
			}
		} else if ch == ']' {
			if p == '[' {
				s.pop()
				continue
			} else {
				fmt.Println("Wrong")
				return
			}
		} else {
			s.push(ch)
		}
	}
	if s.size == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("Wrong")
	}
}

func getStack() *stack {
	s := &stack{}
	s.data = make([]byte, 255)
	return s
}

func (s *stack) push(d byte) {
	if s.size >= 255 {
		return
	}
	s.data[s.size] = d
	s.size++
}

func (s *stack) pop() byte {
	if s.size == 0 {
		return 0
	}
	s.size--
	return s.data[s.size]
}

func (s *stack) peek() byte {
	return s.data[s.size-1]
}
