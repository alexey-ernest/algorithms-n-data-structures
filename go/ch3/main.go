package ch3

type SymbolTable interface {
	Put(key, value string)
	Get(key string) string
	Delete(key string)
	Size() int
	IsEmpty() bool
	Contains(key string) bool
}
