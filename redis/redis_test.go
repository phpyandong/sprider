package redis

import (
	"testing"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func TestInitRedis(t *testing.T) {
	InitRedis()
	conn := GetRedis().Get()
	defer conn.Close()
	profile := map[string]string{
		"name":"seiya",
	}
	result, err := conn.Do("HSET", "seiya", "name",profile["name"])
	fmt.Println(result)
	if result ==1 && err != nil {
		panic(err)
	}

	res,err := redis.String(conn.Do("HGET","seiya","name"))
	if res != profile["name"] {
		t.Errorf("expect %s but %s",profile["name"],res)

	}

	Set("nn","bb")
	fmt.Println(Get("nn"))
}