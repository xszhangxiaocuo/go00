package Queue

/*
自定义链队列,存储string类型数据
2023年4月7日10:37:51
*/

type Queue struct {
	head *Node
	tail *Node
	size uint
}

type Node struct {
	Value  string  //进程名
	Arrive uint    //到达时间
	Run    uint    //运行时间
	Start  uint    //到达时间
	Finish uint    //完成时间
	Ti     uint    //周转时间
	Wi     float64 //带权周转时间
	T      uint    //该任务被执行的次数
	next   *Node
}

// NewQueue 返回队列的头节点，头节点不存储数据
func NewQueue() *Queue {
	tmp := &Node{}
	return &Queue{head: tmp, tail: tmp, size: 0}
}

// Push 在队尾添加元素
func (q *Queue) Push(v string, a uint, r uint) {
	tmp := &Node{Value: v, Arrive: a, Run: r, next: nil}
	q.tail.next = tmp
	q.tail = tmp
	q.size++
}

func (q *Queue) PushNode(node *Node) {
	node.next = nil
	q.tail.next = node
	q.tail = node
	q.size++
}

// Pop 弹出队首元素并返回
func (q *Queue) Pop() *Node {
	if q.size == 0 {
		return nil
	}
	tmp := q.head.next
	if q.size == 1 {
		q.tail = q.head
	} else {
		q.head.next = q.head.next.next
	}
	q.size--
	return tmp
}

// GetHead 查看队首元素
func (q *Queue) GetHead() *Node {
	return q.head.next
}

// IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}
