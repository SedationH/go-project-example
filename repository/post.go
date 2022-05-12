package repository

import (
	"encoding/json"
	"os"
	"sync"
	"time"
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
	v, _ := postIndexMap.Get(parentId)
	return v
}

func (*PostDAO) AddNewPost(postList []*Post, parentId int64) error {
	for _, post := range postList {
		maxPostId++
		post.Id = maxPostId
		post.ParentId = parentId
		post.CreateTime = time.Now().Unix()
	}
	postIndexMap.Set(parentId, postList)
	f, err := os.OpenFile("./data/post", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	var b []byte
	for _, post := range postList {
		f.WriteString("\n")
		b, _ = json.Marshal(post)
		_, err = f.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}
