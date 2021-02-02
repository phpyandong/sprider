package subject

import (
	"testing"
	"fmt"
)

func TestSubset(t *testing.T) {
	backends := []string{"123","124","125","126","127","128","129","130","131","132","133"}
	clientId := 3
	subsetSize := 5
	list := Subset(backends,clientId,subsetSize)
	fmt.Println(list)//[124 132 123 128 125]
}