package main

import (
	"math/rand"
	"testing"
)

var stringKeys [1000]string // random string keys
const bytesPerKey = 30

func init() {
	// string keys
	for i := 0; i < len(stringKeys); i++ {
		key := make([]byte, bytesPerKey)
		if _, err := rand.Read(key); err != nil {
			panic("error generating random byte slice")
		}
		stringKeys[i] = string(key)
	}
}

func BenchmarkInsert(b *testing.B) {

	trie := InitTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Insert(stringKeys[i%len(stringKeys)])
	}
}

func BenchmarkFind(b *testing.B) {
	trie := InitTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Find(stringKeys[i%len(stringKeys)])
	}
}

func BenchmarkPut(b *testing.B) {
	trie := InitTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Put(stringKeys[i%len(stringKeys)], i)
	}
}

func BenchmarkGet(b *testing.B) {
	trie := InitTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Get(stringKeys[i%len(stringKeys)])
	}
}

func BenchmarkKeys(b *testing.B) {
	trie := InitTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Keys(stringKeys[i%len(stringKeys)][:3])
	}
}

func BenchmarkSearch(b *testing.B) {
	trie := InitTrie()
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		trie.Search(stringKeys[i%len(stringKeys)][:3])
	}
}
