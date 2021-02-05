package basic

import (
	"testing"
	"github.com/pkg/errors"
	"fmt"
)

func TestIs(t *testing.T) {

	Aperr := ApError{errors.New("我是测试test")}
	err := errors.Wrap(Aperr.Err,"增加内容")
	fmt.Println("是否err:",Is(err,Aperr.Err))
}
