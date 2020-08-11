package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

// Build builds a tree of nodes from given records.
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	n := make([]*Node, len(records))
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})
	for i, r := range records {
		n[i] = &Node{ID: r.ID}
		if i == 0 && (r.ID != 0 || r.Parent != 0) {
			return nil, errors.New("invalid root record")
		} else if i == 0 {
			continue
		} else if i != r.ID || r.ID <= r.Parent {
			return nil, errors.New("invalid record")
		}

		if n[r.Parent] != nil {
			n[r.Parent].Children = append(n[r.Parent].Children, n[i])
		}
	}

	return n[0], nil
}
