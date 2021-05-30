// Package tree reconstructs tree layouts from an unsorted set of records.
package tree

// Import packages.
import (
	"errors"
	"fmt"
)

// Record type stores an ID and a parent ID.
type Record struct {
	ID, Parent int
}

// Node type stores an ID and the IDs of all children.
type Node struct {
	ID       int
	Children []*Node
}

// Mismatch type is used to throw errors when needed.
func (r *Record) String() string {
	return fmt.Sprintf("(ID: %d, Parent: %d)", r.ID, r.Parent)
}

// Error generates errors in case of mismatch.

// checkAndSort orders the record set and orders them by ID.
func checkAndSort(records []Record) ([]Record, error) {
	n := len(records)
	for key, item := range records {
		if item.ID >= n {
			return nil, errors.New("Error: Invalid ID value, there is a risk of non-continuity or duplicate node!")
		}
		for i := key + 1; i < n; i++ {
			if records[i].ID == records[key].ID {
				return nil, errors.New("Error: Duplicate node detected!")
			}
		}
		if item.ID != 0 && item.Parent >= item.ID {
			return nil, errors.New("Error: Parent ID cannot exceed one of its children's ID!")
		}
		if item.ID == 0 && item.Parent != 0 {
			return nil, errors.New("Error: Node 0 cannot have a parent!")
		}
	}
	sortedRecords := make([]Record, n)
	for _, item := range records {
		sortedRecords[item.ID] = item
	}
	return sortedRecords, nil
}

// Build generates a tree based on an unsorted set of records.
func Build(records []Record) (*Node, error) {
	if len(records) <= 0 {
		return nil, nil
	}
	sortedRecords, err := checkAndSort(records)
	if err != nil {
		return nil, err
	}
	tree := make([]Node, len(records))
	for _, item := range sortedRecords {
		tree[item.ID].ID = item.ID
		if item.ID != 0 {
			tree[item.Parent].Children = append(tree[item.Parent].Children, &tree[item.ID])
		}
	}
	return &tree[0], nil
}
