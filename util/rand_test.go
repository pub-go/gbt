package util_test

import (
	"testing"

	"code.gopub.tech/commons/assert"
	"code.gopub.tech/gbt/util"
)

func TestRand(t *testing.T) {
	s := util.RandStr(10)
	t.Logf("rand: %v", s)
	assert.True(t, len(s) == 10)

	s = util.RandString(10, "abcd-0123")
	t.Logf("rand: %v", s)
	assert.True(t, len(s) == 10)

	b := util.RandBytes(10)
	t.Logf("%v %s", b, b)
	assert.True(t, len(b) == 10)
}
