package main

import (
	"strconv"
	"testing"
)


func BenchmarkInsert(b *testing.B){
  tr := InitTrie()
  for i := 0 ; i < b.N ; i++ {
    tr.Insert("رقم" + strconv.Itoa(i))
  }
}


