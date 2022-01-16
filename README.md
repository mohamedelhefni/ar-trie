# Ar Trie


### Implement Trie Datastructure that supports rune types with functionalities :
-  **InitTrie**  initialize the trie
-  **Insert**  Insert single key in your tree 
-  **Find** check if your key is in the tree
- **Put** map a key to a value 
- **Get** get value for your key 

### Trie Benchmark 

```
goos: linux
goarch: amd64
pkg: aratrie
cpu: AMD Ryzen 5 PRO 4650G with Radeon Graphics
BenchmarkInsert-12    	4684521	      235.4 ns/op	     63 B/op	      2 allocs/op
BenchmarkSearch-12    	8535787	      143.9 ns/op	      7 B/op	      0 allocs/op
BenchmarkPut-12       	4150489	      302.8 ns/op	     93 B/op	      3 allocs/op
BenchmarkGet-12       	8219924	      147.3 ns/op	      7 B/op	      0 allocs/op
PASS
ok  	aratrie	6.885s
```


