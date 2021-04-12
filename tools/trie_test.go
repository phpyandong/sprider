package tools

import (
	"testing"
	"fmt"
)

func TestTrie(t *testing.T) {
	tree := NewTrie()

	tree.Add("1234","我干","我日","我草")
	//second := tree.Root.Children
	//fmt.Println(second)
	//third := second[49].Children
	//fmt.Println(third)
	//three := third[50].Children
	//fmt.Println(three)
	//fmt.Println(tree.Root.Character)
	//fmt.Println(tree.Root.Children[49].Character)
	//fmt.Println(tree.Root.Children[49].Children[50].Character)
	//fmt.Println(tree.Root.Children[49].Children[50].Children[51].Character)


	//fmt.Println(tree.Validate("c12d111235"))
	fmt.Println(tree.Validate2("c12d1112"))

	//fmt.Println(tree.Replace("123345", '*'))
	//
	fmt.Println(tree.Filter2("1212341"))
	//fmt.Println(tree.Filter("你好吗 我支持学习题目， 他的名字叫做习近平"))


}