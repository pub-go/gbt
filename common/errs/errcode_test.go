package errs_test

import (
	"fmt"
	"testing"

	"code.gopub.tech/commons/assert"
	"code.gopub.tech/errors"
	"code.gopub.tech/gbt/common/errs"
)

var errFmt = fmt.Errorf("format error")

type code int

func (c code) Code() int       { return int(c) }
func (c code) Message() string { return fmt.Sprintf("my code is %d", int(c)) }
func (c code) Error() string   { return c.Message() }

var errCode = code(123)
var errStack = errors.Errorf("i have stack")

func TestOf(t *testing.T) {
	var err error
	assert.Nil(t, err)
	assert.True(t, err == nil)

	err = errs.Of(nil)
	assert.Nil(t, err)
	assert.True(t, err == nil)

	err = errs.Of(errFmt)
	assert.NotNil(t, err)
	assert.True(t, err != nil)
	assert.True(t, errors.Is(err, errFmt))

	err = errs.Of(errCode)
	assert.NotNil(t, err)
}

func TestOr(t *testing.T) {
	var err error
	assert.Nil(t, err)

	err = errs.Or(nil, errs.ErrBadRequest)
	assert.Nil(t, err)

	err = errs.Or(errFmt, errs.ErrBadRequest)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, errs.ErrBadRequest))

	err = errs.Or(errCode, errs.ErrBadRequest)
	assert.NotNil(t, err)
	assert.True(t, errors.Is(err, errCode))
}

func TestErrCode(t *testing.T) {
	var err error
	err = errs.New(123, "basic error")
	t.Logf("err=%v", err)
	t.Logf("err+v=%+v", err)
	t.Logf("err+Wrap=%+v", errors.Wrapf(err, "prefix"))

	err = errs.ErrBadRequest.WithCause(errStack)
	t.Logf("WithCause=%+v", err)

	err = errs.ErrBadRequest.WithCause(errCode)
	t.Logf("WithCauseCode=%+v", err)
	assert.True(t, errors.Is(err, errCode))

	err = errs.ErrBadRequest.WithCause(nil)
	assert.True(t, err == nil)
}
