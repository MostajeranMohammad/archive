package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
)

var pool = newPool()

type car struct {
	Color string
	Now int
}

func main() {

}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
func ping(){
	c := pool.Get()
	defer c.Close()
	pong, err := redis.String(c.Do("PING"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("PING Response = %s\n", pong)
}
func set(key string,value []byte)error{
	c := pool.Get()
	defer c.Close()
	_ , err:= c.Do("SET",key , value)
	return err
}
func get(key string)([]byte , error){
	c := pool.Get()
	defer c.Close()
	result , err := redis.Bytes(c.Do("GET",key))
	if err != nil {
		return []byte{}, err
	}
	return result, nil
}
func del(key string)(error){
	c := pool.Get()
	defer c.Close()
	_ , err := c.Do("DEL" , key)
	return err
}