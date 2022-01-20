<<<<<<< HEAD
# ar-trie
=======
=======
>>>>>>> dddb29e878527525d344ac06ed8e6c1c01db7701
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
<<<<<<< HEAD
BenchmarkInsert-12    	4616820	      238.6 ns/op	     63 B/op	      2 allocs/op
BenchmarkSearch-12    	8523724	      140.3 ns/op	      7 B/op	      0 allocs/op
BenchmarkPut-12       	4158210	      303.4 ns/op	     93 B/op	      3 allocs/op
BenchmarkGet-12       	7823724	      152.0 ns/op	      7 B/op	      0 allocs/op
BenchmarkKeys-12      	1000000000	        0.5880 ns/op	      0 B/op	      0 allocs/op
PASS
ok  	aratrie	28.368s
```

```
goos: linux
goarch: amd64
pkg: aratrie
cpu: Intel(R) Xeon(R) CPU @ 2.20GHz
BenchmarkInsert-2        2882205               364.1 ns/op            34 B/op          1 allocs/op
BenchmarkSearch-2        4970304               254.2 ns/op             7 B/op          0 allocs/op
BenchmarkPut-2           2340908               504.7 ns/op            80 B/op          3 allocs/op
BenchmarkGet-2           5024829               254.0 ns/op             7 B/op          0 allocs/op
PASS
ok      aratrie 11.603s

```

## Examples 
```go 
func main() {

	tr := InitTrie()
	tr.Insert("mohamed")
	tr.Insert("mohmed")
	tr.Insert("modaser")
	tr.Insert("monzer")
	tr.Insert("momen")
	tr.Insert("mohsen")
	fmt.Println(tr.Keys("mo")) // [mohamed  mohmed mohsen modaser monzer momen]

	tr.Insert("محمد")
	tr.Insert("محمود")
	tr.Insert("محمي")
	fmt.Println(tr.Keys("مح")) // [محمود محمي محمد]
}

```
=======
BenchmarkInsert-12    	4684521	      235.4 ns/op	     63 B/op	      2 allocs/op
BenchmarkSearch-12    	8535787	      143.9 ns/op	      7 B/op	      0 allocs/op
BenchmarkPut-12       	4150489	      302.8 ns/op	     93 B/op	      3 allocs/op
BenchmarkGet-12       	8219924	      147.3 ns/op	      7 B/op	      0 allocs/op
PASS
ok  	aratrie	6.885s
```


>>>>>>> dddb29e878527525d344ac06ed8e6c1c01db7701
