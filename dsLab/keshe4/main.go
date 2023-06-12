package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
二叉排序树与平衡二叉树
2023年6月12日17:18:04
*/

// 二叉树节点
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// 二叉排序树
type BinarySearchTree struct {
	Root *TreeNode
}

// 平衡二叉树
type AVLTree struct {
	Root *TreeNode
}

// 向二叉排序树插入节点
func (bst *BinarySearchTree) Insert(val int) {
	node := &TreeNode{Val: val}

	if bst.Root == nil {
		bst.Root = node
		return
	}

	current := bst.Root
	for {
		if val < current.Val {
			if current.Left == nil {
				current.Left = node
				return
			}
			current = current.Left
		} else if val > current.Val {
			if current.Right == nil {
				current.Right = node
				return
			}
			current = current.Right
		} else {
			// 如果已存在相同节点值，则忽略
			return
		}
	}
}

// 中序遍历二叉树
func (bst *BinarySearchTree) InorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	bst.InorderTraversal(node.Left)
	fmt.Printf("%d ", node.Val)
	bst.InorderTraversal(node.Right)
}

// 计算二叉排序树的平均查找长度
func (bst *BinarySearchTree) CalculateAverageSearchLength() float64 {
	sumDepths, nodeCount := bst.calculateSumDepthsAndCount(bst.Root, 1)
	return float64(sumDepths) / float64(nodeCount)
}

// 递归计算二叉排序树的总深度和节点个数
func (bst *BinarySearchTree) calculateSumDepthsAndCount(node *TreeNode, depth int) (int, int) {
	if node == nil {
		return 0, 0
	}

	leftSumDepths, leftNodeCount := bst.calculateSumDepthsAndCount(node.Left, depth+1)
	rightSumDepths, rightNodeCount := bst.calculateSumDepthsAndCount(node.Right, depth+1)

	sumDepths := depth + leftSumDepths + rightSumDepths
	nodeCount := 1 + leftNodeCount + rightNodeCount

	return sumDepths, nodeCount
}

// 在二叉排序树中查找并删除节点，然后中序遍历
func (bst *BinarySearchTree) DeleteAndInorderTraversal(val int) {
	bst.Root = bst.deleteNode(bst.Root, val)
	bst.InorderTraversal(bst.Root)
}

// 在二叉排序树中递归查找并删除节点
func (bst *BinarySearchTree) deleteNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val < root.Val {
		root.Left = bst.deleteNode(root.Left, val)
	} else if val > root.Val {
		root.Right = bst.deleteNode(root.Right, val)
	} else {
		// 找到要删除的节点

		// 节点为叶子节点，直接删除
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Left == nil { // 节点只有右子树，用右子节点替代该节点
			root = root.Right
		} else if root.Right == nil { // 节点只有左子树，用左子节点替代该节点
			root = root.Left
		} else { // 节点有左右子树，找到右子树中最小的节点替代该节点
			minNode := bst.findMinNode(root.Right)
			root.Val = minNode.Val
			root.Right = bst.deleteNode(root.Right, minNode.Val)
		}
	}

	return root
}

// 在二叉排序树中找到最小的节点
func (bst *BinarySearchTree) findMinNode(node *TreeNode) *TreeNode {
	if node.Left == nil {
		return node
	}
	return bst.findMinNode(node.Left)
}

// 构建平衡二叉树
func (avl *AVLTree) BuildBalancedTree(arr []int) {
	avl.Root = avl.buildBalancedTreeHelper(arr, 0, len(arr)-1)
}

// 递归构建平衡二叉树
func (avl *AVLTree) buildBalancedTreeHelper(arr []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	node := &TreeNode{Val: arr[mid]}

	node.Left = avl.buildBalancedTreeHelper(arr, start, mid-1)
	node.Right = avl.buildBalancedTreeHelper(arr, mid+1, end)

	return node
}

// 中序遍历平衡二叉树
func (avl *AVLTree) InorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	avl.InorderTraversal(node.Left)
	fmt.Printf("%d ", node.Val)
	avl.InorderTraversal(node.Right)
}

// 计算平衡二叉树的平均查找长度
func (avl *AVLTree) CalculateAverageSearchLength() float64 {
	sumDepths, nodeCount := avl.calculateSumDepthsAndCount(avl.Root, 1)
	return float64(sumDepths) / float64(nodeCount)
}

// 递归计算平衡二叉树的总深度和节点个数
func (avl *AVLTree) calculateSumDepthsAndCount(node *TreeNode, depth int) (int, int) {
	if node == nil {
		return 0, 0
	}

	leftSumDepths, leftNodeCount := avl.calculateSumDepthsAndCount(node.Left, depth+1)
	rightSumDepths, rightNodeCount := avl.calculateSumDepthsAndCount(node.Right, depth+1)

	sumDepths := depth + leftSumDepths + rightSumDepths
	nodeCount := 1 + leftNodeCount + rightNodeCount

	return sumDepths, nodeCount
}

// 可视化输出二叉树
func visualizeTree(node *TreeNode, prefix string, isLeft bool) {
	if node == nil {
		return
	}

	fmt.Printf(prefix)
	if isLeft {
		fmt.Print("├── ")
	} else {
		fmt.Print("└── ")
	}
	fmt.Println(node.Val)

	visualizeTree(node.Left, prefix+"│   ", true)
	visualizeTree(node.Right, prefix+"    ", false)
}

func main() {
	// 读取整数序列
	file, err := os.Open("keshe4/input.txt")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	input := bufio.NewReader(file)
	var sequence []int
	str, _, _ := input.ReadLine()
	strs := strings.Split(string(str), " ")
	for _, s := range strs {
		val, _ := strconv.Atoi(s)
		sequence = append(sequence, val)
	}

	// 构建二叉排序树并输出中序遍历结果
	bst := &BinarySearchTree{}
	for _, val := range sequence {
		bst.Insert(val)
	}
	fmt.Print("二叉排序树中序遍历结果: ")
	bst.InorderTraversal(bst.Root)
	fmt.Println()

	// 计算二叉排序树的平均查找长度
	averageSearchLength := bst.CalculateAverageSearchLength()
	fmt.Printf("二叉排序树查找成功的平均查找长度: %.2f\n", averageSearchLength)

	// 输入元素并查找删除节点，然后输出中序遍历结果
	var x int
	fmt.Print("请输入要查找删除的元素: ")
	fmt.Scanln(&x)
	fmt.Print("中序遍历结果: ")
	bst.DeleteAndInorderTraversal(x)
	fmt.Println()

	// 构建平衡二叉树并输出中序遍历结果
	avl := &AVLTree{}
	avl.BuildBalancedTree(sequence)
	fmt.Print("平衡二叉树中序遍历结果: ")
	avl.InorderTraversal(avl.Root)
	fmt.Println()

	// 计算平衡二叉树的平均查找长度
	averageSearchLength = avl.CalculateAverageSearchLength()
	fmt.Printf("平衡二叉树查找成功的平均查找长度: %.2f\n", averageSearchLength)

	// 可视化输出二叉树
	fmt.Println("二叉排序树:")
	visualizeTree(bst.Root, "", true)
	fmt.Println("平衡二叉树:")
	visualizeTree(avl.Root, "", true)

	fmt.Println("程序执行完毕，请按下回车键继续...")
	fmt.Scanln() // 程序将在用户按下回车键后继续执行
}
