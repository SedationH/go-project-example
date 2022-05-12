package service

import (
	"errors"
	"go-project-example/repository"
	"sync"
)

// type PageInfo struct {
// 	Topic    *repository.Topic
// 	PostList []*repository.Post
// }

type PageInfo struct {
	Topic    *repository.Topic
	PostList []*repository.Post
}

func QueryPageInfo(topicId int64) (*PageInfo, error) {
	return NewQueryPageInfoFlow(topicId).Do()
}

type QueryPageInfoFlow struct {
	topicId  int64
	pageInfo *PageInfo

	topic *repository.Topic
	posts []*repository.Post
}

func NewQueryPageInfoFlow(topId int64) *QueryPageInfoFlow {
	return &QueryPageInfoFlow{
		topicId: topId,
	}
}

func (f *QueryPageInfoFlow) Do() (*PageInfo, error) {
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	if err := f.prepareInfo(); err != nil {
		return nil, err
	}
	if err := f.packPageInfo(); err != nil {
		return nil, err
	}
	return f.pageInfo, nil
}

func (f *QueryPageInfoFlow) checkParam() error {
	if f.topicId <= 0 {
		return errors.New("topic id must be larger than 0")
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo() error {
	// 去并行获取 topic 和 post 信息
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		topic := repository.NewTopicDaoInstance().QueryTopicById(f.topicId)
		f.topic = topic
	}()
	go func() {
		defer wg.Done()
		posts := repository.NewPostDaoInstance().QueryPostsByParentId(f.topicId)
		f.posts = posts
	}()

	wg.Wait()
	return nil
}

func (f *QueryPageInfoFlow) packPageInfo() error {
	f.pageInfo = &PageInfo{
		Topic:    f.topic,
		PostList: f.posts,
	}

	return nil
}

func AddNewPage(newPageInfo *PageInfo) error {
	// TODO
	return nil
}
