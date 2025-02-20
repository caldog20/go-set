package set

import (
	"testing"
)

func TestSetContains(t *testing.T) {
	s := From("hello", "set")
	if !s.Contains("set") {
		t.Fatal("expected set to contain 'set'")
	}
}

func TestSetUnion(t *testing.T) {
	s1 := From("hello")
	s2 := From("set")

	s3 := s1.Union(s2)
	if len(s3) != 2 {
		t.Fatalf("invalid set length: got %d - want: %d", len(s3), 2)
	}

	if !s3.Contains("hello") {
		t.Fatal("expected set to contain 'hello'")
	}

	if !s3.Contains("set") {
		t.Fatal("expected set to contain 'set'")
	}
}

func TestSetUnionChain(t *testing.T) {
	s1 := From("hello", "set")
	s2 := From("testing", "union", "set")
	s3 := From("hello", "sets")

	s4 := s1.Union(s2).Union(s3)

	if len(s4) != 5 {
		t.Fatalf("invalid set length: got %d - want: %d", len(s4), 5)
	}
}

func TestSetIntersect(t *testing.T) {
	s1 := From("hello", "set")
	s2 := From("bye", "set")

	s3 := s1.Intersect(s2)

	if len(s3) != 1 {
		t.Fatalf("invalid set length: got %d - want: %d", len(s3), 1)
	}

	if !s3.Contains("set") {
		t.Fatal("invalid set: expected set to contain 'hello'")
	}
}

func TestSetDifference(t *testing.T) {
	s1 := From("hello", "new", "set")
	s2 := New[string]()

	s3 := s2.Difference(s1)
	if len(s3) != 3 {
		t.Fatalf("invalid set length: got %d - want %d", len(s3), 3)
	}

	if !s3.Contains("hello") || !s3.Contains("new") || !s3.Contains("set") {
		t.Fatal("invalid set from difference")
	}
}

func TestSetSortedItems(t *testing.T) {
	s := New[byte]()
	s.Insert(4, 3, 2, 1)

	sorted := SortedItems(s)

	for i, v := range sorted {
		expected := i + 1
		if v != byte(expected) {
			t.Fatalf("invalid set order: got %d - want: %d", v, expected)
		}
	}
}

func TestSetRange(t *testing.T) {
	s := From("your", "mom", "is", "hot")
	count := 0
	expected := len(s)

	s.Range(func(s string) bool {
		count++
		return true
	})
	if count != expected {
		t.Fatalf("invalid interator count: got %d - want %d", count, expected)
	}
}

func TestSetIterator(t *testing.T) {
	s := From("your", "mom", "is", "hot")

	count := 0
	expected := len(s)

	for range s.Iter() {
		count++
	}

	if count != expected {
		t.Fatalf("invalid interator count: got %d - want %d", count, expected)
	}
}
