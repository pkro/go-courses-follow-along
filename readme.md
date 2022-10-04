Golang stuff from various videos on youtube

# Notes from "Learn Go Programming - Golang Tutorial for Beginners"

https://www.youtube.com/watch?v=YS4e4q9oBaU

## Introduction

### Overview / motivation

- Strongly (type can't be changed after declaration) and statically (types have to be defined at compile time) typed
- Strong community
- key features:
    - simplicity (which also means less features that would add complexity)
    - fast compile times
    - fast garbage collection
    - built-in concurrency
    - compiles to standalone libraries (no external runtime or libraries needed)
- focused on server applications (no UI libraries)

### Ressources

- https://go.dev/ including a little in-built [IDE](https://go.dev/play/)
- Documentation: https://go.dev/doc/
- https://go.dev/doc/effective_go required reading for serious development
- https://pkg.go.dev/std Package documentation
- [go forum](https://forum.golangbridge.org/categories)
- https://gobyexample.com/


### Packages

Every go program ist structured into packages. Every file needs to be part of a package. The package `main` is the entrypoint of the application.

`package main`

### Imports

Syntax:

`import(
"fmt"
"anotherPackage"
)`

or optionally `import "fmt"` if only one library is used.

All imports have to be used somewhere in order for the program to be compiled.


## Syntax

- no semicolons, thus no multiple statements in a line.
- comments: the usual `//` and `/* blah */`

## Installation / IDE

Install and put bin path in `~/.bashrc`: `export PATH=$PATH:/usr/local/go/bin` (default go path on linux).

Don't define `GOROOT`, as go will figure it out from its own location, and it can lead to problems with multiple go versions.

`GOPATH` is a monolithic directory where packages and binaries are stored, similar to a global `node_modules`. Packages go to `~/go/pkg`, binaries to `~/go/bin`.

If no GOPATH is defined, `~/go` is assumed, otherwise define GOPATH in `.bashrc` as well.

Multiple GOPATH paths can be defined, e.g. `/home/pk/go:/home/pk/projects/go/helloworld`

The *first* path is where libraries are installed, but all paths will be *searched* when looking for libraries.

I guess the IDE should take care of this for projects I guess.

https://www.digitalocean.com/community/tutorials/understanding-the-gopath

*Deprecated*: Libraries can be installed (globally) using `go get`, e.g. `go get github.com/nsf/gocode`

'go get' is no longer supported outside a module.

**Update**: "new" way to install global libs / binaries: `go install  github.com/nsf/gocode@latest`

A go workspace is any folder with a `src` subfolder. Other (generated) ones are `bin` and `pkg` just as in the monolithic `~/go`path.

## Variables

### Declaration

Variable names must start with a letter or underscore, including unicode.

**Variables starting with an uppercase letter are exported, lowercase variables are scoped to the package**

Declaration syntax:

1) `var varname type`, e.g. `var speed int`; declaration without initialization
2) `var varname type = value`, e.g. `var x float32 = 10.8` declaration WITH initialization; use when there's not enough information for the compiler to figure out the type from the value
3) `varname := value`, e.g. `x := 10` (the compiler figures the type out from the value)
- values have to be unambiguous to be inferred correctly; when adding a decimal sign behind a whole number, `float64` is automatically inferred

Variables are initialized to `0`.

    package main
    
    import (
        "fmt"
    )
    func main() {
        //var i int
        // i = 42
        //var j float32 = 27
        k := 99.
        fmt.Printf("%v, %T", k, k) // %v = value, %T = type
        // output: 99, float64
    }

Variables can be declared inside a function or at *package level*. At package level, types can't be inferred so `:=` can't be used there.

When declaring several variables, the declaration can be put in a `var` block so individual `var` keywords are unnecessary and indicate a relation of the variables to the reader:

    var actorName string = "joe"
    var companion string = "Sara"
    var doctorNumber int = 3
    var season int = 11

becomes

    var (
        actorName    string = "joe"
        companion    string = "Sara"
        doctorNumber int    = 3
        season       int    = 11
    )

### Redeclaration and shadowing

Variables declared in an outer scope can be redeclared (shadowed) by a new variable of the same name in an inner scope:

    var x int = 27
    
    func main() {
        fmt.Printf("%v", x) // 27
        var x int = 42
        fmt.Printf("%v", x) // 42
    }

**All <u>locally</u> declared variables have to be used somewhere for the program to compile!**

### Variable scope / visibility

- lowercase variables at the package level are visible anywhere in the package
- Uppercase variables at the package level are exported from the package and are globally visible
- locally scoped variables are visible in their respective scope and enclosed scopes
- there is no private scope for the current file for package level variables (only package scope, which can include several source files which all can access the package level variables)

### Naming conventions

- the length of the variable name should reflect the lifespan of the variable (e.g. `i` for a loop variable)
- others should be more verbose / clear
- acronyms should be all uppercase, e.g. `theURL` instead of `theUrl`

### Casting

Variable can be cast / converted to another type by using the type like a function, e.g.

    var i int = 42
	var j float32
	j = float32(i)
	fmt.Printf("%v %T", j, j) // 42 float32

Converting a higher precision type (e.g. `float32`) to one with lower precision loses information, e.g. `42.8` cast to int becomes `42`.

When trying to convert a number to a string, it converts it to a unicode character of that ord number, e.g. `string(42)` becomes the character `*`.

To convert variables of various types to strings properly, the package [strconv](https://pkg.go.dev/strconv) has to be used.

    var x int = 27  
	var j = strconv.Itoa(x)
	fmt.Printf("%v %T", j, j) // 42 float32

## Primitive datatypes

All primitives are initialized to a zero value if no value is given on initialization.

- boolean (defaults to `false` if uninitialized)

### numeric

#### Integer types

- int - integer of system dependent, unspecified size, at least 32 bits; initialized with `0`
- uint - same but unsigned - only positive values
    - int8 / uint8 / byte
    - int16 / uint16
    - int32 / uint32
    - int64 - no unsigned version

`int` types are not automatically converted between each other but must be cast accordingly:

    var a int = 10
    var b int8 = 3
    int c = a + (int(b))

Operators that can be used with integers:

	a := 10
	b := 3
	
	// arithmetic operators
	fmt.Println(a + b) // 13
	fmt.Println(a - b) // 10
	fmt.Println(a / b) // 3 - int division result without remainder
	fmt.Println(a % b) // 1 - remainder of the int division
	
	// logical operators
	// are applied between each bit of the values, e.g.
	// a & b:
	// 10 = 1010
	// 3  = 0011
	// ---------
	// 2  = 0010
	
	fmt.Println(a & b) // 2 - and
	fmt.Println(a | b) // 11 - or
	fmt.Println(a ^ b) // 9 - xor
	fmt.Println(a &^ b) // 8 - and not - only the first one has the bit set, but not the other 
	fmt.Println(b &^ a) // 1 - so a &^ b has a different result than b &^ a, unlike the previous logical operators
	
	// bit shifting operators
	a = 8
	fmt.Println(a << 3) // 64 - 00001000 becomes 01000000
	fmt.Println(a >> 3) // 1  - 00001000 becomes 00000001

#### Decimal types
