package errorutil

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

var SampleValueError = errors.New("sample value error")

type SampleTypeError struct {
}

func (e SampleTypeError) Error() string {
	return "sample type error"
}

func TestUnwrap(t *testing.T) {
	err := errors.New("sample error")
	assert.Equal(t, errors.New("sample error"), err)

	err = errors.Unwrap(err)
	assert.Nil(t, err)
	assert.NoError(t, err)
}

func TestUnwrap2(t *testing.T) {
	err := SampleValueError
	assert.Equal(t, err, errors.New("sample value error"))

	err = fmt.Errorf("action=some action err=%w", err)

	err = errors.Unwrap(err)

	assert.Error(t, err)
	assert.ErrorIs(t, err, SampleValueError)
}

func TestErrorsIsMultipleWrapError(t *testing.T) {
	err := SampleValueError
	err1 := fmt.Errorf("wrap=1 err=[%w]", err)
	err2 := fmt.Errorf("wrap=2 err=[%w]", err1)
	err3 := fmt.Errorf("wrap=3 err=[%w]", err2)

	fmt.Println("err\t\t:", err)
	fmt.Println("err1\t:", err1)
	fmt.Println("err2\t:", err2)
	fmt.Println("err3\t:", err3)

	is := errors.Is(err, SampleValueError)
	is1 := errors.Is(err1, SampleValueError)
	is2 := errors.Is(err2, SampleValueError)
	is3 := errors.Is(err3, SampleValueError)

	assert.True(t, is)
	assert.True(t, is1)
	assert.True(t, is2)
	assert.True(t, is3)
}

func TestErrorsAsMultipleWrapError(t *testing.T) {
	err := SampleTypeError{}
	err1 := fmt.Errorf("wrap=1 err=[%w]", err)
	err2 := fmt.Errorf("wrap=2 err=[%w]", err1)
	err3 := fmt.Errorf("wrap=3 err=[%w]", err2)

	fmt.Println("err\t\t:", err)
	fmt.Println("err1\t:", err1)
	fmt.Println("err2\t:", err2)
	fmt.Println("err3\t:", err3)

	var sampleTypeError SampleTypeError
	as := errors.As(err, &SampleTypeError{})
	as1 := errors.As(err1, &sampleTypeError)
	as2 := errors.As(err2, &sampleTypeError)
	as3 := errors.As(err3, &sampleTypeError)

	assert.True(t, as)
	assert.True(t, as1)
	assert.True(t, as2)
	assert.True(t, as3)

	assert.ErrorAs(t, err, &sampleTypeError)
	assert.ErrorAs(t, err1, &sampleTypeError)
	assert.ErrorAs(t, err2, &sampleTypeError)
	assert.ErrorAs(t, err3, &sampleTypeError)
}
