package main

import (
	"fmt"
)

type Trie interface {
	Get(key string) interface{}
	Put(key string, value interface{}) bool
	Insert(key string)
	Find(key string) bool
	Delete(key string) bool
	Keys(key string) interface{}
}

type araTrie struct {
	value    interface{}
	children map[rune]*araTrie
	isWord   bool
}

type nodeRune struct {
	node *araTrie
	r    rune
}

func InitTrie() *araTrie {
	return new(araTrie)
}

func (t *araTrie) Get(key string) interface{} {
	node := t
	for _, r := range key {
		node = node.children[r]
		if node == nil {
			return nil
		}
	}
	return node.value
}

func (t *araTrie) Put(key string, value interface{}) bool {
	node := t
	for _, r := range key {
		child, _ := node.children[r]
		if child == nil {
			if node.children == nil {
				node.children = map[rune]*araTrie{}
			}
			child = new(araTrie)
			node.children[r] = child
		}
		node = child
	}
	isNewVal := node.value == nil
	node.value = value
	node.isWord = true
	return isNewVal
}

func (t *araTrie) Insert(key string) {
	node := t
	for _, r := range key {
		child, _ := node.children[r]
		if child == nil {
			if node.children == nil {
				node.children = map[rune]*araTrie{}
			}

			child = new(araTrie)
			node.children[r] = child
		}
		node = child
	}
	node.isWord = true
}

func (t *araTrie) Find(key string) bool {

	node := t
	for _, r := range key {
		node = node.children[r]
		if node == nil {
			return false
		}
	}
	return true
}

func (t *araTrie) Delete(key string) bool {
	path := make([]nodeRune, len(key))
	node := t
	for i, r := range key {
		path[i] = nodeRune{r: r, node: node}
		node = node.children[r]
		if node == nil {
			return false // node doesn't exist
		}
	}
	node.value = nil

	if node.isLeaf() {
		for i := len(key) - 1; i > 0; i-- {
			parent := path[i].node
			r := path[i].r

			delete(parent.children, r)

			if !parent.isLeaf() {
				break
			}

			parent.children = nil

			if parent.value != nil {
				break
			}

		}
	}

	return true
}

func (t *araTrie) walk(ch *araTrie) {

	for r, child := range ch.children {

		if child.isWord {

			fmt.Println(" ", string(r), child.isWord)
		} else {

			fmt.Println(string(r), child.isWord)
		}

		ch.walk(child)
	}

}

func (t *araTrie) Keys(key string) {

	node := t
	for _, r := range key {
		node = node.children[r]
	}
	t.walk(node)
}

func (t *araTrie) isLeaf() bool {
	return len(t.children) == 0
}

func main() {
	tr := InitTrie()
	tr.Insert("hello")
	tr.Insert("help")

	tr.Insert("mohamed")
	tr.Insert("momen")
	tr.Insert("mohsen")

	fmt.Println(tr.Find("mohamed"))
	tr.Delete("mohamed")
	fmt.Println(tr.Find("mohamed"))

	tr.Insert("محمد")
	tr.Insert("محمود")
	tr.Insert("محمي")
}
