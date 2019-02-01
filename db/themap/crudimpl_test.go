package themap

import (
	"testing"

	"github.com/nicob/db/thedata"
)

func TestCreate(t *testing.T) {
	x := NewMap()
	b := *thedata.NewContainer("a", "b")
	x.Create("a", b)
	s := x.Read("a")
	if s.AValue() != "b" {
		t.Errorf("The value in not ok")
	}
}
