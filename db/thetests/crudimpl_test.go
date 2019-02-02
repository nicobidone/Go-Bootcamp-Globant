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
	s, _ := x.Read("a")
	if s.AValue() != "a" {
		t.Error("The value A in not ok ", s.AValue())
	}
	if s.BValue() != "b" {
		t.Error("The value B is not ok ", s.BValue())
	}

	c := *thedata.NewContainer("c", "d")
	x.Update("a", c)
	s, _ = x.Read("a")
	if s.AValue() != "a" {
		t.Error("The value A in not ok ", s.AValue())
	}
	if s.BValue() != "b" {
		t.Error("The value B is not ok ", s.BValue())
	}

	x.Delete("a")

	s, e := x.Read("a")
	if e != nil {
		t.Error(e)
	}

}
