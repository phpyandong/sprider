package rpcdemo

import "github.com/kataras/iris/core/errors"

type DemoService struct {

}
type Args struct {
	A,B int
}
func (DemoService) Div(args Args,res *float64) error{
	if args.B == 0 {
		return errors.New("devision by zero	")
	}
	*res = float64(args.B)/  float64(args.A)
	return nil
}
