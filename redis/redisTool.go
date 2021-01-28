package redis


func Get(key string) (reply interface{}, err error) {
	conn := GetRedis().Get()
	defer conn.Close()
	row, err := conn.Do("get", key)
	return row, err
}

func Set(key, field string) (reply interface{}, err error) {
	conn := GetRedis().Get()
	defer conn.Close()
	row, err := conn.Do("set", key, field)
	return row, err
}
