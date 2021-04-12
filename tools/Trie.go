package tools

// Trie 短语组成的Trie树.
type Trie struct {
	Root *Node
}

// Node Trie树上的一个节点.
type Node struct {
	isRootNode bool
	isPathEnd  bool
	Character  rune
	Children   map[rune]*Node
}

// NewTrie 新建一棵Trie
func NewTrie() *Trie {
	return &Trie{
		Root: NewRootNode(0),
	}
}

// Add 添加若干个词
func (tree *Trie) Add(words ...string) {
	for _, word := range words {
		tree.add(word)
	}
}

//func (tree *Trie) add(word string) {
//	var current = tree.Root
//	var runes = []rune(word)
//	for position := 0; position < len(runes); position++ {
//		r := runes[position]
//		if next, ok := current.Children[r]; ok {
//			current = next
//		} else {
//			newNode := NewNode(r)
//			current.Children[r] = newNode
//			current = newNode
//		}
//		if position == len(runes)-1 {
//			current.isPathEnd = true
//		}
//	}
//}
func (tree *Trie) add(word string){
	var current = tree.Root
	var wordsArr = []rune(word)
	for idx := 0; idx <= len(wordsArr)-1; idx++  {
		currWord := wordsArr[idx]
		if childNode,ok := current.Children[currWord] ;ok {
			current = childNode
		}else{
			childNode := NewNode(currWord)
			current.Children[currWord] = childNode
			current = childNode
		}
		if idx == len(wordsArr) -1 {
			current.isPathEnd = true
		}
	}
}

func (tree *Trie) Del(words ...string) {
	for _, word := range words {
		tree.del(word)
	}
}

func (tree *Trie) del(word string) {
	var current = tree.Root
	var runes = []rune(word)
	for position := 0; position < len(runes); position++ {
		r := runes[position]
		if next, ok := current.Children[r]; !ok {
			return
		} else {
			current = next
		}

		if position == len(runes)-1 {
			current.SoftDel()
		}
	}
}

// Replace 词语替换
func (tree *Trie) Replace(text string, character rune) string {
	var (
		parent  = tree.Root
		current *Node
		runes   = []rune(text)
		length  = len(runes)
		left    = 0
		found   bool
	)

	for position := 0; position < len(runes); position++ {
		current, found = parent.Children[runes[position]]

		if !found || (!current.IsPathEnd() && position == length-1) {
			parent = tree.Root
			position = left
			left++
			continue
		}

		// println(string(current.Character), current.IsPathEnd(), left)
		if current.IsPathEnd() && left <= position {
			for i := left; i <= position; i++ {
				runes[i] = character
			}
		}

		parent = current
	}

	return string(runes)
}

// Filter 直接过滤掉字符串中的敏感词
func (tree *Trie) Filter(text string) string {
	var (
		parent      = tree.Root
		current     *Node
		left        = 0
		found       bool
		runes       = []rune(text)
		length      = len(runes)
		resultRunes = make([]rune, 0, length)
	)

	for position := 0; position < length; position++ {
		current, found = parent.Children[runes[position]]

		if !found || (!current.IsPathEnd() && position == length-1) {
			resultRunes = append(resultRunes, runes[left])
			parent = tree.Root
			position = left   //left的思想总是排除掉上一个位置
			left++
			continue
		}

		if current.IsPathEnd() {//当前节点为最后一个
			left = position+1  //找到了；；；【left 的赋值逻辑根据if else 的逻辑的不同而变更】
			parent = tree.Root //说明已经到头了，要从根节点重新查找，这是一个无限循环的过程，只要文字还没有读取完毕
		} else {
			parent = current
		}

	}

	resultRunes = append(resultRunes, runes[left:]...)
	return string(resultRunes)
}
func (tree *Trie) Filter2(words string) string{
	worsArr := []rune(words)
	var (
		parent = tree.Root
		current *Node
		left = 0
		length = len(words)
		result = make([]rune,0,length)
		found bool
	)
	for idx := 0 ; idx < length ; idx++ {
		currentWord := worsArr[idx]
		current, found = parent.Children[currentWord]
		if !found {

			result = append(result,worsArr[left])
			idx = left
			parent = tree.Root
			left ++
			continue
		}

		if current.IsPathEnd() {
			left = idx+1
			//循环要继续
			parent = tree.Root
		} else {
			parent = current

		}
	}
		result = append(result,worsArr[left:]...)
		return string(result)

}

func (tree *Trie) Validate2(text string) (bool, string){
	wordArr := []rune(text)
	var parent = tree.Root
	var currNode *Node
	var left = 0
	var found bool
	for idx := 0; idx< len(wordArr) ; idx ++ {
		currWord := wordArr[idx]
		currNode,found = parent.Children[currWord]
		if !found ||(!currNode.IsPathEnd() && idx == len(wordArr) -1 ) {//还是没有找出来原因
			parent = tree.Root //从根节点开始
			idx = left
			left++

			continue
		}

		//判断是否是根节点，返回数据
		if currNode.IsPathEnd() && left <=idx {
			return false,string(wordArr[left:idx+1])
		}
		parent = currNode


	}

	return true,""
}
// Validate 验证字符串是否合法，如不合法则返回false和检测到
// 的第一个敏感词
func (tree *Trie) Validate(text string) (bool, string) {
	//const (
	//	Empty = ""
	//)
	var (
		parent  = tree.Root
		current *Node
		runes   = []rune(text)
		length  = len(runes)
		left    = 0  //用来记录敏感词的开始位置
		found   bool
	)

	for position := 0; position < len(runes); position++ {
		current, found = parent.Children[runes[position]]
		//如果找不到数据，或者找到了但是当前词组位置已经移动到最后了，并且节点不是跟节点
		//比如 敏感词是123 ，句子中是12  则不能匹配
		if !found || (!current.IsPathEnd() && position == length-1) {
			left++
			runes = runes[left:length]//length
			return tree.Validate(string(runes))//递归比较好处理和理解,但也是建立在总是从上一个发现的数据开始处理的方案之上的
			parent = tree.Root
			position = left  //解决的场景是 b1d123  b11123有问题 ，需要重置position，否则会将d123截取出来，即第一个1 是误导的数据，
			left++
			continue
		}
		//0 0 0   //0  0
				  //0  1
		//1 1 0
		//2 2 1
		//2 3 2
		//2 4 3
		//两种情况
		//a : 123 敏感词 和 123 内容 一一对应，数据，从循环三次从开头到结尾拿到数据，敏感词的位置则是从0 to end比较好理解
		//b : 123 敏感词 和 cd123 ，需要跳过 cd 并把left 置为 +=2，开始循环3次 敏感词的位置则是从2 to end

		//截取敏感词的开始位置 到当前位置//判断条件position > left 因此初步理解是要把postion减少，而不是和当前left相等
		if current.IsPathEnd() && left <= position {
			return false, string(runes[left : position+1])
		}

		parent = current
	}

	return true, ""
}

// FindIn 判断text中是否含有词库中的词
func (tree *Trie) FindIn(text string) (bool, string) {
	validated, first := tree.Validate(text)
	return !validated, first
}

// FindAll 找有所有包含在词库中的词
func (tree *Trie) FindAll(text string) []string {
	var matches []string
	var (
		parent  = tree.Root
		current *Node
		runes   = []rune(text)
		length  = len(runes)
		left    = 0
		found   bool
	)

	for position := 0; position < length; position++ {
		current, found = parent.Children[runes[position]]

		if !found {
			parent = tree.Root
			position = left
			left++
			continue
		}

		if current.IsPathEnd() && left <= position {
			matches = append(matches, string(runes[left:position+1]))
		}

		if position == length-1 {
			parent = tree.Root
			position = left
			left++
			continue
		}

		parent = current
	}

	var i = 0
	if count := len(matches); count > 0 {
		set := make(map[string]struct{})
		for i < count {
			_, ok := set[matches[i]]
			if !ok {
				set[matches[i]] = struct{}{}
				i++
				continue
			}
			count--
			copy(matches[i:], matches[i+1:])
		}
		return matches[:count]
	}

	return nil
}

// NewNode 新建子节点
func NewNode(character rune) *Node {
	return &Node{
		Character: character,
		Children:  make(map[rune]*Node, 0),
	}
}

// NewRootNode 新建根节点
func NewRootNode(character rune) *Node {
	return &Node{
		isRootNode: true,
		Character:  character,
		Children:   make(map[rune]*Node, 0),
	}
}

// IsLeafNode 判断是否叶子节点
func (node *Node) IsLeafNode() bool {
	return len(node.Children) == 0
}

// IsRootNode 判断是否为根节点
func (node *Node) IsRootNode() bool {
	return node.isRootNode
}

// IsPathEnd 判断是否为某个路径的结束
func (node *Node) IsPathEnd() bool {
	return node.isPathEnd
}

// SoftDel 置软删除状态
func (node *Node) SoftDel() {
	node.isPathEnd = false
}
