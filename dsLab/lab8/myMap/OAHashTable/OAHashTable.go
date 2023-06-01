package OAHashTable

import "math"

/*
一个以string类型为key的简易hash表
散列函数：移位叠加法
解决冲突的方法：开放定址法
只提供了Put,Get方法
实现了自动扩容
2023年5月31日16:48:34
*/
var nowCapacity uint = 500 //当前容量

var maxCapacity uint = 50000 //最大容量
const loadFactor = 0.75      //负载因子

type Entry struct {
	key   string
	value interface{}
	count int //线性探测次数
}

type OAHashTable struct {
	Size   int     //当前map中，bucket中已经有key存在的桶的个数
	bucket []Entry //存放entry的桶
}

// CreatOAHashTable CreatMap 创建一个hash table并返回指针
func CreatOAHashTable() *OAHashTable {
	return &OAHashTable{0, make([]Entry, nowCapacity, maxCapacity)}
}

func SetOAMaxCapacity(capacity uint) {
	maxCapacity = capacity
}

// OAhashCode 计算hash code，传入的字符串为纯数字
func OAhashCode(k string, capacity uint) uint {
	var sum uint = 0
	var n uint = 0
	l := 0
	for capacity != 0 {
		capacity /= 10
		l++
	}
	for i, ch := range []byte(k) {
		if (i+1)%l == 0 {
			sum += n
			n = 0
		} else {

			n = n*10 + uint(ch-'0')
		}
	}
	sum += n

	return sum % uint(math.Pow(10, float64(l)))
}

// Put 对外暴露的插入方法，内部实际上直接调用insert
func (mm *OAHashTable) Put(k string, v interface{}) {
	mm.insert(Entry{k, v, 1})

	//扩容
	if float64(mm.Size)/float64(nowCapacity) > loadFactor {
		if nowCapacity*2 > maxCapacity { //每次扩容大小乘以2，如果结果大于maxCapacity，就将最大值赋值给nowCapacity
			nowCapacity = maxCapacity
		} else {
			nowCapacity *= 2
		}
		newTable := OAHashTable{0, make([]Entry, nowCapacity, maxCapacity)}

		for _, e := range mm.bucket { //e不是指针，是一个entry类型的结构体
			if e.value == "" {
				continue
			}
			newTable.insert(e)
		}
		*mm = newTable
	}
}

// 不对外暴露的插入方法
func (mm *OAHashTable) insert(entry Entry) {
	mm.Size++
	index := OAhashCode(entry.key, nowCapacity) //计算key所对应的哈希值，即当前要插入的下标
	e := &mm.bucket[index]                      //获取当前key所对应位置的第一个entry指针
	if e.key == "" {
		*e = entry
	}
	res := index
	for {
		if res == nowCapacity {
			res = 0
		}
		if mm.bucket[res].key == "" || mm.bucket[res].key == entry.key {
			mm.bucket[res] = entry
			break
		} else {
			entry.count++
			res++
			if res == index {
				break
			}
		}
	}
}

// Get 获取key为k的value
func (mm *OAHashTable) Get(k string) interface{} {
	index := OAhashCode(k, nowCapacity)
	res := index
	if mm.bucket[res].key == "" {
		return nil
	}
	for {
		if res == nowCapacity {
			res = 0
		}
		if mm.bucket[res].key == k {
			return mm.bucket[res]
		} else {
			res++
			if res == index {
				break
			}
		}
	}
	return nil
}

// GetKey 获取当前hash table中所有的key以及key的个数
func (mm *OAHashTable) GetKey() ([]string, int) {
	keys := make([]string, 0, mm.Size) //最多有size个entry节点
	num := 0                           //key的个数
	for _, e := range mm.bucket {
		if e.key == "" { //key为空说明当前位置没有entry节点
			continue
		}
		keys = append(keys, e.key)
		num++
	}
	return keys, num
}
