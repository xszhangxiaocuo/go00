package myMap

/*
一个以string类型为key的简易map
只提供了Put,Get方法
实现了自动扩容
2023年3月25日12:56:35
*/
var nowCapacity = 20 //当前容量

var maxCapacity = 200   //最大容量
const loadFactor = 0.75 //负载因子

type Entry struct {
	key   string
	value interface{}
	next  *Entry //数组相同位置，下一个Entry的指针
}

type MyMap struct {
	Size   int     //当前map中，key的个数
	bucket []Entry //存放entry的桶
}

// CreatMap 创建一个map并返回指针
func CreatMap() *MyMap {
	return &MyMap{0, make([]Entry, nowCapacity, maxCapacity)}
}

func SetMaxCapacity(capacity int) {
	maxCapacity = capacity
}

// 计算hash code，将字符串k转化为byte数组并将所有值相加对length取余
func hashCode(k string, length int) int {
	sum := 0
	for _, num := range []byte(k) {
		sum += int(num)
	}
	return sum % length
}

// Put 对外暴露的插入方法，内部实际上直接调用insert
func (mm *MyMap) Put(k string, v interface{}) {
	mm.insert(Entry{k, v, nil})

	//扩容
	if float64(mm.Size)/float64(nowCapacity) > loadFactor {
		if nowCapacity*2 > maxCapacity { //每次扩容大小乘以2，如果结果大于maxCapacity，就将最大值赋值给nowCapacity
			nowCapacity = maxCapacity
		} else {
			nowCapacity *= 2
		}
		newMap := MyMap{0, make([]Entry, nowCapacity, maxCapacity)}

		for _, e := range mm.bucket { //e不是指针，是一个entry类型的结构体
			if e.value == "" {
				continue
			}
			for e.next != nil {
				newMap.insert(e)
				e = *e.next //直接复制e.next节点的结构体内容，而不是复制指针
			}
			newMap.insert(e)
		}
		*mm = newMap
	}
}

// 不对外暴露的插入方法
func (mm *MyMap) insert(entry Entry) {
	mm.Size += 1
	index := hashCode(entry.key, nowCapacity) //计算key所对应的哈希值，即当前要插入的下标
	e := &mm.bucket[index]                    //获取当前key所对应位置的第一个entry指针
	if e.key == "" {
		*e = entry
	}
	for e.next != nil {
		if e.key == entry.key {
			entry.next = e.next
			*e = entry
			return
		}
		e = e.next
	}
	if e.key == entry.key {
		entry.next = e.next
		*e = entry
		return
	}
	e.next = &entry
}

// Get 获取key为k的value
func (mm *MyMap) Get(k string) interface{} {
	index := hashCode(k, nowCapacity)
	e := &mm.bucket[index]
	for e.next != nil {
		if e.key == k {
			return e.value
		}
		e = e.next
	}
	if e.key == k {
		return e.value
	}
	return nil
}

// GetKey 获取当前map中所有的key
func (mm *MyMap) GetKey() []string {
	keys := make([]string, 0, mm.Size) //至少有size个entry节点
	for _, e := range mm.bucket {
		tmpe := &e
		if tmpe.key == "" { //key为空说明当前位置没有entry节点
			continue
		}
		for tmpe.next != nil {
			keys = append(keys, tmpe.key)
			tmpe = tmpe.next
		}
		keys = append(keys, tmpe.key)
	}
	return keys
}
