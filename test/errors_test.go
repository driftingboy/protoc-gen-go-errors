package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrors(t *testing.T) {
	e := ErrorTestNotFound("resource not found")
	assert.True(t, IsTestNotFound(e))
	assert.Equal(t, 100001, BizErrorCode(e))
}
