package tree

import "fmt"

type Tree struct {
    value       int
    left, right *Tree
}

func Add(t *Tree, value int) *Tree {
    if t == nil {
        // Equivalent to return &Tree{value: value}.
        t = new(Tree)
        t.value = value
        return t
    }
    if value < t.value {
        t.left = Add(t.left, value)
    } else {
        t.right = Add(t.right, value)
    }
    return t
}

func InOrder(t *Tree) {
	if (t != nil) {
		InOrder(t.left)
		fmt.Println(t.value)
		InOrder(t.right)
	}
}