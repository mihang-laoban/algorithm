package hash

func IsAnagram2(s, t string) bool {
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}

type MyHashSet struct {
	bitset []uint64
}

/** Initialize your data structure here. */
func HashSetConstructor() MyHashSet {
	return MyHashSet{bitset: []uint64{}}
}

func (s *MyHashSet) Add(key int) {
	bit := key % 64
	length := key / 64
	for i := len(s.bitset); i <= length; i++ {
		s.bitset = append(s.bitset, 0)
	}
	s.bitset[length] = s.bitset[length] | (1 << uint(bit))
}

func (s *MyHashSet) Remove(key int) {
	bit := key % 64
	length := key / 64
	if length >= len(s.bitset) {
		return
	}
	s.bitset[length] = s.bitset[length] & ^(1 << uint(bit))
}

/** Returns true if this set contains the specified element */
func (s *MyHashSet) Contains(key int) bool {
	bit := key % 64
	length := key / 64
	if length >= len(s.bitset) {
		return false
	}
	return s.bitset[length]&(1<<uint(bit)) != 0
}
