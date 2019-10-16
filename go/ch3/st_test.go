package ch3

import "testing"

func TestSequentialSearchSTBasic(t *testing.T) {
	st := SequentialSearchST{}
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