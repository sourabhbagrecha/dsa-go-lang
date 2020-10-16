package hashtable

import (
	"fmt"
	"math"
)

// HashTable data structure
type HashTable struct {
	keyMap [50][][]string
}

//Hash the provided key
func (h *HashTable) _hash(key string) int {
	total := 0
	weirdPrime := 31
	runes := []rune(key)
	for i := 0; i < int(math.Min(float64(len(runes)), float64(100))); i++ {
		total = ((total * weirdPrime) + int(runes[i])) % len(h.keyMap)
	}
	return total
}

//Set the value corresponding to the key
func (h *HashTable) Set(key, value string) int {
	hash := h._hash(key)
	newArr := []string{key, value}
	h.keyMap[hash] = append(h.keyMap[hash], newArr)
	return hash
}

//Get the value corresponding to the key
func (h *HashTable) Get(key string) string {
	hash := h._hash(key)
	hashPointer := h.keyMap[hash]
	for i := 0; i < len(hashPointer); i++ {
		if hashPointer[i][0] == key {
			fmt.Println("The value for the key", key, "is", hashPointer[i][1])
			return hashPointer[i][1]
		}
	}
	return ""
}

//GetAllKeys returns the complete list of all the keys present in the Hash Table
func (h *HashTable) GetAllKeys() []string {
	allKeys := []string{}
	for i := 0; i < len(h.keyMap); i++ {
		if len(h.keyMap[i]) > 0 {
			for j := 0; j < len(h.keyMap[i]); j++ {
				allKeys = append(allKeys, h.keyMap[i][j][0])
			}
		}
	}
	return allKeys
}

//GetAllValues returns the complete list of all the values present in the Hash Table
func (h *HashTable) GetAllValues() []string {
	allValues := []string{}
	for i := 0; i < len(h.keyMap); i++ {
		if len(h.keyMap[i]) > 0 {
			for j := 0; j < len(h.keyMap[i]); j++ {
				allValues = append(allValues, h.keyMap[i][j][1])
			}
		}
	}
	return allValues
}
