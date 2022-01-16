package main

import (
	"fmt"
)

type Trie interface {
	Get(key string) interface{}
	Put(key string, value interface{}) bool
	Insert(key string)
	Find(key string) bool
	Keys(key string) interface{}
	walk(key string)
}

type araTrie struct {
	value    interface{}
	children map[rune]*araTrie
	isWord   bool
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


func (t *araTrie) walk(ch *araTrie) {

	for r, child := range ch.children {

		if child.isWord {

			fmt.Println(" ", string(r), child.isWord)
		} else {

			fmt.Println(string(r), child.isWord)
		}

		ch.walk(child)
	}

	//for r, child := range t.children {

	//fmt.Println(child.children)
	//child.walk(key + string(r))

	//}

}

//keys := make([]interface{}, 5)
func (t *araTrie) Keys(key string) {
	
	node := t
	for _, r := range key {
		node = node.children[r]
	}
	t.walk(node)
}

func main() {
	tr := InitTrie()
	tr.Insert("hello")
	tr.Insert("help")

	tr.Insert("mohamed")
	tr.Insert("momen")
	tr.Insert("mohsen")
  tr.Keys("mo")

  // bench 
/*  for i := 0 ; i < 10_000_000; i++ {*/
    /*tr.Insert("رقم" + strconv.Itoa(i))*/
  /*}*/

 
  tr.Insert("محمد")
  tr.Insert("محمود")
  tr.Insert("محمي")
  tr.Keys("مح")
  tr.walk(tr)
}







