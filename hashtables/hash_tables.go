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

// HashTable structure
type HashTable struct {
	t         [][]HashTableEntry
	TableSize int
}

// NewHashTable creates new hash table
func NewHashTable(tableSize int) *HashTable {
	return &HashTable{make([][]HashTableEntry, tableSize), tableSize}
}

// HashFunc applies hash function
func (ht *HashTable) HashFunc(key int) int {
	return key % ht.TableSize
}

// Insert inserts into hash table
func (ht *HashTable) Insert(key, value int) {
	hash := ht.HashFunc(key)
	if len(ht.t[hash]) > 0 {
		ht.t[hash] = append(ht.t[hash], *NewHashTableEntry(key, value))
	} else {
		ht.t[hash] = make([]HashTableEntry, 1)
		ht.t[hash][0] = *NewHashTableEntry(key, value)
	}
}

// Search searches in hash table
func (ht *HashTable) Search(key int) []HashTableEntry {
	hash := ht.HashFunc(key)
	return ht.t[hash]
}

// Remove key
func (ht *HashTable) Remove(key int) {
	hash := ht.HashFunc(key)
	ht.t[hash] = []HashTableEntry{}
}
