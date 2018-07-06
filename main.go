package main

import "fmt"

func main() {

	bar := Foo{1, "red"}
	fmt.Println(bar) // > "{1, "red"}

	DirectSetter(bar, "blue")
	fmt.Println(bar) // > "{1, "red"}

	pointer := &bar
	pointer.PointerSetter("yellow")
	fmt.Println(bar) // > "{1, "yellow"}

	bar.Y = "black"
	fmt.Println(bar) // > "{1, "black"}
}

type Foo struct {
	X int
	Y string
}

func (f *Foo) PointerSetter(y string) {
	f.Y = y
}

func DirectSetter(f Foo, y string) {
	f.Y = y
}
