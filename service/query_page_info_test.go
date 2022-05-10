package service

import (
	"encoding/json"
	"go-project-example/repository"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 注意这里是 m
func TestMain(m *testing.M) {
	repository.Init("../data/")
	os.Exit(m.Run())
}

func TestQueryPageInfo(t *testing.T) {
	pageInfo, err := QueryPageInfo(1)
	pageInfo2, err1 := QueryPageInfo(-1)
	res, _ := json.Marshal(pageInfo)
	t.Log(string(res), err, pageInfo2, err1)

	expectedLessThanZeroErrorMsg := "topic id must be larger than 0"
	assert.Equal(t, expectedLessThanZeroErrorMsg, err1.Error())
	assert.Equal(t, 5, len(pageInfo.PostList))
}
