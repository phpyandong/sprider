package main

import (
	"sprider/load_blance"
	"fmt"
)

func main() {
	loadBalance := load_balance.LoadBalanceFactory(load_balance.LbConsistentHash)
	//loadBalance.Add("126 20","127 30")
	//loadBalance.Add("126 20","127 30")
	loadBalance.Add("126 20","127 30")

	for i:=0;i<10 ;i++  {
		ip,_:= loadBalance.Get("")
		fmt.Printf("loadBalance :%v ip:%+v \n",loadBalance,ip)
	}


}
