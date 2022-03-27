package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/putukrisna6/golang-api/entity"
)

type ReceiptCache interface {
	Set(key string, value []entity.Receipt)
	Get(key string) []entity.Receipt
	Del(key string)
}

type receiptCache struct {
	cache   *redis.Client
	expires time.Duration
}

func NewReceiptCache(cache *redis.Client, expires time.Duration) ReceiptCache {
	return &receiptCache{
		cache:   cache,
		expires: expires,
	}
}

func (rc *receiptCache) Set(key string, value []entity.Receipt) {
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	rc.cache.Set(context.TODO(), "receipt:"+key, json, rc.expires*time.Second)
}

func (rc *receiptCache) Get(key string) []entity.Receipt {
	val, errGet := rc.cache.Get(context.TODO(), "receipt:"+key).Result()

	if errGet == redis.Nil {
		return nil
	} else if errGet != nil {
		panic(errGet)
	}

	var receipt []entity.Receipt
	err := json.Unmarshal([]byte(val), &receipt)
	if err != nil {
		panic(err)
	}

	return receipt
}

func (rc *receiptCache) Del(key string) {
	err := rc.cache.Del(context.TODO(), "receipt:"+key)

	log.Println(err)
}
