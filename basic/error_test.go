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
	//======== %w 包装  可使用Is()判断 ========
	err3 := fmt.Errorf("notfound:%w",NotFoundError)//这里的重点是 %w 可以包装
	if errors2.Is(err3,NotFoundError){
		fmt.Println("err3 is notfound")
	}else{
		fmt.Println("err3 not is notfound")

	}



	if errors2.As(err,&a){
		fmt.Println("err is Aperr")
	}else{
		fmt.Println("err is not Aperr")

	}

}
