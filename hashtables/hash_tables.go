package hashtables

// HashTableEntry it's created on insertion of hash table
type HashTableEntry struct {
	key   int
	value int
}

// NewHashTableEntry creates new hash table entry
func NewHashTableEntry(key, value int) *HashTableEntry {
	return &HashTableEntry{key, value}
}
