package fpl

import (
	"testing"
)

func TestListTransfers(t *testing.T) {

	c := NewClient(nil)

	got, _ := c.ListTransfers("12")
	if got != nil {
		t.Errorf("Could do not be succed, got %+v", got)
	}
}
