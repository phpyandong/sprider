package subject

import (
	"math/rand"
	"fmt"
)
//todo 子集算法背景 由于实现的是客户端心跳检测，过多的客户端检测服务端，服务端压力很大，因此需要从全部的节点取一部分，
// from google sre
//为什么上面这个算法可以保证可以均匀分布？
//首先，shuffle 算法保证在 round 一致的情况下，backend 的排列一定是一致的。
//因为每个实例拥有从 0 开始的连续唯一的自增 id，且计算过程能够保证每个 round 内所有实例拿到的服务列表的排列一致，
// 因此在同一个 round 内的 client 会分别 backend 排列的不同部分的切片作为选中的后端服务来建立连接。
//所以只要 client id 是连续的，那么 client 发向 后端的连接就一定是连续的
//参考资料:
/**
	backends 后端节点
	clientId 消费者Id
	subsetSize 子集大小
 */
func Subset(backends []string, clientID, subsetSize int) []string {
	subsetCount := len(backends) / subsetSize //几批子集
	fmt.Println("subsetcount:11/5",subsetCount)
	// Group clients into rounds; each round uses the same shuffled list:
	//对客户端 进行分组;每一轮使用相同的洗牌列表:
	round := clientID / subsetCount//种子
	fmt.Println("round:3/2=",round)

	r := rand.New(rand.NewSource(int64(round)))
	//洗牌算法
	r.Shuffle(len(backends),
		func(i, j int) {
			backends[i], backends[j] = backends[j], backends[i]
		},
		)

	// The subset id corresponding to the current client:
	// 当前客户端对应的子集id:
	subsetID := clientID % subsetCount //3 %种子
	fmt.Printf("clentId:%v,subsetCount:%v:SubsetID:%v\n",clientID,subsetCount,subsetID)
	start := subsetID * subsetSize
	fmt.Printf("subsetId:%v;subsetSize:%v start :%v\n",subsetID,subsetSize,start)
	fmt.Printf("backends:%v\n",backends)
	return backends[start : start+subsetSize]
}