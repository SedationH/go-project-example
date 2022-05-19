# 相关阅读

v0 与 v1 的内容 看 [2. 课程记录-Go语言上手-工程实践｜ 青训营笔记](https://juejin.cn/post/7096848424917532680)

v2 的内容看 [3. 课程记录-Go与数据库｜ 青训营笔记](https://juejin.cn/post/7099397274337280014) 

# v0

[重要参考](https://juejin.cn/post/7095327443585597453#heading-12)

[ER 图](https://app.diagrams.net/#G1LZJdPQrZxJDsCtTSmlu8BvFkQWEz0B8L)

![](https://s2.loli.net/2022/05/12/wGYj5y8A1cvJE7B.png)

# v1

## 从 Server 开始

![image-20220512102327270](https://s2.loli.net/2022/05/12/oeT4lMKyShuvfwd.png)



## 视图层

![image-20220512104447590](https://s2.loli.net/2022/05/12/IWnJZ6FANYx5KLt.png)

## 逻辑层

注意 PostList 中的 parentID 和 topic 的 id产生了关联

![image-20220512110504773](https://s2.loli.net/2022/05/12/IrqC4anfgHVsvWE.png)



## 数据层

成功完成了post  和 topic id的自增行为, 这一处涉及 map 锁的行为

![image-20220512122753592](https://s2.loli.net/2022/05/12/mvfRBTxu6VjkhtL.png)



![image-20220512122906550](https://s2.loli.net/2022/05/12/t2CKABDF8WqIkG7.png)

### 关注 map 锁的处理

以 topic 为例子，嵌套一层读写锁的处理

```go
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
```

