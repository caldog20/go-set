# go-set
A set type to use in my Go projects. The set is not threadsafe so synchronization is requred if reading/writing to the set from multiple goroutines.

Usage:
```go
// Create a new empty set of type `string`
s1 := set.New[string]()

// Insert a single item to the set:
s1.Insert("hello")

// Insert multiple items to a set:
s1.Insert("this", "is", "a", "set")

// or from a slice:
s1.Insert([]string{"this", "is", "a", "set"}...)


//A set can also be created from multiple items:
s2 := set.From("hello", "world")

// Remove an item from the set - returns true if the item existed:
removed := s2.Remove("world")

// Create a union of two sets as a new set:
union := s1.Union(s2)

// Get the difference of two sets as a new set:
diff := s1.Difference(s2)

// Get the intersection of the two sets as a new set:
common := s1.Intersect(s2)

// Check if the set contains an item:
contains := s1.Contains("hello")

// Get a slice of all of the items of the set
items := s1.ToSlice()

// Range over all items in the set:
s1.Range(func (s string) bool {
    if s == "hello" {
        return false
    }
    fmt.Println(s)
    return true
})

// Get an iter.Seq[T] of the items in the set:
iter := s.Iter()

// Get the current size (cardinality) of the set:
s1.Size()

// Clear all items in the set:
s1.Clear()


// If type T of Set[T] is orderable, you can get a slice of sorted items from a set:
intSet := set.From(5, 4, 3, 2, 1)
sorted := set.SortedItems(intSet)
// sorted: [1, 2, 3, 4, 5]
```


