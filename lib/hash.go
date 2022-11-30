package lib

import (
	"fmt"
	"strings"
)

type HashData[T any] struct {
	data map[int]T
}

// Hash private function will do some math operation and convert
// the key into byte
func (h *HashData[_]) hash(k string) byte {
	key := fmt.Sprintf("%d", k)

	hash := 0
	for i := 0; i < len(key); i++ {
		strSplit := strings.Split(k, "")
		strByte := []byte(strSplit[0])
		hash = int(byte(hash)) + int(key[i]) + int(strByte[0])
	}

	return byte(hash)

}

// Put will call Hash() method, returning byte and set it as
// array indices and set its value
func (h *HashData[T]) Put(key string, value T) *HashData[T] {
	index := h.hash(key)

	if h.data == nil {
		h.data = make(map[int]T)
	}

	h.data[int(index)] = value

	return h
}

// Get will hash inputted key and get the map of index
func (h *HashData[T]) Get(key string) T {
	index := h.hash(key)
	return h.data[int(index)]
}

// Remove will delete specific element in key
// It will call delete() built-in function
func (h *HashData[_]) Remove(key string) {
	index := h.hash(key)
	delete(h.data, int(index))
}
