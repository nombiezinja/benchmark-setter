package main

import "testing"

var bar = Foo{1, "red"}

//BenchmarkPointerSetter-8   	2000000000	         0.82 ns/op	       0 B/op	       0 allocs/op
//BenchmarkPointerSetter-8   	2000000000	         0.79 ns/op	       0 B/op	       0 allocs/op
//BenchmarkPointerSetter-8   	2000000000	         0.84 ns/op	       0 B/op	       0 allocs/op
func BenchmarkPointerSetter(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pointer := &bar
		pointer.PointerSetter("yello")
	}
}

//BenchmarkNoFuncSet-8   	2000000000	         0.78 ns/op	       0 B/op	       0 allocs/op
//BenchmarkNoFuncSet-8   	2000000000	         0.79 ns/op	       0 B/op	       0 allocs/op
//BenchmarkNoFuncSet-8   	2000000000	         0.85 ns/op	       0 B/op	       0 allocs/op
func BenchmarkNoFuncSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bar.Y = "yellow"
	}
}
