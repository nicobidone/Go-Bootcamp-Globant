package thetest

import (
	"testing"

	"github.com/nicob/db/thedata"
	"github.com/nicob/db/themap"
)

func TestCreate(t *testing.T) {
	x := themap.NewMap()
	b := *thedata.NewContainer("a", "b")
	x.Create("a", b)
	s := x.Read("a")
	if s.AValue() != "a" {
		t.Errorf("The value A in not ok")
	}
	if s.BValue() != "b" {
		t.Errorf("The value B is not ok")
	}
}
