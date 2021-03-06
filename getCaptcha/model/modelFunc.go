package model

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

// 存储图片id 到redis 数据库
func SaveImgCode(code, uuid string) error {
	log.Print(code, uuid)
	// 1. 链接数据库
	conn, err := redis.Dial("tcp", "192.168.6.108:6379")
	if err != nil {
		fmt.Println("redis Dial err:", err)
		return err
	}
	defer conn.Close()
	// 2. 写数据库  --- 有效时间 5 分钟
	_, err = conn.Do("setex", uuid, 60*5, code)

	return nil  // 不需要回复助手!
}