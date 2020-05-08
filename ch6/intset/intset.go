package intset

import (
	"bytes"
	"fmt"
)

var intSize int

func init() {
	if 32<<(^uint(0)>>63) == 0 {
		intSize = 32
		return
	}
	intSize = 64
}

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
// The set contains i if the i-th bit is set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/intSize, x%intSize
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set
func (s *IntSet) Add(x int) {
	word, bit := x/intSize, uint(x%intSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll allows a list of values to be added
func (s *IntSet) AddAll(x ...int) {
	for _, val := range x {
		s.Add(val)
	}
}

// Remove removes x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/intSize, uint(x%intSize)
	if word < len(s.words) {
		s.words[word] &^= (1 << bit)
	}
}

// UnionWith sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t
func (s *IntSet) IntersectWith(t *IntSet) {
	var counter int
	for _, tword := range t.words {
		if counter < len(s.words) {
			s.words[counter] &= tword
		}
		counter++
	}
	for counter < len(s.words) {
		s.words[counter] = 0
		counter++
	}
}

// DifferenceWith sets s to the difference between s and t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

// SymmetricDifference sets s to the symmetric difference between s and t.
// (elements not shared by s or t)
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Len returns the number of elements
func (s *IntSet) Len() int {
	var n int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				n++
			}
		}
	}
	return n
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy returns a copy of s
func (s *IntSet) Copy() *IntSet {
	var x IntSet
	x.words = append(x.words, s.words...)
	return &x
}

// Elems returns a slice containing the elements of the set
func (s *IntSet) Elems() []uint {
	return s.words
}

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < intSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", intSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
