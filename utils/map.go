package expire_map

import (
	"sync"
	"time"
)

type Value struct {
	data        interface{}
	expiredTime int64
}

type ExpiredMap struct {
	m map[interface{}]*Value
	// 过期时间作为key，放在map中
	timeMap map[int64][]interface{}
	lock    *sync.Mutex
	stop    chan struct{}
}

func NewExpiredMap() *ExpiredMap {
	e := ExpiredMap{
		m:       make(map[interface{}]*Value),
		timeMap: make(map[int64][]interface{}),
		lock:    new(sync.Mutex),
		stop:    make(chan struct{}),
	}
	go e.run(time.Now().Unix())
	return &e
}

type deleteMsg struct {
	keys []interface{}
	time int64
}

//background goroutine 主动删除过期的key
//数据实际删除时间比应该删除的时间稍晚一些，这个误差会在查询的时候被解决。
func (e *ExpiredMap) run(now int64) {
	t := time.NewTicker(time.Second * 1)
	defer t.Stop()

	deleteChannel := make(chan *deleteMsg)

	// 从channel接收，并进行删除
	go func() {
		for v := range deleteChannel {
			e.DeleteMulti(v.keys, v.time)
		}
	}()

	// 遍历timeMap，将每秒过期的传送到channel
	for {
		select {
		case <-t.C:
			now++ //这里用now++的形式，直接用time.Now().Unix()可能会导致时间跳过1s，导致key未删除。
			e.lock.Lock()
			if keys, found := e.timeMap[now]; found {
				deleteChannel <- &deleteMsg{
					keys: keys,
					time: now,
				}
			}
			e.lock.Unlock()
		case <-e.stop:
			close(deleteChannel)
			return
		}
	}
}

func (e *ExpiredMap) Put(key, val interface{}, expireSecond int64) {
	if expireSecond <= 0 {
		return
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	expiredTime := time.Now().Unix() + expireSecond
	e.m[key] = &Value{
		data:        val,
		expiredTime: expiredTime,
	}
	e.timeMap[expiredTime] = append(e.timeMap[expireSecond], e.m[key])
}

func (e *ExpiredMap) Get(key interface{}) (value interface{}, found bool) {
	e.lock.Lock()
	defer e.lock.Unlock()
	if found = e.checkDeleteKey(key); !found {
		return nil, false
	}
	return e.m[key].data, true
}

func (e *ExpiredMap) GetOrDefault(key, defaultValue interface{}) (value interface{}, found bool) {
	e.lock.Lock()
	defer e.lock.Unlock()
	if found = e.checkDeleteKey(key); !found {
		return defaultValue, false
	}
	return e.m[key].data, true
}

func (e *ExpiredMap) Delete(key interface{}) {
	e.lock.Lock()
	defer e.lock.Unlock()
	delete(e.m, key)
}

func (e *ExpiredMap) DeleteMulti(keys []interface{}, time int64) {
	e.lock.Lock()
	defer e.lock.Unlock()
	for _, key := range keys {
		delete(e.timeMap, time)
		// 这里调用checkDeleteKey，防止出现对一个key反复Put误删
		e.checkDeleteKey(key)
	}
}

func (e *ExpiredMap) Length() int {
	e.lock.Lock()
	defer e.lock.Unlock()
	return len(e.m)
}

// TTL 返回key的剩余生存时间 key不存在返回-1
func (e *ExpiredMap) TTL(key interface{}) int64 {
	e.lock.Lock()
	defer e.lock.Unlock()
	if !e.checkDeleteKey(key) {
		return -1
	}
	return e.m[key].expiredTime - time.Now().Unix()
}

func (e *ExpiredMap) Clear() {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.m = make(map[interface{}]*Value)
	e.timeMap = make(map[int64][]interface{})
}

func (e *ExpiredMap) Foreach(handle func(key, val interface{})) {
	e.lock.Lock()
	defer e.lock.Unlock()
	for k, v := range e.m {
		if !e.checkDeleteKey(k) {
			continue
		}
		handle(k, v.data)
	}
}

func (e *ExpiredMap) checkDeleteKey(key interface{}) bool {
	if val, found := e.m[key]; found {
		if val.expiredTime <= time.Now().Unix() {
			delete(e.m, key)
			return false
		}
		return true
	}
	return false
}
