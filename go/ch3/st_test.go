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