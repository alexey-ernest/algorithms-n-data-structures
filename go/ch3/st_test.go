package ch3

import (
	"testing"
	"bufio"
	"os"
	"strings"
)

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

func TestBSTKeys(t *testing.T) {
	st := NewBST()
	nodes := []string{"S", "E", "X", "A", "R", "C", "H", "M"}
	for i, v := range nodes {
		st.Put(v, string(i))
	}

	keys := st.Keys("H", "X")
	if len(keys) != 5 {
		t.Errorf("keys len should equal to 5")
	}

	exp := []string{"H", "M", "R", "S", "X"}
	for i := range keys {
		if exp[i] != keys[i] {
			t.Errorf("keys range validation failed")
		}
	}
}

func TestRedBlackEmpty(t *testing.T) {
	rb := NewRedBlackBST()
	if rb.Size() != 0 || !rb.IsEmpty() {
		t.Errorf("Red Black BST should be empty")
	}
}

func TestRedBlackBasic(t *testing.T) {
	st := NewRedBlackBST()
	for i := 0; i < 10; i+=1 {
		k := string('a' + i)
		st.Put(k, i)
	}

	if st.Size() != 10 {
		t.Errorf("size should equals 10, got %d", st.Size())
	}
	if st.IsEmpty() {
		t.Errorf("st should not be empty")	
	}

	for i := 0; i < 10; i+=1 {
		k := string('a' + i)
		if !st.Contains(k) {
			t.Errorf("st should contain the key %q", k)
		}
		if st.Get(k) != i {
			t.Errorf("value %s != %s", st.Get(k), i)
		}
	}
}

func TestRedBlackHeight(t *testing.T) {
	st := NewRedBlackBST()
	n := 1024
	for i := 0; i < n; i+=1 {
		k := string(i)
		st.Put(k, i)
	}

	if st.Size() != n {
		t.Errorf("size should equals %d, got %d", n, st.Size())
	}
	if st.IsEmpty() {
		t.Errorf("st should not be empty")	
	}

	height := st.Height()
	if height < 10 || height > 20 {
		t.Errorf("red black bst height should be in range lgN <= height <= 2*lgN, in our case from 10 to 20, but we got %d", height)
	}
}

func TestRedBlackLeipzig1M(t *testing.T) {

	file, err := os.Open("leipzig1M.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    st := NewRedBlackBST()
    n := 0
    for scanner.Scan() {
        line := scanner.Text()
        
        words := strings.Fields(line)
        n += len(words)
        
        for _, w := range words {
        	if !st.Contains(w) {
        		st.Put(w, 1)
        	} else {
        		st.Put(w, st.Get(w) + 1)
        	}
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    allwords := 622098
    if n != allwords {
    	t.Errorf("number of all words should be %d, got %d", allwords, n)
    }

    m := 69321
    if st.Size() != m {
		t.Errorf("size should equals %d, got %d", m, st.Size())
	}

	height := st.Height()
	if height < 16 || height > 32 {
		t.Errorf("red black bst height should be in range lgN <= height <= 2*lgN, in our case from 16 to 32, but we got %d", height)
	}

    if !st.IsRedBlack() {
    	t.Errorf("certification failed")
    }
}

func TestRedBlackMinMax(t *testing.T) {
	st := NewRedBlackBST()
	for i := 0; i < 10; i+=1 {
		k := string('a' + 9 - i)
		st.Put(k, i)
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

func TestRedBlackFloor(t *testing.T) {
	st := NewRedBlackBST()
	for i := 0; i < 10; i += 1 {
		k := string('a' + 20 - 2*i)
		st.Put(k, i)
	}

	keymiss := string('a' + 3)
	flmiss := string('a' + 2)
	if st.Floor(keymiss) != flmiss {
		t.Errorf("floor != %s", st.Floor(keymiss))
	}

	keyhit := string('a' + 10)
	if st.Floor(keyhit) != keyhit {
		t.Errorf("floor != %s", st.Floor(keyhit))
	}
}

func TestRedBlackCeiling(t *testing.T) {
	st := NewRedBlackBST()
	for i := 0; i < 10; i += 1 {
		k := string('a' + 20 - 2*i)
		st.Put(k, i)
	}

	keymiss := string('a' + 3)
	clmiss := string('a' + 4)
	if st.Ceiling(keymiss) != clmiss {
		t.Errorf("ceiling != %s", st.Ceiling(keymiss))
	}

	keyhit := string('a' + 10)
	if st.Ceiling(keyhit) != keyhit {
		t.Errorf("ceiling != %s", st.Ceiling(keyhit))
	}
}

func TestRedBlackSelect(t *testing.T) {
	st := NewRedBlackBST()
	for i := 0; i < 10; i+=1 {
		k := string('a' + 10 - i)
		st.Put(k, i)
	}

	key := string('a' + 3)
	if st.Select(2) != key {
		t.Errorf("element with rank=2 should be %s", key)
	}

	key = string('a' + 10)
	if st.Select(9) != key {
		t.Errorf("element with rank=9 should be %s", key)
	}
}

func TestRedBlackRank(t *testing.T) {
	st := NewRedBlackBST()
	nodes := []string{"S", "E", "X", "A", "R", "C", "H", "M"}
	for i, v := range nodes {
		st.Put(v, i)
	}

	for i := range nodes {
		k := st.Select(i)
		if st.Rank(k) != i {
			t.Errorf("rank of %q != %d", k, i)
		}
	}

	if st.Rank("Y") != len(nodes) {
		t.Errorf("rank of new maximum should equal to the number of nodes in the tree")
	}

	if st.Rank("Y") != st.Rank("Z") {
		t.Errorf("rank of new maximum should not depend on the new maximum concrete value")
	}
}

func TestRedBlackKeys(t *testing.T) {
	st := NewRedBlackBST()
	for i := 0; i < 10; i+=1 {
		k := string('a' + 10 - i)
		st.Put(k, i)
	}

	lo := string('a' + 3)
	hi := string('a' + 6)
	keys := st.Keys(lo, hi)
	if len(keys) != 4 {
		t.Errorf("keys len should equals 4, %+v", keys)
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

