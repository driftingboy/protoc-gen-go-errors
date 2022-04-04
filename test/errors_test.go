package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	e := ErrorTestnotfound("resource not found")
	assert.True(t, IsTestnotfound(e))
	assert.Equal(t, 100001, BizErrorCode(e))
}
