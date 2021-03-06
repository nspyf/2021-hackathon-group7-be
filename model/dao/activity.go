package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strconv"
	"time"
)

const (
	ActivityCacheTime = 10 * time.Minute
)

type Activity struct {
	gorm.Model
	UserId    uint
	Title     string `gorm:"varchar(128)"`
	Content   string `gorm:"text"`
	StartTime string `gorm:"varchar(64)"`
	EndTime   string `gorm:"varchar(64)"`
	Place     string `gorm:"varchar(128)"`
	Digest    string `gorm:"varchar(255)"`
}

//mysql
//添加新活动
func (s *Activity) Create() error {
	return DB.Create(s).Error
}

//获取所有活动列表
func (s *Activity) GetAllActivities() ([]Activity, error) {
	var t []Activity
	err := DB.Model(s).Order("start_time DESC").Find(&t).Error
	for _, content := range t {
		err = content.SetCacheActivity()

	}
	return t, err
}

//获取活动详细信息
func (s *Activity) GetActivity() (Activity, error) {
	var t Activity
	err := DB.Model(s).Find(&t).Error
	return t, err
}

//按地点获取活动
func (s *Activity) GetActivitiesByPlace() ([]Activity, error) {
	tim := time.Now().Unix()
	var t []Activity
	err := DB.Model(s).Where(s).Where("end_time > ?", tim).Order("start_time DESC").Find(&t).Error
	return t, err
}

//按组织获取活动
func (s *Activity) GetActivitiesByHost() ([]Activity, error) {
	var t []Activity
	err := DB.Model(s).Where(s).Order("start_time DESC").Find(&t).Error
	return t, err
}

//redis
//建立
func (s *Activity) SetCacheActivity() error {
	fmt.Println(s)
	b := strconv.Itoa(int(s.ID))
	key := CacheConfigObj.Prefix + "activity" + b
	DataBytes, err := json.Marshal(s)
	if err != nil {
		return err
	}
	err = Cache.Set(key, string(DataBytes), ActivityCacheTime).Err()
	Cache.Expire(key, ActivityCacheTime)
	return err
}

func (s *Activity) SetCacheActivityList() error {
	key := CacheConfigObj.Prefix + "activity_list"
	data, err := s.GetAllActivities()
	if err != nil {
		return err
	}

	for _, temp := range data {
		err = Cache.SAdd(key, temp.ID).Err()
		Cache.Expire(key, ActivityCacheTime)
		if err != nil {
			return err
		}
	}

	return err
}

func (s *Activity) SetCachePlaceList() error {
	key := CacheConfigObj.Prefix + "activity_place"
	data, err := s.GetActivitiesByPlace()
	if err != nil {
		return err
	}

	for _, temp := range data {
		err = Cache.SAdd(key, temp.ID).Err()
		Cache.Expire(key, ActivityCacheTime)
		if err != nil {
			return err
		}
	}

	return err
}

func (s *Activity) SetCacheHostList() error {
	key := CacheConfigObj.Prefix + "activity_host"
	data, err := s.GetActivitiesByHost()
	if err != nil {
		return err
	}

	for _, temp := range data {
		err = Cache.SAdd(key, temp.ID).Err()
		Cache.Expire(key, ActivityCacheTime)
		if err != nil {
			return err
		}
	}

	return err
}

//获取

func (s *Activity) GetCacheActivity() (interface{}, error) {
	key := CacheConfigObj.Prefix + "activity" + strconv.Itoa(int(s.ID))

	value, err := Cache.Get(key).Result()
	if err != nil {
		return nil, err
	}
	var t interface{}
	err = json.Unmarshal([]byte(value), &t)

	if err != nil {
		return nil, err
	}

	return t, err
}

func (s *Activity) GetCacheAllActivities() ([]string, error) {
	key := CacheConfigObj.Prefix + "activity_list"
	members, err := Cache.SCard(key).Result()
	if err != nil {
		return nil, err
	}
	if members != 0 {
		val, err := Cache.SMembers(key).Result()
		return val, err
	}

	return nil, errors.New("false")
}

func (s *Activity) GetCacheActivitiesByPlace() ([]string, error) {
	key := CacheConfigObj.Prefix + "activity_place"
	members, err := Cache.SCard(key).Result()
	if err != nil {
		return nil, err
	}
	if members != 0 {
		val, err := Cache.SMembers(key).Result()
		return val, err
	}

	return nil, errors.New("false")
}

func (s *Activity) GetCacheActivitiesByHost() ([]string, error) {
	key := CacheConfigObj.Prefix + "activity_host"
	members, err := Cache.SCard(key).Result()
	if err != nil {
		return nil, err
	}
	if members != 0 {
		val, err := Cache.SMembers(key).Result()
		return val, err
	}

	return nil, errors.New("false")
}

//删除缓存
func (s *Activity) DelCacheList(name string) error {
	key := CacheConfigObj.Prefix + "activity_" + name
	_, err := Cache.Del(key).Result()
	return err
}
