package intset

import (
	"testing"
)

func TestAdd(t *testing.T) {
	var x IntSet
	x.Add(1)
	if x.String() != "{1}" {
		t.Errorf("Got %s, should be %s", x.String(), "{1}")
	}
	x.Add(235)
	x.Add(25)
	if x.String() != "{1 25 235}" {
		t.Errorf("Got %s, should be %s", x.String(), "{1 25 235}")
	}
}

func TestAddAll(t *testing.T) {
	var x IntSet
	x.AddAll(1, 235, 25)
	if x.String() != "{1 25 235}" {
		t.Errorf("Got %s, should be %s", x.String(), "{1 25 235}")
	}
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.AddAll(1, 235, 25)
	x.Remove(235)
	if x.String() != "{1 25}" {
		t.Errorf("Got %s, should be %s", x.String(), "{1 25}")
	}
}

func TestUnion(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 235, 25)
	y.AddAll(1, 50, 100)
	x.UnionWith(&y)
	if x.String() != "{1 25 50 100 235}" {
		t.Errorf("Got %s, should be %s", x.String(), "{1 25 50 100 235}")
	}
}

func TestIntersect(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 235, 25)
	y.AddAll(1, 100)
	x.IntersectWith(&y)
	if x.String() != "{1}" {
		t.Errorf("Got %s, should be %s", x.String(), "{1}")
	}
}

func TestDifference(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 300, 45, 10)
	y.AddAll(1, 200)
	x.DifferenceWith(&y)
	if x.String() != "{10 45 300}" {
		t.Errorf("Got %s, should be %s", x.String(), "{10 45 300}")
	}

	x, y = IntSet{}, IntSet{}
	y.AddAll(1, 300, 45, 10)
	x.AddAll(1, 200)
	x.DifferenceWith(&y)
	if x.String() != "{200}" {
		t.Errorf("Got %s, should be %s", x.String(), "{200}")
	}
}

func TestSymmetricDifference(t *testing.T) {
	var x, y IntSet
	x.AddAll(1, 300, 10)
	y.AddAll(1, 200, 10, 23)
	x.SymmetricDifference(&y)
	if x.String() != "{23 200 300}" {
		t.Errorf("Got %s, should be %s", x.String(), "{23 200 300}")
	}

	x, y = IntSet{}, IntSet{}
	x.AddAll(1, 300, 10)
	y.AddAll(1)
	x.SymmetricDifference(&y)
	if x.String() != "{10 300}" {
		t.Errorf("Got %s, should be %s", x.String(), "{10 300}")
	}
}
