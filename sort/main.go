package main

import (
	"sort"
	"math/rand"
	"fmt"
)
type Hero struct{
	name string
	Age int
}
type Herosclice []Hero
func (heros Herosclice ) Len() int{
	return len(heros)
}
// Less reports whether the element with
// index i should sort before the element with index j.
func (heros Herosclice ) Less(i, j int) bool{
	if heros[i].Age > heros[j].Age {
		return true
	}
	return false
}
// Swap swaps the elements with indexes i and j.
func (heros Herosclice)Swap(i, j int){
	heros[i],heros[j] = heros[j],heros[i]
}
func main()  {
	var heros Herosclice
	for i:=1;i<108 ;i++  {
		hero := Hero{fmt.Sprintf("水浒%d将",i),rand.Intn(60)}
		heros = append(heros,hero)
	}
	fmt.Println(heros)

	sort.Sort(heros)
	fmt.Println(heros)
}
