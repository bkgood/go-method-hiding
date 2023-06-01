package upstream

import "fmt"

type SayIt struct{}

func (_ *SayIt) Foo() {
	fmt.Println("foo")
}

func (_ *SayIt) Baz() {
	fmt.Println("baz")
}

func (_ *SayIt) Bar() {
	fmt.Println("bar")
}
