package bloom

import (
"hash"
"hash/fnv"
"fmt"
)

// The standard bloom filter, which allows adding of
// elements, and checking for their existence
//标准的bloom过滤器，它允许添加元素，并检查元素是否存在
type BloomFilter struct {
	bitmap []bool      // 布隆过滤器的位图
	hashNums      int         //k 哈希函数的数目
	itemSize      int         //n 过滤器中元素的数目
	bloomSize      int         //m bloom过滤器的容量
	hashFunc hash.Hash64 // 定义 哈希函数
}

// Returns a new BloomFilter object, if you pass the
// number of Hash Functions to use and the maximum
// size of the Bloom Filter
//如果您传递要使用的散列函数的数量和Bloom过滤器的最大大小，则返回一个新的BloomFilter对象
func NewBloomFilter(numHashFuncs, bfSize int) *BloomFilter {
	boolFilter := new(BloomFilter)

	boolFilter.bitmap = make([]bool, bfSize)

	boolFilter.hashNums =  numHashFuncs

	boolFilter.bloomSize = bfSize

	boolFilter.itemSize = 0  //元素的数量

	boolFilter.hashFunc = fnv.New64()//具体的hash函数
	return boolFilter
}

func (bf *BloomFilter) getHash(b []byte) (uint32, uint32) {
	bf.hashFunc.Reset()
	bf.hashFunc.Write(b)
	hash64No := bf.hashFunc.Sum64()//64位的二进制
	fmt.Printf("hash64    二:%b 十：%d\n",hash64No,hash64No)
	fmt.Printf("hash64>>32二:%b 十：%d\n",hash64No >> 32,hash64No >> 32)

	fmt.Printf("1<<32 十:%d 二（32）：%b 二：%b\n",1 << 32,1 << 32,(1 << 32)-1)//unsigend long int是无符号整数类型， 能表示的整数范围是0~4294967295，
	h1 := uint32( hash64No & ( (1 << 32) - 1)) //32位的1
	fmt.Printf("32位的1& hash64:%b\n",h1)

	h2 := uint32(hash64No >> 32)  //除以 2^32;直接表现为 将64位的前32位保留，后32位去掉
	fmt.Printf("hash64 >> 32:%b\n",h2)
	fmt.Printf("h1:h2:%b %b \n",h1,h2)
	//h1 是64位（实际后32位，因为与运算如果空位为0，32个1 只有32位，前32位都是0）与运算 32个1
	// ；h2 是前32位
	return h1, h2

}

// Adds an element (in byte-array form) to the Bloom Filter
//添加一个元素(以字节数组的形式)到Bloom过滤器
func (bf *BloomFilter) Add(e []byte) {
	h1, h2 := bf.getHash(e)
	//for i := 0; i < bf.hashNums; i++ {
	//								% 布隆过滤器的容量
	ind := (h1 + h2) % uint32(bf.bloomSize)
	bf.bitmap[ind] = true
	//}
	bf.itemSize++
}

// Checks if an element (in byte-array form) exists in the
// Bloom Filter
//检查是否有一个元素(以字节数组形式)存在于Bloom过滤器中
func (bf *BloomFilter) Check(x []byte) bool {
	h1, h2 := bf.getHash(x)
	result := true
	for i := 0; i < bf.hashNums; i++ {
		ind := (h1 + uint32(i)*h2) % uint32(bf.bloomSize)
		result = result && bf.bitmap[ind]
	}
	return result
}
