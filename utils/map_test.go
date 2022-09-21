package expire_map

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMap(t *testing.T) {
	cache := NewExpiredMap()

	for i := 1; i <= 100; i++ {
		cache.Put(i, i, int64(i))
	}
	cache.Delete(8)
	cache.Delete(9)
	cache.Delete(10)

	time.Sleep(time.Second * 10)
	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond * 1)
		fmt.Println(cache.Get(i))
	}
}

func TestExpiredMap_Foreach(t *testing.T) {
	cache := NewExpiredMap()
	for i := 1; i <= 10; i++ {
		cache.Put(i, i, int64(i))
	}
	cache.Foreach(func(key, val interface{}) {
		fmt.Println("key: ", key, "Value: ", val)
	})
}

func TestExpiredMap_Get(t *testing.T) {
	cache := NewExpiredMap()
	value := "geek"
	cache.Put(value, value, 10)
	result, _ := cache.Get(value)
	assert.Equal(t, value, result)
}

func TestExpiredMap_TTL(t *testing.T) {
	cache := NewExpiredMap()
	cache.Put(1, 1, 10)
	time.Sleep(time.Second)
	assert.Equal(t, int64(9), cache.TTL(1))
}

func TestExpiredMap_Clear(t *testing.T) {
	cache := NewExpiredMap()
	for i := 1; i <= 100; i++ {
		cache.Put(i, i, int64(i))
	}
	assert.Equal(t, 100, cache.Length())
	cache.Clear()
	assert.Equal(t, 0, cache.Length())
}
