package nqb_test

import (
	"fmt"
	"testing"

	. "github.com/wheniwork/n1ql-query-builder"
)

func TestSelect(t *testing.T) {
	query := Select(X("foo")).String()
	fmt.Println(query)
}
