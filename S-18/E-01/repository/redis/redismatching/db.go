package redismatching

import "E-01/adapter/redis"

type DB struct {
	adapter redis.Adapter
}

func New (adapter redis.Adapter)DB {
	return DB{
		adapter: adapter,
	}
}



