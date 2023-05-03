package main

import (
	"container/list"
	"dataStruct/test16/test1/btnode"
	"fmt"
)

var index = 0

func main() {
	root := creatTree("-+a *b -c d /e f ")
	//fmt.Print("前序遍历：")
	//preorder(root)
	//fmt.Println()
	//fmt.Print("中序遍历：")
	//inorder(root)
	//fmt.Println()
	//fmt.Print("后序遍历：")
	//postorder(root)
	//fmt.Println()
	//
	//leaf := countleaf(root)
	//fmt.Println("叶子数：", leaf)
	//
	//h := high(root)
	//fmt.Println("深度：", h)
	//
	//fmt.Print("带括号的中缀表达式：")
	//infixExpression(root)
	//fmt.Println()

	fmt.Print("非递归的中序遍历：")
	nonRecInorder(root)
	fmt.Println()

	fmt.Print("复制二叉树并中序遍历输出：")
	copyRoot := copyBtree(root)
	inorder(copyRoot)
	fmt.Println()

	leaf := nonRecCountLeaf(root)
	fmt.Println("叶子数：", leaf)

	n := countSingleC(root)
	fmt.Println("度为1的节点数：", n)

	fmt.Print("按层遍历：")
	layerorder(root)
}

//传入一个先序遍历序列
func creatTree(str string) *btnode.Btnode {
	if str[index] == ' ' || index >= len(str) {
		return nil
	}
	node := &btnode.Btnode{Data: str[index]}
	index++
	left := creatTree(str)
	if left == nil {
		return node
	}
	index++
	right := creatTree(str)
	node.Left = left
	node.Right = right
	return node
}

//前序遍历
func preorder(root *btnode.Btnode) {
	if root == nil {
		return
	}
	fmt.Printf("%c", root.Data)
	preorder(root.Left)
	preorder(root.Right)
}

//中序遍历
func inorder(root *btnode.Btnode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Printf("%c", root.Data)
	inorder(root.Right)
}

//后序遍历
func postorder(root *btnode.Btnode) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	fmt.Printf("%c", root.Data)
}

//求叶子数
func countleaf(root *btnode.Btnode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := countleaf(root.Left)
	right := countleaf(root.Right)
	return left + right
}

//求深度
func high(root *btnode.Btnode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	left := high(root.Left) + 1
	right := high(root.Right) + 1
	if left > right {
		return left
	} else {
		return right
	}
}

//输出带括号的中缀表达式
func infixExpression(root *btnode.Btnode) {
	if root == nil {
		return
	} else if root.Left == nil || root.Right == nil {
		fmt.Printf("%c", root.Data)
		return
	}
	fmt.Print("(")
	infixExpression(root.Left)
	fmt.Printf("%c", root.Data)
	infixExpression(root.Right)
	fmt.Print(")")
}

//非递归的中序遍历
func nonRecInorder(root *btnode.Btnode) {
	stack := []*btnode.Btnode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]   //取出栈顶元素
		stack = stack[:len(stack)-1] //弹出栈顶元素
		fmt.Printf("%c", root.Data)
		root = root.Right
	}
}

//复制二叉树
func copyBtree(root *btnode.Btnode) *btnode.Btnode {
	if root == nil {
		return nil
	}
	copyRoot := &btnode.Btnode{
		Data:  root.Data,
		Left:  root.Left,
		Right: root.Right,
	}
	left := copyBtree(root.Left)
	right := copyBtree(root.Right)
	copyRoot.Left = left
	copyRoot.Right = right
	return copyRoot
}

//非递归计算叶子数
func nonRecCountLeaf(root *btnode.Btnode) int {
	stack := []*btnode.Btnode{}
	count := 0
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]   //取出栈顶元素
		stack = stack[:len(stack)-1] //弹出栈顶元素
		if root.Left == nil && root.Right == nil {
			count++
		}
		root = root.Right
	}
	return count
}

//用递归计算度为1的节点数
func countSingleC(root *btnode.Btnode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right == nil || root.Right != nil && root.Left == nil {
		return 1
	}

	left := countSingleC(root.Left)
	right := countSingleC(root.Right)

	return left + right
}

//按层遍历二叉树
func layerorder(root *btnode.Btnode) {
	queue := &list.List{}
	queue.PushBack(root)
	for queue.Len() > 0 {
		tmp := queue.Front().Value.(*btnode.Btnode)
		queue.Remove(queue.Front())
		fmt.Printf("%c", tmp.Data)
		if tmp.Left != nil {
			queue.PushBack(tmp.Left)
		}
		if tmp.Right != nil {
			queue.PushBack(tmp.Right)
		}
	}
}
