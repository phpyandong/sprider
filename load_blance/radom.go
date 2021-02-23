package load_balance

import (
"errors"
"math/rand"
)

//随机负载均衡
type RandomBalance struct {
	curIndex int

	rss []string
}
//loadBalance.Add("126","127")

func (r *RandomBalance) Add(params ...string) error {
	if len(params) == 0 {
		return errors.New("params len 1 at least")
	}
	for i:=0; i<len(params);i++  {
		addr := params[i]
		r.rss = append(r.rss, addr)
	}


	return nil
}

func (r *RandomBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}
	r.curIndex = rand.Intn(len(r.rss))
	return r.rss[r.curIndex]
}

func (r *RandomBalance) Get(string) (string, error) {
	return r.Next(), nil
}