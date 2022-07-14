package model

import (
	"github.com/jinzhu/gorm"
)

type Task struct {
	gorm.Model
	User      User   `gorm:"Foreignkey:Uid"`
	Uid       uint   `gorm:"not null"`
	Title     string `gorm:"index;not null"`
	Status    string `gorm:"default:'0'"`
	Content   string `gorm:"tye:longtext"`
	StartTime int64
	EndTime   int64
}

//func (Task *Task) View() uint64 {
//	//增加点击数
//	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(Task.ID)).Result()
//	count, _ := strconv.ParseUint(countStr, 10, 64)
//	return count
//}
//
////AddView
//func (Task *Task) AddView() {
//	cache.RedisClient.Incr(cache.TaskViewKey(Task.ID))                      //增加视频点击数
//	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Task.ID))) //增加排行点击数
//}
