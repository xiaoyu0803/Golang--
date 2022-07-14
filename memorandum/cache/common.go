package cache

//
//import (
//	"github.com/go-redis/redis"
//	"gopkg.in/ini.v1"
//	"log"
//	"strconv"
//)
//
//var (
//	RedisClient *redis.Client
//	RedisAddr   string
//	RedisPw     string
//	RedisDbName string
//)
//
//func init() {
//	file, err := ini.Load("conf/config.ini")
//	if err != nil {
//		log.Println("配置文件错误", err)
//	}
//	loadRedis(file)
//	Redis()
//}
//func Redis() {
//	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
//	client := redis.NewClient(&redis.Options{
//		Addr:     RedisAddr,
//		Password: RedisPw,
//		DB:       int(db),
//	})
//	_, err := client.Ping().Result()
//	if err != nil {
//		log.Println(err)
//		panic(err)
//	}
//	RedisClient = client
//}
//func loadRedis(file *ini.File) {
//	RedisAddr = file.Section("redis").Key("RedisAddr").String()
//	RedisPw = file.Section("redis").Key("RedisPw").String()
//	RedisDbName = file.Section("redis").Key("RedisDbName").String()
//}
