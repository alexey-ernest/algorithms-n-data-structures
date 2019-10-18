package ch3

import "testing"

func TestSequentialSearchSTBasic(t *testing.T) {
	st := NewSequentialSearchST()
	for i := 0; i < 10; i+=1 {
		k := string('a' + i)
		v := string(i)
		st.Put(k, v)
	}

	if st.Size() != 10 {
		t.Errorf("size should equals 10")
	}
	if st.IsEmpty() {
		t.Errorf("st should not be empty")	
	}

	for i := 0; i < 10; i+=1 {
		k := string('a' + i)
		v := string(i)
		if !st.Contains(k) {
			t.Errorf("st should contain the key %s", k)
		}
		if st.Get(k) != v {
			t.Errorf("value %s != %s", st.Get(k), v)
		}
	}

	st.Delete(string('a'))
	if st.Size() != 9 {
		t.Errorf("st size should be 9")
	}
	if st.Contains(string('a')) {
		t.Errorf("st should not contain deleted key")
	}
}

func TestBinarySearchSTSingle(t *testing.T) {
	st := NewBinarySearchST()
	key := "key"
	val := "val"
	st.Put(key, val)
	if st.IsEmpty() {
		t.Errorf("st should not be empty")
	}
	if st.Size() != 1 {
		t.Errorf("st size should be 1")
	}
	if !st.Contains(key) {
		t.Errorf("st should contain %s", key)
	}
	st.Delete(key)
	if !st.IsEmpty() || st.Size() != 0 {
		t.Errorf("st should be empty now")
	}
}

func TestBinarySearchSTBasic(t *testing.T) {
	st := NewBinarySearchST()
	for i := 0; i < 10; i+=1 {
		k := string('a' + 10 - i)
		v := string(i)
		st.Put(k, v)
	}

	if st.Size() != 10 {
		t.Errorf("size should equals 10")
	}
	if st.IsEmpty() {
		t.Errorf("st should not be empty")	
	}

	for i := 0; i < 10; i+=1 {
		k := string('a' + 10 - i)
		v := string(i)
		if !st.Contains(k) {
			t.Errorf("st should contain the key %s", k)
		}
		if st.Get(k) != v {
			t.Errorf("value %s != %s", st.Get(k), v)
		}
	}

	st.Delete(string('a' + 5))
	if st.Size() != 9 {
		t.Errorf("st size should be 9")
	}
	if st.Contains(string('a' + 5)) {
		t.Errorf("st should not contain deleted key")
	}
}

func TestBinarySearchSTMinMax(t *testing.T) {
	st := NewBinarySearchST()
	for i := 0; i < 10; i+=1 {
		k := string('a' + 10 - i)
		v := string(i)
		st.Put(k, v)
	}

	min := string('a' + 1)
	if st.Min() != min {
		t.Errorf("min != %s", min)
	}

	max := string('a' + 10)
	if st.Max() != max {
		t.Errorf("max != %s", max)
	}
}

func TestBinarySearchSTFloorCeiling(t *testing.T) {
	st := NewBinarySearchST()
	for i := 0; i < 10; i += 1 {
		k := string('a' + 20 - 2*i)
		v := string(i)
		st.Put(k, v)
	}

	key := string('a' + 3)
	fl := string('a' + 2)
	ceiling := string('a' + 4)
	if st.Floor(key) != fl {
		t.Errorf("floor != %s", st.Floor(key))
	}
	if st.Ceiling(key) != ceiling {
		t.Errorf("ceiling != %s", st.Ceiling(key))
	}
}

func TestBinarySearchSTSelect(t *testing.T) {
	st := NewBinarySearchST()
	for i := 0; i < 10; i+=1 {
		k := string('a' + 10 - i)
		v := string(i)
		st.Put(k, v)
	}

	key := string('a' + 3)
	if st.Select(2) != key {
		t.Errorf("element with rank=2 should be %s", key)
	}
}

func TestBinarySearchSTKeys(t *testing.T) {
	st := NewBinarySearchST()
	for i := 0; i < 10; i+=1 {
		k := string('a' + 10 - i)
		v := string(i)
		st.Put(k, v)
	}

	lo := string('a' + 3)
	hi := string('a' + 6)
	keys := st.Keys(lo, hi)
	if len(keys) != 4 {
		t.Errorf("keys len should equals 4")
	}

	if keys[0] != lo {
		t.Errorf("first key should be %s", lo)
	}

	if keys[len(keys)-1] != hi {
		t.Errorf("last key should be %s", hi)
	}

	for i := 1; i < len(keys); i += 1 {
		if keys[i] < keys[i-1] {
			t.Errorf("non-decreasing keys order validation failed")
		}
	}
}

func TestBSTBasic(t *testing.T) {
	st := NewBST()
	for i := 0; i < 10; i+=1 {
		k := string('a' + i)
		v := string(i)
		st.Put(k, v)
	}

	if st.Size() != 10 {
		t.Errorf("size should equals 10")
	}
	if st.IsEmpty() {
		t.Errorf("st should not be empty")	
	}

	for i := 0; i < 10; i+=1 {
		k := string('a' + i)
		v := string(i)
		if !st.Contains(k) {
			t.Errorf("st should contain the key %s", k)
		}
		if st.Get(k) != v {
			t.Errorf("value %s != %s", st.Get(k), v)
		}
	}
}

func TestBSTMinMax(t *testing.T) {
	st := NewBST()
	for i := 0; i < 10; i+=1 {
		k := string('a' + 9 - i)
		v := string(i)
		st.Put(k, v)
	}

	min := string('a')
	if st.Min() != min {
		t.Errorf("min %q != %q", st.Min(), min)
	}

	max := string('a' + 9)
	if st.Max() != max {
		t.Errorf("min %q != %q", st.Max(), max)
	}
}

func TestBSTFloor(t *testing.T) {
	st := NewBST()
	for i, v := range []string{"S", "E", "X", "A", "R", "C", "H", "M"} {
		st.Put(v, string(i))
	}

	floorG := "E"
	if st.Floor("G") != floorG {
		t.Errorf("%q != %q", st.Floor("G"), floorG)
	}
}

func TestBSTCeiling(t *testing.T) {
	st := NewBST()
	for i, v := range []string{"S", "E", "X", "A", "R", "C", "H", "M"} {
		st.Put(v, string(i))
	}

	ceilingP := "R"
	if st.Ceiling("P") != ceilingP {
		t.Errorf("%q != %q", st.Ceiling("P"), ceilingP)
	}
}

func TestBSTSelect(t *testing.T) {
	st := NewBST()
	for i, v := range []string{"S", "E", "X", "A", "R", "C", "H", "M"} {
		st.Put(v, string(i))
	}

	exp0 := "A"
	if st.Select(0) != exp0 {
		t.Errorf("%q != %q", st.Select(0), exp0)
	}

	exp1 := "H"
	if st.Select(3) != exp1 {
		t.Errorf("%q != %q", st.Select(3), exp1)
	}

	exp2 := "X"
	if st.Select(7) != exp2 {
		t.Errorf("%q != %q", st.Select(7), exp2)
	}
}

func TestBSTRank(t *testing.T) {
	st := NewBST()
	nodes := []string{"S", "E", "X", "A", "R", "C", "H", "M"}
	for i, v := range nodes {
		st.Put(v, string(i))
	}

	for i := range nodes {
		k := st.Select(i)
		if st.Rank(k) != i {
			t.Errorf("rank of %q != %d", k, i)
		}
	}
}

func TestBSTDeleteMin(t *testing.T) {
	st := NewBST()
	nodes := []string{"S", "E", "X", "A", "R", "C", "H", "M"}
	for i, v := range nodes {
		st.Put(v, string(i))
	}

	st.DeleteMin()
	if st.Size() != 7 {
		t.Errorf("bst size should decrease by 1")
	}

	if st.Min() != "C" {
		t.Errorf("bst min key should be C")
	}	
}

func TestBSTDelete(t *testing.T) {
	st := NewBST()
	nodes := []string{"S", "E", "X", "A", "R", "C", "H", "M"}
	for i, v := range nodes {
		st.Put(v, string(i))
	}

	rnk := st.Rank("E")
	st.Delete("E")
	if st.Size() != 7 {
		t.Errorf("bst size should decrease by 1")
	}

	if st.Rank("H") != rnk {
		t.Errorf("bst node H should has rank=%d as E had before deletion", rnk)
	}
	
}