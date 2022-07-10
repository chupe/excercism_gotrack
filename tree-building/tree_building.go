package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	length := len(records)
	if length == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	var nodes = make([]*Node, length)
	for _, r := range records {
		if r.ID < 0 || r.ID >= length {
			return nil, errors.New("id out of bounds")
		}
		if r.Parent == r.ID && r.ID != 0 {
			return nil, errors.New("r.id > parent id")
		}
		if r.Parent > r.ID {
			return nil, errors.New("root has parent")
		}
		if nodes[r.ID] != nil {
			return nil, errors.New("r.id already inserted")
		}
		nodes[r.ID] = &Node{ID: r.ID}
		if r.ID == 0 {
			continue
		}
		parent := nodes[r.Parent]
		parent.Children = append(parent.Children, nodes[r.ID])
	}

	return nodes[0], nil
}
