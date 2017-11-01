package main

import (
	"encoding/json"
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//testRedis()
	//test1()
	// var str string
	// if len(str) == 0 {
	// 	fmt.Println("failed")
	// }
	testPushRedis()
}

func test1() {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	key := "profile"
	imap := map[string]string{"username": "666", "phonenumber": "888"}
	value, _ := json.Marshal(imap)

	n, err := c.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
	}
	if n == int64(1) {
		fmt.Println("success")
	}

	var imapGet map[string]string

	valueGet, err := redis.Bytes(c.Do("GET", key))
	if err != nil {
		fmt.Println(err)
	}

	errShal := json.Unmarshal(valueGet, &imapGet)
	if errShal != nil {
		fmt.Println(err)
	}
	fmt.Println(imapGet["username"])
	fmt.Println(imapGet["phonenumber"])
}

func testRedis() {
	r, err := connectRedis("127.0.0.1:6379")
	defer r.Close()
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	key := "helloking"
	fmt.Println("redis set ", setRedis(key, "helloking", 12, r))
	fmt.Println("redis get ", getRedis(key, r))
	fmt.Println("exists redis ", existsRedis(key, r))
	fmt.Println("delete redis ", deleteRedis(key, r))
	fmt.Println("exists redis ", existsRedis(key, r))
	jsonRedis(r)
}

func connectRedis(conn string) (redis.Conn, error) {
	r, err := redis.Dial("tcp", conn)
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return nil, err
	}
	return r, nil
}

func testPushRedis() {
	r, err := connectRedis("127.0.0.1:6379")
	defer r.Close()
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	key := "helloking"
	pushRedis(key, "java", r)
	pushRedis(key, "C#", r)
	pushRedis(key, "javascript", r)
	pushRedis(key, "python", r)
	pushRedis(key, "golang", r)
	values, _ := redis.Values(r.Do("lrange", key, "0", "100"))
	for _, value := range values {
		fmt.Println(string(value.([]byte)))
	}
}

func pushRedis(key, value string, conn redis.Conn) bool {
	if len(key) == 0 || len(value) == 0 || conn == nil {
		return false
	}
	_, err := conn.Do("lpush", key, value)
	if err != nil {
		return false
	}
	return true
}

func setRedis(key, value string, time int, conn redis.Conn) bool {
	if len(key) == 0 || len(value) == 0 {
		return false
	}
	if conn == nil {
		return false
	}
	_, err := conn.Do("SET", key, value, "EX", time)
	if err != nil {
		return false
	}
	return true
}

func getRedis(key string, conn redis.Conn) string {
	if len(key) == 0 || conn == nil {
		return ""
	}
	info, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return ""
	}
	return info
}

func existsRedis(key string, conn redis.Conn) bool {
	if len(key) == 0 || conn == nil {
		return false
	}
	flag, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return flag
}

func deleteRedis(key string, conn redis.Conn) bool {
	if len(key) == 0 || conn == nil {
		return false
	}
	if !existsRedis(key, conn) {
		return true
	}
	flag, err := redis.Bool(conn.Do("DEL", key))
	if err != nil {
		return false
	}
	return flag
}

func setTimeRedis(key string, time int, conn redis.Conn) bool {
	if len(key) == 0 || time <= 0 || conn == nil {
		return false
	}
	n, err := conn.Do("EXPIRE", time, time)
	if err != nil {
		return false
	}
	if n == int64(1) {
		return true
	}
	return false
}

func jsonRedis(conn redis.Conn) {
	if conn == nil {
		return
	}
	key := "king1"
	imap := map[string]string{"userName": "king", "passWord": "123456"}
	value, _ := json.Marshal(imap)
	fmt.Println(value)
	n, err := conn.Do("SETNX", key, value)
	if err != nil {
		fmt.Println(err)
		return
	}
	if n == int64(1) {
		fmt.Println("success")
	}
	var impGet map[string]string
	valueGet, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		return
	}
	err = json.Unmarshal(valueGet, &impGet)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(impGet)
}
