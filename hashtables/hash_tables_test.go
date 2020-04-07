package hashtables

import "testing"

func TestNewHashTableEntry(t *testing.T) {
	got := NewHashTableEntry(1, 2)
	if got.key != 1 && got.key != 2 {
		t.Error("Error creating new hash table entry")
	}
}

func TestNewHashTable(t *testing.T) {
	ht := NewHashTable(3)
	if len(ht.t) != 3 && ht.TableSize != 3 {
		t.Error("Hash not created")
	}
}

func TestInsert(t *testing.T) {
	ht := NewHashTable(3)
	ht.Insert(1, 2)
	ht.Insert(0, 1)
	if ht.t[1][0].value != 2 && ht.t[0][1].value != 1 {
		t.Error("no insertion")
	}
}
