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
		cache.Set(i, i, int64(i))
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

func TestExpiredMap_TTL(t *testing.T) {
	cache := NewExpiredMap()
	cache.Set(1, 1, 10)
	time.Sleep(time.Second)
	assert.Equal(t, int64(9), cache.TTL(1))
}

func TestExpiredMap_Clear(t *testing.T) {
	cache := NewExpiredMap()
	for i := 1; i <= 100; i++ {
		cache.Set(i, i, int64(i))
	}
	assert.Equal(t, 100, cache.Length())
	cache.Clear()
	assert.Equal(t, 0, cache.Length())
}
