package bst

import (
	"errors"
	"fmt"
)

var root int

type bstNode struct {
	val   int
	left  *bstNode
	right *bstNode
	level int
}

func Root(val int) *bstNode {
	root = val
	return &bstNode{
		val:   val,
		level: 0,
	}
}

func (t *bstNode) Value() int {
	if t == nil {
		return 0
	}
	return t.val
}

func (t *bstNode) Left() *bstNode {
	if t == nil {
		return nil
	}
	return t.left
}
func (t *bstNode) Right() *bstNode {
	if t == nil {
		return nil
	}
	return t.right
}

func (t *bstNode) PrintInorder() {
	if t == nil {
		return
	}
	t.left.PrintInorder()
	fmt.Print(t.val, " ")
	t.right.PrintInorder()
}

var pos int

func (t *bstNode) PrintTree() {
	matrix := [6][41]int{}
	pos = 20
	matrix[0][pos] = t.val
	t.printTreeCall(&matrix)
	for _, elem := range matrix {
		for _, val := range elem {
			if val == 0 {
				fmt.Print(" ")
				continue
			}
			fmt.Print(val)
		}
		fmt.Print("\n")

	}
}

func (t *bstNode) printTreeCall(m *[6][41]int) {
	if t == nil {
		return
	}
	if t.left != nil {
		if t.level == 0 {
			pos = 12
		} else {
			pos -= 3
		}

		m[t.left.level][pos] = t.left.val
		t.left.printTreeCall(m)
	}
	if t.right != nil {
		if t.level == 0 {
			pos = 28
		} else {
			pos += 3
		}
		m[t.right.level][pos] = t.right.val
		t.right.printTreeCall(m)
	}
}

func (t *bstNode) Insert(value int) error {

	if t == nil {

		return errors.New("bst is nil")
	}
	if t.val == value {
		return errors.New("this node value already exists")
	}

	if value < t.val {

		if t.left == nil {
			t.left = &bstNode{
				val:   value,
				level: t.level + 1,
			}
			return nil
		}

		return t.left.Insert(value)
	}

	if value > t.val {

		if t.right == nil {
			t.right = &bstNode{
				val:   value,
				level: t.level + 1,
			}

			return nil
		}

		return t.right.Insert(value)
	}

	return nil
}

func (t *bstNode) Find(value int) (bstNode, bool) {

	if t == nil {
		return bstNode{}, false
	}

	switch {
	case value == t.val:
		return *t, true
	case value < t.val:
		return t.left.Find(value)
	default:
		return t.right.Find(value)
	}
}

func (t *bstNode) Delete(value int) {
	t.remove(value)
}

func (t *bstNode) remove(value int) *bstNode {

	if t == nil {
		return nil
	}

	if value < t.val {
		t.left = t.left.remove(value)
		return t
	}
	if value > t.val {
		t.right = t.right.remove(value)
		return t
	}

	if t.left == nil && t.right == nil {
		t = nil
		return nil
	}

	if t.left == nil {
		t = t.right
		return t
	}
	if t.right == nil {
		t = t.left
		return t
	}

	smallestValOnRight := t.right
	for {
		if smallestValOnRight != nil && smallestValOnRight.left != nil {
			smallestValOnRight = smallestValOnRight.left
		} else {
			break
		}
	}

	t.val = smallestValOnRight.val
	t.right = t.right.remove(t.val)
	return t
}

func (t *bstNode) FindMax() int {
	if t.right == nil {
		return t.val
	}
	return t.right.FindMax()
}

func (t *bstNode) FindMin() int {
	if t.left == nil {
		return t.val
	}
	return t.left.FindMin()
}
