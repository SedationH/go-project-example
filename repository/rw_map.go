// 以下实现参考《Go 并发之三种线程安全的 map - AFreeCoder的文章 - 知乎》
// 仅仅实现 Get Set 的锁
// 链接: https://zhuanlan.zhihu.com/p/356739568

package repository

import (
	"sync"
)

type RWTopicMap struct {
	sync.RWMutex
	m map[int64]*Topic
}

func NewRWTopicMap(m map[int64]*Topic) *RWTopicMap {
	return &RWTopicMap{
		m: m,
	}
}

func (rWTopicMap *RWTopicMap) Get(k int64) (*Topic, bool) {
	rWTopicMap.RLock()
	defer rWTopicMap.RUnlock()
	v, existed := rWTopicMap.m[k]
	return v, existed
}

func (rWTopicMap *RWTopicMap) Set(k int64, v *Topic) {
	rWTopicMap.Lock()
	defer rWTopicMap.Unlock()
	rWTopicMap.m[k] = v
}

type RWPostMap struct {
	sync.RWMutex
	m map[int64][]*Post
}

func NewRWPostMap(m map[int64][]*Post) *RWPostMap {
	return &RWPostMap{
		m: m,
	}
}

func (rWPostMap *RWPostMap) Get(k int64) ([]*Post, bool) {
	rWPostMap.RLock()
	defer rWPostMap.RUnlock()
	v, existed := rWPostMap.m[k]
	return v, existed
}

func (rWPostMap *RWPostMap) Set(k int64, v []*Post) {
	rWPostMap.Lock()
	defer rWPostMap.Unlock()
	rWPostMap.m[k] = v
}
