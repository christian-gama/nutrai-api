package conn

func MakeRedis() {
	if redisInstance != nil {
		return
	}

	redisInstance = NewConn(0)
}
