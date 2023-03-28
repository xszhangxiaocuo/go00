package myStack

/*
自定义一个stack数据结构
*/

type MyStack struct {
	data     []interface{} //栈中的数据
	size     int           //栈中当前元素个数
	capacity int           //栈中最大容量
}

func NewMySack() *MyStack {
	ms := &MyStack{}
	ms.capacity = 10
	ms.data = make([]interface{}, ms.capacity)
	return ms
}

// Push 入栈操作
func (ms *MyStack) Push(d interface{}) {
	if ms.size >= ms.capacity { //扩容
		ms.capacity *= 2
		tmp := make([]interface{}, ms.capacity)
		copy(tmp[:ms.size], ms.data[:])
		ms.data = tmp
	}
	ms.data[ms.size] = d
	ms.size++
}

// Pop 弹出栈顶元素
func (ms *MyStack) Pop() interface{} {
	if ms.size == 0 {
		return nil
	}
	ms.size--
	return ms.data[ms.size]
}

// GetPop 获取栈顶元素但不弹出
func (ms *MyStack) GetPop() interface{} {
	if ms.size == 0 {
		return nil
	}
	return ms.data[ms.size-1]
}

// GetSize 获取当前栈中的元素个数
func (ms *MyStack) GetSize() int {
	return ms.size
}
