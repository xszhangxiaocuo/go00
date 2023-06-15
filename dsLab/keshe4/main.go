package main

import (
	"bufio"
	"fmt"
	"github.com/fogleman/gg"
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
	Height      int
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

	sumDepths := depth + leftSumDepths + rightSumDepths //当前节点所在的深度加上左子树的深度和右子树的深度
	nodeCount := 1 + leftNodeCount + rightNodeCount     //当前节点加上左子树节点个数加上右子树节点个数

	return sumDepths, nodeCount
}

// 在二叉排序树中查找并删除节点，然后中序遍历
func (bst *BinarySearchTree) DeleteAndInorderTraversal(val int) {
	delFlag = false
	bst.Root = bst.deleteNode(bst.Root, val)
	if !delFlag {
		fmt.Println("该元素不存在！")
	}
	fmt.Print("删除后的中序遍历结果: ")
	bst.InorderTraversal(bst.Root)
}

var delFlag bool

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
		delFlag = true
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

// Insert 插入节点到平衡二叉树
func (avl *AVLTree) Insert(data int) {
	// 从根节点开始插入数据
	// 根节点在动态变化，所以需要不断刷新
	avl.Root = avl.Root.Insert(data)
}

func (node *TreeNode) Insert(data int) *TreeNode {
	// 如果节点为空，则初始化该节点并返回，完成插入操作
	if node == nil {
		return &TreeNode{Val: data, Height: 1}
	}
	// 如果值重复，则什么都不做
	if node.Val == data {
		return node
	}

	// 辅助变量，用于存储旋转后子树根节点
	var newTreeNode *TreeNode

	if data > node.Val { // 插入的值大于当前节点值，要从右子树插入
		node.Right = node.Right.Insert(data)
		// 计算插入节点后当前节点的平衡因子
		// 平衡因子绝对值不能大于 1
		bf := node.BalanceFactor()
		// 如果右子树高度变高了，导致左子树-右子树的高度从 -1 变成了 -2
		if bf == -2 {
			if data > node.Right.Val { // 在右子树中插入右子节点导致失衡，需要单左旋
				newTreeNode = LeftRotate(node)
			} else { // 在右子树中插上左子节点导致失衡，需要先右旋后左旋
				newTreeNode = RightLeftRotation(node)
			}
		}
	} else {
		// 插入的值小于当前节点值，要从左子树插入
		node.Left = node.Left.Insert(data)
		bf := node.BalanceFactor()
		// 左子树的高度变高了，导致左子树-右子树的高度从 1 变成了 2
		if bf == 2 {
			if data < node.Left.Val {
				// 在左子树中插入左子节点导致失衡，需要单右旋
				newTreeNode = RightRotate(node)
			} else {
				// 在左子树中插入右子节点导致失衡，需要先左旋后右旋
				newTreeNode = LeftRightRotation(node)
			}
		}
	}

	if newTreeNode == nil {
		// 根节点没变，直接更新子树高度，并返回当前节点指针
		node.UpdateHeight()
		return node
	} else {
		// 经过旋转处理后，子树根节点变了，需要更新新的子树高度，然后返回新的子树根节点指针
		newTreeNode.UpdateHeight()
		return newTreeNode
	}
}

// UpdateHeight 更新节点树高度
func (node *TreeNode) UpdateHeight() {
	if node == nil {
		return
	}

	// 分别计算左子树和右子树的高度
	leftHeight, rightHeight := 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}

	// 以更高的子树高度作为节点树高度
	maxHeight := leftHeight
	if rightHeight > maxHeight {
		maxHeight = rightHeight
	}

	// 最终高度要加上节点本身所在的那一层
	node.Height = maxHeight + 1
}

// BalanceFactor 计算节点平衡因子（即左右子树的高度差）
func (node *TreeNode) BalanceFactor() int {
	leftHeight, rightHeight := 0, 0
	if node.Left != nil {
		leftHeight = node.Left.Height
	}
	if node.Right != nil {
		rightHeight = node.Right.Height
	}
	return leftHeight - rightHeight
}

// RightRotate 右旋操作
func RightRotate(node *TreeNode) *TreeNode {
	pivot := node.Left    // pivot 表示新插入的节点
	pivotR := pivot.Right // 暂存 pivot 右子树入口节点
	pivot.Right = node    // 右旋后最小不平衡子树根节点 node 变成 pivot 的右子节点
	node.Left = pivotR    // pivot 原本的右子节点需要挂载到 node 节点的左子树上

	// 只有 node 和 pivot 的高度改变了
	node.UpdateHeight()
	pivot.UpdateHeight()

	// 返回右旋后的子树根节点指针，即 pivot
	return pivot
}

// LeftRotate 左旋操作
func LeftRotate(node *TreeNode) *TreeNode {
	pivot := node.Right  // pivot 表示新插入的节点
	pivotL := pivot.Left // 暂存 pivot 左子树入口节点
	pivot.Left = node    // 左旋后最小不平衡子树根节点 node 变成 pivot 的左子节点
	node.Right = pivotL  // pivot 原本的左子节点需要挂载到 node 节点的右子树上

	// 只有 node 和 pivot 的高度改变了
	node.UpdateHeight()
	pivot.UpdateHeight()

	// 返回旋后的子树根节点指针，即 pivot
	return pivot
}

// LeftRightRotation 先左旋后右旋
func LeftRightRotation(node *TreeNode) *TreeNode {
	node.Left = LeftRotate(node.Left)
	return RightRotate(node)
}

// RightLeftRotation 先右旋后左旋
func RightLeftRotation(node *TreeNode) *TreeNode {
	node.Right = RightRotate(node.Right)
	return LeftRotate(node)
}

// InorderTraversal 中序遍历平衡二叉树
func (avl *AVLTree) InorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	avl.InorderTraversal(node.Left)
	fmt.Printf("%d ", node.Val)
	avl.InorderTraversal(node.Right)
}

// CalculateAverageSearchLength 计算平衡二叉树的平均查找长度
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
func drawTree(dc *gg.Context, node *TreeNode, x, y, dx int) {
	if node == nil {
		return
	}

	const (
		radius = 20 // 节点半径
		dy     = 60 // 纵向间距
	)

	// 绘制节点
	dc.DrawCircle(float64(x), float64(y), float64(radius))
	dc.SetRGB(0, 0, 0) // 设置节点颜色为黑色
	dc.Fill()

	// 绘制节点值
	dc.SetRGB(1, 1, 1) // 设置文本颜色为白色
	dc.DrawStringAnchored(fmt.Sprintf("%d", node.Val), float64(x), float64(y), 0.5, 0.5)
	dc.SetRGB(0, 0, 0) // 设置线条颜色为黑色
	// 绘制左子节点和连接线
	if node.Left != nil {
		xl := x - dx
		yl := y + dy
		dc.DrawLine(float64(x), float64(y+radius), float64(xl), float64(yl-radius))
		dc.Stroke()
		drawTree(dc, node.Left, xl, yl, dx/2)
	}

	// 绘制右子节点和连接线
	if node.Right != nil {
		xr := x + dx
		yr := y + dy
		dc.DrawLine(float64(x), float64(y+radius), float64(xr), float64(yr-radius))
		dc.Stroke()
		drawTree(dc, node.Right, xr, yr, dx/2)
	}
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
	var data []int
	str, _, _ := input.ReadLine()
	strs := strings.Split(string(str), " ")
	for _, s := range strs {
		val, _ := strconv.Atoi(s)
		data = append(data, val)
	}

	// 构建二叉排序树并输出中序遍历结果
	bst := &BinarySearchTree{}
	for _, val := range data {
		bst.Insert(val)
	}
	fmt.Print("二叉排序树中序遍历结果: ")
	bst.InorderTraversal(bst.Root)
	fmt.Println()

	// 计算二叉排序树的平均查找长度
	averageSearchLength := bst.CalculateAverageSearchLength()
	fmt.Printf("二叉排序树查找成功的平均查找长度: %.2f\n", averageSearchLength)

	//输入元素并查找删除节点，然后输出中序遍历结果
	var num int
	fmt.Print("请输入要查找删除的元素: ")
	fmt.Scanln(&num)
	bst.DeleteAndInorderTraversal(num)
	fmt.Println()

	// 构建平衡二叉树并输出中序遍历结果
	avl := &AVLTree{}
	for _, val := range data {
		avl.Insert(val)
	}
	fmt.Print("平衡二叉树中序遍历结果: ")
	avl.InorderTraversal(avl.Root)
	fmt.Println()

	// 计算平衡二叉树的平均查找长度
	averageSearchLength = avl.CalculateAverageSearchLength()
	fmt.Printf("平衡二叉树查找成功的平均查找长度: %.2f\n", averageSearchLength)

	fmt.Println("二叉排序树:")
	visualizeTree(bst.Root, "", true)
	fmt.Println("平衡二叉树:")
	visualizeTree(avl.Root, "", true)

	// 创建绘图上下文
	const (
		width  = 1024
		height = 1024
	)
	dc := gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1) // 设置背景颜色为白色
	dc.Clear()

	// 绘制二叉树
	x := width / 4
	y := 20
	drawTree(dc, bst.Root, x, y, width/8)

	// 保存绘图结果到文件
	dc.SavePNG("bst_tree.png")

	dc = gg.NewContext(width, height)
	dc.SetRGB(1, 1, 1) // 设置背景颜色为白色
	dc.Clear()

	// 绘制二叉树
	x = width / 4
	y = 20
	drawTree(dc, avl.Root, x, y, width/8)

	// 保存绘图结果到文件
	dc.SavePNG("avl_tree.png")

}
