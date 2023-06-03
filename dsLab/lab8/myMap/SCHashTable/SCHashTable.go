package SCHashTable

import "math"

/*
一个以string类型为key的简易哈希表
散列函数：除留余数法
解决冲突的方法：拉链法
只提供了Put,Get方法
实现了自动扩容
2023年5月31日16:48:27
*/
var nowCapacity uint = 2000 //当前容量

var maxCapacity uint = 50000 //最大容量
const loadFactor = 0.2       //负载因子

type Entry struct {
	key   string
	value interface{}
	count int    //探测次数
	next  *Entry //数组相同位置，下一个Entry的指针
}

type SCHashTable struct {
	Size     int     //当前map中，bucket中已经有key存在的桶的个数
	bucket   []Entry //存放entry的桶
	Count    int     //当前表中探测次数总数
	Conflict int     //当前哈希表中冲突的元素个数（即探测次数大于1的元素个数）
}

// CreatSCHashTable 创建一个hash table并返回指针
func CreatSCHashTable() *SCHashTable {
	return &SCHashTable{0, make([]Entry, nowCapacity, maxCapacity), 0, 0}
}

func SetMaxCapacity(capacity uint) {
	maxCapacity = capacity
}

// SChashCode 计算hash code，将字符串k转化为byte数组并将所有值相加对length取余
//func SChashCode(k string, length uint) uint {
//	var sum uint = 0
//	for _, num := range []byte(k) {
//		sum += uint(num)
//	}
//	return sum % length
//}

func SChashCode(k string, capacity uint) uint {
	var sum uint = 0
	var n uint = 0
	l := 0
	for capacity != 0 {
		capacity /= 10
		l++
	}
	for i, ch := range []byte(k) {
		if ch >= 'A' && ch <= 'Z' {
			n = n*10 + uint(ch-'A')%10
		} else if ch >= 'a' && ch <= 'z' {
			n = n*10 + uint(ch-'a')%10
		}

		if (i+1)%l == 0 {
			sum += n
			n = 0
		}
	}
	sum += n
	hashcode := sum % uint(math.Pow(10, float64(l)))
	if hashcode >= capacity {
		hashcode = sum % uint(math.Pow(10, float64(l-1)))
	}
	return hashcode
}

// Put 对外暴露的插入方法，内部实际上直接调用insert
func (mm *SCHashTable) Put(k string, v interface{}) {
	mm.insert(Entry{k, v, 1, nil})

	//扩容
	if float64(mm.Size)/float64(nowCapacity) > loadFactor {
		if nowCapacity*2 > maxCapacity { //每次扩容大小乘以2，如果结果大于maxCapacity，就将最大值赋值给nowCapacity
			nowCapacity = maxCapacity
		} else {
			nowCapacity *= 2
		}
		newSC := SCHashTable{0, make([]Entry, nowCapacity, maxCapacity), 0, 0}

		for _, e := range mm.bucket { //e不是指针，是一个entry类型的结构体
			if e.value == nil {
				continue
			}
			for e.next != nil {
				newSC.insert(e)
				e = *e.next //直接复制e.next节点的结构体内容，而不是复制指针
			}
			newSC.insert(e)
		}
		*mm = newSC
	}
}

// 不对外暴露的插入方法
func (mm *SCHashTable) insert(entry Entry) {
	mm.Size++
	index := SChashCode(entry.key, nowCapacity) //计算key所对应的哈希值，即当前要插入的下标
	e := &mm.bucket[index]                      //获取当前key所对应位置的第一个entry指针
	if e.key == "" {
		*e = entry
		mm.Count += entry.count
		return
	}
	mm.Conflict++
	for e.next != nil {
		if e.key == entry.key {
			entry.next = e.next
			*e = entry
			mm.Count += entry.count
			return
		}
		e = e.next
		entry.count++
	}
	if e.key == entry.key {
		entry.next = e.next
		*e = entry
		mm.Count += entry.count
		return
	}
	entry.count++
	e.next = &entry
	mm.Count += entry.count
}

// Get 获取key为k的value
func (mm *SCHashTable) Get(k string) interface{} {
	index := SChashCode(k, nowCapacity)
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

// GetCount 获取key为k的节点的count,0表示节点不存在
func (mm *SCHashTable) GetCount(k string) int {
	index := SChashCode(k, nowCapacity)
	e := &mm.bucket[index]
	for e.next != nil {
		if e.key == k {
			return e.count
		}
		e = e.next
	}
	if e.key == k {
		return e.count
	}
	return 0
}

// GetKey 获取当前hash table中所有的key以及key的个数
func (mm *SCHashTable) GetKey() ([]string, int) {
	keys := make([]string, 0, mm.Size) //最多有size个entry节点
	num := 0                           //key的个数
	for _, e := range mm.bucket {
		tmpe := &e
		if tmpe.key == "" { //key为空说明当前位置没有entry节点
			continue
		}
		for tmpe.next != nil {
			keys = append(keys, tmpe.key)
			tmpe = tmpe.next
			num++
		}
		keys = append(keys, tmpe.key)
		num++
	}
	return keys, num
}
