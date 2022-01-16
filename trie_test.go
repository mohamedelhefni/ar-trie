package main

import (
	"strconv"
	"testing"
)

var t = InitTrie()
var tr = InitTrie()

func BenchmarkInsert(b *testing.B){

  for i := 0 ; i < b.N ; i++ {
    t.Insert("رقم" + strconv.Itoa(i))
  }
}


func BenchmarkSearch( b *testing.B){
  for i := 0 ; i < b.N; i++ {
    t.Find("رقم" + strconv.Itoa(i))
  }
}


func BenchmarkPut(b *testing.B){
  for i := 0 ; i < b.N; i++ {
    val := "رقم" + strconv.Itoa(i)
    tr.Put(val, val)
  }
}


func BenchmarkGet(b *testing.B){
  for i := 0 ; i < b.N; i++ {
    val := "رقم" + strconv.Itoa(i)
    tr.Get(val)
  }
}
