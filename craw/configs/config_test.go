package configs

import (
	"testing"
	"github.com/bilibili/kratos/pkg/cache/redis"
	"github.com/ghodss/yaml"
	"common"
)

func TestInit(t *testing.T) {
	Init()

}
//=========================================
//伪代码 最佳实践
//读取yaml文件，并赋值给redis config struct
func ApplyYAML(s *redis.Config,yml string) error{
	js , err := yaml.YAMLToJSON([]byte(yml))
	if err != nil {
		return err
	}
	return ApplyJSON(s,string(js))
}
//package redis
type Option interface{
	apply(*options)
}

func Options(c *redis.Config) []redis.Options{
	return []redis.Options{
		redis.DialDatabase(c.Database),
		redis.DialPassword(c.Password),
		redis.DialReadTimeout(c.ReadTimeout),
	}
}
func TestRedisConf(t *testing.T) {

	c := new (redis.Config)
	_= ApplyYAML(c,LoadConfig())
	r ,_:= redis.Dial(c.NetWork,c.Address,Options(c))
}