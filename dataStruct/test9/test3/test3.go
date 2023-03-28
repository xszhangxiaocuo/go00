package main

import (
	"bufio"
	"dataStruct/test9/test3/myStack"
	"fmt"
	"os"
)

/*
编写一个简单的逆波兰式计算器，它接受用户输入的整型数和运算符 +、-、*、/。
2023年3月28日15:27:59
*/

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("read failed!")
	}
	cal(input)
}

func cal(str string) {
	num := myStack.NewMySack()  //存放数的栈
	oper := myStack.NewMySack() //存放计算符的栈
	n := 0
	i := 0
	for i < len(str)-2 {
		ch := str[i]
		if ch >= '0' && ch <= '9' {
			n = n*10 + int(ch-'0')
			i++
		} else if isOper(ch) {
			i++
			num.Push(n)
			n = 0
			pop := oper.GetPop()
			if pop != nil && isHigher(pop.(byte), ch) { //如果oper栈不为空并且栈顶运算符的优先级大于等于当前运算符就计算一次
				count(num, oper)
				oper.Push(ch)
				continue
			}
			oper.Push(ch)

		}
	}
	num.Push(n)
	for oper.GetSize() != 0 {
		count(num, oper)
	}
	fmt.Println(num.Pop())
}

//判断是否是运算符
func isOper(ch byte) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

//判断pop的优先级是否高于ch,若栈中运算符pop为‘*’或‘/’优先级肯定大于等于当前运算符ch，进行一次运算，同理ch为‘+’或‘-’时同样满足条件
func isHigher(pop byte, ch byte) bool {
	return (pop == '*' || pop == '/') || (ch == '+' || ch == '-')
}

//进行一次运算
func count(num *myStack.MyStack, oper *myStack.MyStack) {
	num2 := num.Pop().(int)
	num1 := num.Pop().(int)
	o := oper.Pop().(byte)
	switch o {
	case '+':
		num.Push(num1 + num2)
	case '-':
		num.Push(num1 - num2)
	case '*':
		num.Push(num1 * num2)
	case '/':
		num.Push(num1 / num2)
	}
}
