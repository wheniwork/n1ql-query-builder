package nqb

import (
	"bytes"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildCondition_Error(t *testing.T) {
	buf := &bytes.Buffer{}
	err := buildCondition(buf, "", func(buf *bytes.Buffer) error {
		return nil
	})
	assert.Error(t, err)
}

func TestBuildCondition_BuildFunc_Error(t *testing.T) {
	buf := &bytes.Buffer{}
	err := buildCondition(buf, "foo", func(buf *bytes.Buffer) error {
		return errors.New("something went wrong")
	})
	assert.Error(t, err)
}
