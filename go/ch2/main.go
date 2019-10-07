package ch2

type Sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Compare(i, j int) int
}

type Sorter interface {
	Sort(data Sortable)
}


type SortableInt []int

func (s SortableInt) Len() int {
	return len(s)
}

func (s SortableInt) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortableInt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortableInt) Compare(i, j int) int {
	if s[i] < s[j] {
		return -1
	} else if s[i] > s[j] {
		return 1
	} else {
		return 0
	}
}

type SortableByte []byte

func (s SortableByte) Len() int {
	return len(s)
}

func (s SortableByte) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortableByte) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortableByte) Compare(i, j int) int {
	if s[i] < s[j] {
		return -1
	} else if s[i] > s[j] {
		return 1
	} else {
		return 0
	}
}
