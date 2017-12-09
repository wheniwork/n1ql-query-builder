package nqb

import (
	"bytes"
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
