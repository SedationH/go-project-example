package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryPageInfo(t *testing.T) {
	pageInfo, err := QueryPageInfo(1)
	pageInfo2, err1 := QueryPageInfo(-1)
	t.Log(pageInfo, err, pageInfo2, err1)

	expectedLessThanZeroErrorMsg := "topic id must be larger than 0"
	assert.Equal(t, expectedLessThanZeroErrorMsg, err1.Error())
}
