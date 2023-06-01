package wrap

import (
	"fmt"
	"go-method-hiding/upstream"
)

type privateSayIt = upstream.SayIt

type Fooer struct {
	*privateSayIt
}

func (_ *Fooer) Foo() {
	fmt.Println("our foo!")
}

type private struct{}

func (_ *Fooer) Bar(_ private) {}
