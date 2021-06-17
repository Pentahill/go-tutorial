package main

import (
	"testing"

	"github.com/marmotedu/errors"
)

func testErrorCode(t *testing.T) {
	errors.WithCode(1, "user 'Lingfei Kong' not found.")
}
