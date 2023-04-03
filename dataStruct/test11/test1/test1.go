package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

/*
将一对小括号()，插入到一个括号序列中，其中插入的规则是，左括号'('的位置要小于右括号')'的位置，不要求插入的左右括号相邻，
例如以下，为了方便区分，我们拿ab代表原括号序列
将()插入到()中可形成()ab (a)b (ab) a()b a(b) ab()等等序列，其中a代表原括号序列的左括号,b代表原括号序列的右括号。
小明认为一个括号序列是一个优美的序列当且仅当这个括号序列可以被如下方法构造出来：
一开始有一个空串，然后执行0次或者若干次操作，每次操作将()插入到当前的括号序列中。
根据上面的定义：() , (()) , (()())都是优美的括号序列，(() , )( , ()))都不是优美的括号序列

输入格式：
多组输入
每行输入给定一个仅由'(',')'组成的括号序列,长度小于等于1000
题目保证没有空串

输出格式：
对于每个输入输出一行，若当前的括号序列是优美的，则输出"YES"（不含引号）
否则输出"NO"（不含引号）
2023年4月3日14:32:59
*/

func main() {
	input := bufio.NewReader(os.Stdin)
	for {
		line, _, err := input.ReadLine()
		if err != nil || len(line) == 0 {
			return
		}
		stack := &list.List{}
		flag := true
		var ch byte
		for i := 0; i < len(line); i++ {
			ch = line[i]
			if ch == ')' {
				if stack.Len() == 0 || stack.Back().Value.(byte) != '(' {
					flag = false
					break
				} else {
					stack.Remove(stack.Back())
				}
			} else {
				stack.PushBack(ch)
			}
		}
		if stack.Len() != 0 {
			flag = false
		}
		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
