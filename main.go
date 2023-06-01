package main

import "go-method-hiding/wrap"

func main() {
	x := &wrap.Fooer{}

	x.Foo() // can call a method we overrode, of course

	x.Baz() // can still call a method on the embedded struct

	// will fail to compile: not enough arguments in call to x.Bar
	//x.Bar()

	// fails to compile: private not exported by package wrap
	//x.Bar(wrap.private{})
}
