package hashtables

import "testing"

// TestCreateHashTableEntry testing for hash table entry
func TestCreateHashTableEntry(t *testing.T) {
	got := NewHashTableEntry(1, 2)
	if got.key != 1 && got.key != 2 {
		t.Error("Error creating new hash table entry")
	}
}
