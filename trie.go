package main

type Trie interface {
	Get(key string) interface{}
	Put(key string, value interface{}) bool
	Insert(key string)
	Find(key string) bool
	Search(key string) []interface{}
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

var wordList = make([]string, 5)

func dfs(ch *araTrie, word string) {
	for r, child := range ch.children {
		if child.isWord {
			wordList = append(wordList, word+string(r))
		}
		dfs(child, word+string(r))
	}

}

func (t *araTrie) Keys(key string) []string {
	wordList = nil
	node := t
	for _, r := range key {
		node = node.children[r]
	}

  if node.isWord {
    wordList = append(wordList, key)
  }
	if node != nil {
		dfs(node, key)
	}
	return wordList
}

func (t *araTrie) Search(key string) []interface{} {
	results := make([]interface{}, 1)
	keys := t.Keys(key)
	if len(keys) > 0 {
		for _, key := range keys {
			result := t.Get(key)
			if len(results) > 0 {
				results = append(results, result)
			}
		}
	}
	return results[1:]
}

func main() {
	tr := InitTrie()

	tr.Insert("mohamed")
	tr.Insert("mohmed")
	tr.Insert("modaser")
	tr.Insert("monzer")
	tr.Insert("momen")
	tr.Insert("mohsen")
  tr.Keys("mo")

	tr.Insert("محمد")
	tr.Insert("محمود")
	tr.Insert("محمي")
	tr.Keys("مح")

	tr.Put("hello", "world")           
	tr.Put("here", "is a trie search") 
  tr.Search("he")  // ["world", "is a trie search]
}
