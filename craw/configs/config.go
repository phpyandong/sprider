package configs

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"gopkg.in/yaml.v2"

	"fmt"
)

func Init()  {

	c := config.New(
		config.WithSource(
			file.NewSource("../config/config.yaml"),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			// kv.Key
			// kv.Value
			// kv.Metadata
			// 自定义实现对应的数据源解析，如果是配置中心数据源也可以指定metadata进行识别配置类型
			return yaml.Unmarshal(kv.Value, v)
		}),
	)
	// 加载配置源：
	if err := c.Load(); err != nil {
		panic(err)
	}
	//// 获取对应的值内容：
	name, _ := c.Value("server.http").String()
	fmt.Println(name)
	// 解析到结构体（由于已经合并到map[string]interface{}，所以需要指定 jsonName 进行解析）：
	//var v struct {
	//	Service string `json:"service"`
	//	Version string `json:"version"`
	//}
	var v Bootstrap

	if err := c.Scan(&v); err != nil {
		panic(err)
	}
	fmt.Println(v.Data.Redis.Addr)
	//// 监听值内容变更
	//c.Watch("service.name", func(key string, value config.Value) {
	//	// 值内容变更
	//})

}