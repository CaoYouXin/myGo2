package main

import (
	"bytes"
	"fmt"
)

// IntSet : a set for int
type IntSet struct {
	words []uint
}

// Size : word size
var Size = 32 << (^uint(0) >> 63)

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/Size, uint(x%Size)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/Size, uint(x%Size)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// String returns the set as a string of the form "{1 2 3}"
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')

	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < Size; j++ {
			if word&(1<<uint(j)) == 0 {
				continue
			}

			if buf.Len() > len("{") {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(&buf, "%d", i*Size+j)
		}
	}

	buf.WriteByte('}')
	return buf.String()
}

// Len returns the length of the bitset
func (s *IntSet) Len() (res int) {
	for _, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < Size; j++ {
			if word&(1<<uint(j)) != 0 {
				res++
			}
		}
	}
	return
}

// Remove changes the xth bit to 0
func (s *IntSet) Remove(x int) {
	word, bit := x/Size, uint(x%Size)
	if word >= len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)
}

// Clear clears all bit 1 to 0
func (s *IntSet) Clear() {
	s.words = make([]uint, 0)
}

// Copy returns a copy of the set
func (s *IntSet) Copy() (res *IntSet) {
	res = &IntSet{make([]uint, 0, len(s.words))}
	for _, word := range s.words {
		res.words = append(res.words, word)
	}
	return
}

// AddAll adds a bounch of items
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// Elems returns all bit 1 element
func (s *IntSet) Elems() []int {
	bits := make([]int, 0)

	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < Size; j++ {
			if word&(1<<uint(j)) != 0 {
				bits = append(bits, Size*i+j)
			}
		}
	}

	return bits
}

// UnionWith calculates the union with another set
func (s *IntSet) UnionWith(a *IntSet) {
	for _, elem := range a.Elems() {
		s.Add(elem)
	}
}

// IntersectWith calculates the intersection with another set
func (s *IntSet) IntersectWith(a *IntSet) {
	for _, elem := range s.Elems() {
		if !a.Has(elem) {
			s.Remove(elem)
		}
	}
}

// DifferenceWith calculates the difference with another set
func (s *IntSet) DifferenceWith(a *IntSet) {
	for _, elem := range a.Elems() {
		s.Remove(elem)
	}
}

// SymmetricDifference calculates symmetric difference with another set
func (s *IntSet) SymmetricDifference(a *IntSet) {
	acopy := a.Copy()
	acopy.DifferenceWith(s)
	s.DifferenceWith(a)
	s.UnionWith(acopy)
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())
	x.Clear()
	fmt.Println(&x)
	x.AddAll(32, 256, 128)
	fmt.Println(&x)

	y.Add(9)
	y.Add(42)
	fmt.Println(&y)
	fmt.Println(y.Len())
	y.Remove(9)
	fmt.Println(&y)

	z := y.Copy()
	fmt.Println(z.String())

	z.UnionWith(&x)
	fmt.Println(z)

	z.IntersectWith(&y)
	fmt.Println(z)

	var a, b IntSet
	a.AddAll(1, 2, 3)
	b.AddAll(1, 3, 4)
	a.SymmetricDifference(&b)
	fmt.Println()
	fmt.Println(&b)
	fmt.Println(&a)
}
