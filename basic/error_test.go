package basic

import (
	"testing"
	"github.com/pkg/errors"
	errors2 "errors"
	"fmt"
)

func TestIs(t *testing.T) {

	Aperr := &ApError{errors.New("我是测试test")}
	err := errors.Wrap(Aperr,"增加内容")
	fmt.Println("是否err:",Is(err,Aperr))
	var a *ApError
	if errors2.As(err,&a){
		fmt.Println("err is Aperr")
	}else{
		fmt.Println("err is not Aperr")

	}

}
