# go-method-hiding

Some time ago I was trying to sort out how to 'hide' a method on an embedded value in go, in order
to encapsulate some behavior it exposes.

Previously, I found the first part of the solution: hiding the embedding. By nature of the embedding
syntax, if you embed an exported type, the embedding will also be externally (to the defining
package) visible. This can be cirumvented with an unexported type alias, which I learned about via
some Google search that landed me on StackOverflow or a mailing list archive that I can no longer
find.

Unfortunately for my use-case, hiding the embedding still leaves the embedded type's methods
available, and none of my Google searches found anything. I could redefine the method and just make
it panic, but I don't really want to push explosions to runtime.

Eventually, I realized: I can hide the embedding, and then make calls to a method uncompilable
(outside my package) by defining it to receive an argument that can't be created outside my package.
Thus calls to the method I want to hide on the wrapping struct will not compile, and the method is
finally hidden.

## structure

* `upstream/` represents a package from some upstream that exports a type we're going to wrap, which
    exports a method we're going to hide.
* `wrap/` is where we wrap the type from upstream and hide one of its methods. That method should be
    uncallable outside of the wrap package.
* `main.go` demonstrates usage of `wrap`.

## what's the use-case?

sadly I do not recall.
