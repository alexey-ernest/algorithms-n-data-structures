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
