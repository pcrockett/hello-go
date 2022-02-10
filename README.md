## Hello World

My first real foray into Go.

### Getting Started

Install Go. Then run `./tudas`.

### TODO

* [x] Use `log` package from https://go.dev/doc/tutorial/handle-errors
* [x] Refactor greetings package to return a greeting struct with both hello and goodbye
* [x] Public / private methods
* [x] Use interfaces for polymorphism
* [x] Define custom greeting
* [x] Parse command line parameters
* [ ] Testing
* [ ] Play with some crypto!
* [ ] Play with Gin

## Misc Notes

Surprised at how Go uses pointers similar to C / C++, and expects the programmer to know the difference between passing a value and passing a pointer to a method.

For example, a method that has a value receiver cannot modify the original object, because all it has is a _copy_ of the original object. If you need to modify the original object, you need to have a pointer receiver.

