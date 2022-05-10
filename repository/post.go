package repository

import (
	"sync"
)

type Post struct {
	Id         int64  `json:"id"`
	ParentId   int64  `json:"parent_id"`
	Content    string `json:"content"`
	CreateTime int64  `json:"create_time"`
}

type PostDAO struct {
}

var (
	postDAO  *PostDAO
	postOnce sync.Once
)

func NewPostDaoInstance() *PostDAO {
	postOnce.Do(
		func() {
			postDAO = &PostDAO{}
		})
	return postDAO
}
func (*PostDAO) QueryPostsByParentId(parentId int64) []*Post {
	return postIndexMap[parentId]
}
