package repository

import (
	"log"
	"sync"
	"time"
)

type Post struct {
	Id         int64     `json:"id"`
	ParentId   int64     `json:"parent_id"`
	Content    string    `json:"content"`
	CreateTime time.Time `json:"create_time"`
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

func (*PostDAO) QueryPostsByParentId(parentId int64) ([]*Post, error) {
	var posts []*Post
	err := db.Where("parent_id = ?", parentId).Find(&posts).Error
	if err != nil {
		log.Println("find posts by parent_id err:" + err.Error())
		return nil, err
	}
	return posts, nil
}

func (*PostDAO) AddNewPost(postList []*Post, parentId int64) error {
	for _, post := range postList {
		post.CreateTime = time.Now()
		post.ParentId = parentId
	}
	result := db.Create(postList)
	err := result.Error
	if err != nil {
		log.Println("insert post err:" + err.Error())
	}
	return err
}
