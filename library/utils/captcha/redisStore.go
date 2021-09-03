package captcha

import (
	"fmt"
	"goEasy/library/utils/cache"
	"log"
	"time"
)

const CAPTCHA = "captcha:"

type RedisStore struct {
}

//set a capt
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	//120秒
	//_, err := g.Redis().Do("SET", key, value, "EX", "120")
	err := cache.GCacheRedis.Set(key, value, time.Minute*2)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return err
}

//get a capt
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	//val, err := g.Redis().DoVar("GET", key)
	val, err := cache.GCacheRedis.GetVar(key)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		//_, err := g.Redis().Do("DEL", key)
		_, err := cache.GCacheRedis.Remove(key)
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val.String()
}

//verify a capt
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	//fmt.Println("key:"+id+";value:"+v+";answer:"+answer)
	return v == answer
}
