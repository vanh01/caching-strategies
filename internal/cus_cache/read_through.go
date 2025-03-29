package cuscache

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var HasExpired error = errors.New("Has expired")
var NotFound error = errors.New("Not found")
var InvalidType error = errors.New("Invalid type")

var memCahe = &ReadThroughCache{
	memoryCache: make(map[string]*vals),
	locks:       make(map[string]*sync.Mutex),
}

type vals struct {
	val     any
	expried time.Time
}

type ReadThroughCache struct {
	db          *gorm.DB
	memoryCache map[string]*vals
	locks       map[string]*sync.Mutex
}

func New(db *gorm.DB) {
	memCahe.db = db
}

func Set[T any](key string, t T, ttl *time.Duration) error {
	// lok := memCahe.locks["set"+key]

	// if lok == nil {
	// 	lok = &sync.Mutex{}
	// 	memCahe.locks["set"+key] = lok
	// }

	// lok.Lock()
	// defer lok.Unlock()

	expried := time.Now().Add(*ttl)
	memCahe.memoryCache[key] = &vals{
		val:     t,
		expried: expried,
	}

	return nil
}

func Get[T any](key string, ttl *time.Duration) (*T, error) {
	lok := memCahe.locks["get"+key]

	if lok == nil {
		lok = &sync.Mutex{}
		memCahe.locks["get"+key] = lok
	}

	lok.Lock()
	defer lok.Unlock()

	now := time.Now()
	v, ok := memCahe.memoryCache[key]
	if !ok || v.expried.Before(now) {
		fmt.Println("Get from db")
		t, err := getFromDB[T](key)
		if err != nil {
			return nil, err
		}

		Set(key, &t, ttl)
		return &t, nil
	}

	res, ok := v.val.(*T)
	if !ok {
		return nil, InvalidType
	}

	return res, nil
}

func getFromDB[T any](key string) (T, error) { // mock firstly
	var t T
	id, err := uuid.Parse(key)
	if err != nil {
		return t, err
	}

	res := memCahe.db.Where("id = ?", id).Take(&t)

	return t, res.Error
}
