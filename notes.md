# Go

## Glossary

Build in package - an implicit package containing the build-in functions
Initializer - a value we asign to a variable while declaring it
Naked return - a return statement without argumetns, which returns named return values
Name - everything can be declared, e.g. `var myInt int` or `func helper()`
Named return value - values in a function that can be returned by a naked return and have names providing meaning behind the values
Panic - an error

## Packages

`main` - a package with the entry point

This is how we can import other packages:

```
import (
	"fmt"
	"math/rand" // <- use `rand.*` to call items in the package
)
```

You can also use a statemene like `import fmt` for each imported package, but it is a bad style.

### Names

You can export names from a package only if they start with a capital letter e.g. `Pi`. So `pi` is an unexported name.

## Functions

This is an example of a function in Go:
```
func add(x int, y int) int {
    return x + y
}
```

You can skip a repeated type and it is more common in the Go ecosystem:
```
func add(x, y int) int {
	return x + y
}
```

### Multiple return values

A function can return multiple results:
```
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```

### Named values

You can name return values and use so-called "naked return" to return them:
```
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```
However, it is not very common to use them because they decrease readability.

## Variables

`var` can declare a list of variables:
```
var c, python, java bool
```

We can also skip the type if we provide initializers (must be one per a variable):
```
var c, python, java = true, false, "no!"
```

The short declaration allows to omit the keyword `var`, but it can be used only inside a function:
```
a := 15
```

You can also decrare multiple varaibles in a block, as for the import statements:
```
var (
    MyInt    int    = 1
    MyString string = "a" 
 )
 ```

## Basic types

Go has the following basic types (i.e. primitives):
* `bool`
* `string`
* `int  int8  int16  int32  int64` - signed integers. Use `int` if you don't have a strong reason to use something else. 
                                    `int` is either 32 or 64 bits, depends on whether the system is 32 or 64 bit
* `uint uint8 uint16 uint32 uint64 uintptr` - unsigned integers
* `byte` - an alias for `uint8`
* `rune` - an alias for `int32`
* `float32 float64`
* `complex64 complex128`

Varables declared without an initial value are given zeros: `0`, `false` or `""`

### Type conversions


You can use a function with a name of a type to conver a variable to this type:
```
var i int = 10
var f float64 = float64(i)
```

## Constants

This is how you declare a constant:
```
const Pi = 3.12
```

You cannot use `:=` with constants.

Numeric constants are special type of numeric value. It is competable with all numeric basic types, it is high-precision, and it can be larger than 64-bits values.
You define them without specifying a type:
```
const (
    pi = 3.14159265358979323846
    radius = 10
)
```

## Flow control

### Loops

Go has only for-loop where you define the initial statement, the condition evaluated before an iteration, and the post statement evaluated at the end of an iteration:
```
for i := 0; i < 10; i++ {
		sum += i
	}
```

The initial and post statements are optional, and it allows you to "implement" while-loops:
```
sum := 1
for sum < 1000 {
	sum += sum
}
```

You can also omit the condition to have an infinite loop:
```
for {
}
```

Use `break` to exit a loop.

### If


Like for the for-loop, you don't need parenthesis for if:
```
if x > 0 {
    return y
}
```

You can define a statement to be executed before the condition. The scope of the variables defined there is limited by the if and else statements:
```
if v := math.Pow(x, n); v < lim {
	return v
} else {
    fmt.Printf("%g >= %g\n", v, lim)
}
```

### Switch

You can use not only constants but a pretty wide range of statements in the switch-statement in Go:
```
switch x {
    case 1:           // constant
    case y:           // variable allowed!
    case max:         // function result allowed (i.e. we have `max := getMaxValue()` above)
    case y * 2:       // expression allowed!
    case getOther():  // function call allowed!
    default:
}
```

An eveluation of a switch also ends when the case succeeds.

A switch without a condition can be used a syntax-sugar over a long if-statement:
```
switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
}
```

### Defer

The keyword `defer` allows to define a statement that is executed only when a function returns. It is widely used for the resource management:
```
func processFile() error {
    file, err := os.Open("data.txt")
    if err != nil {
        return err
    }
    defer file.Close()  // Right next to Open()
    
    // 100 lines of code...
    // Easy to see that file will be closed
}
```

Deferred functions are stored in a stack, i.e. the latest ones are executed the first.

## Complex types

### Pointers

A pointer holds the memory address of a value. It has a type `*T` where `T` is the value's type:
```
var p *int
```

You can generate a pointer from a variable using the `&` operator:
```
i := 42
p = &i
```

Then you can access or set a value using the operator `*`:
```
fmt.Println(p*) // <- to access the value of `i`
*p = 21 // <- to set the value of `i`
```

### Structs

A struct is a collection of fields:
```
type Vertex struct {
    X int
    Y int
}
```

If you have a pointer to a struct, you can omit the `*` operator:
```
v := Vertex{1, 2}
p := &v
p.X = 2 // i.e. it is not necessery to write (p*).X = 2
```

You can also set values to individual fields in any order:
```
v := Vertex{Y:11} // `X` is 0 then
```

### Arrays

You have to specify a size of an array on its initialization:
```
var a [10]int
```

You can slice an array:
```
var b []int = a[1:3] // i.e. including the element with the index 1 up to excluding 3
```

Since slices don't contain data and refer to an underlying array, if you modify a slice you also modify the reference array.

It is not necessery to initialize an array first before initializing a slice.
You can init a slice using the slice literal an it will create the underlying array out of the box:
```
a := []int{1, 2, 3, 4, 5, 6}
```

A slice had a length and a capacity. Then can be accesed via `len(a)` and `cap(a)`.
The length is a number of elements in the slice.
The capacity is a number of elements in the underlaying array counting from the first element in the slice.
The capacity is important because it allows to control memory allocations.
So the runtime doesn't need to reallocate memory if there is still sufficient capacity.

You can also have nil slices, which have the zero value `nil` and the length and capacity equal to zero:
```
var a []int
```

You can create a slice using the `make` function:
```
a := make(int[], 3, 5) // len(a)=3, cap(a)=5, the allocated elements are zeros
```

To add elements to a slice, use the `append` function:
```
a = append(a, 1, 2) // the first argument is the target slice, and the rest is varargs with new elements
```

Be carefull, it returns a new slice, instead of modifying the provided one!

### Range

The function `range` allows to iterate over a slice or a map:
```
for i, v := range a {
    fmt.Printf("%d:%d\n", i, v)
}

// Skipping the indexes:
for _, v := range a {
    fmt.Printfln(v)
}

// Skipping the values:
for i := range a {
    fmt.Printfln(i)
}
```

### Map

You usually use the `make` function to create a map, otherwise the map has `nil` value and you cannot add items to it:
```
var m map[string]int
m = make(map[string]int)
```

Althernatively, you can use the map literals:
```
var m = map[string]int{
    "a": 1,
    "b": 2,
}
```

Map mutations:
```
m[key] = elem       // <- add or replace an element
elem = m[key]       // <- access an element
delete(m, key)      // <- delete an element
elem, ok = m[key]   // <- access an element and check if it is present in the map
```

### Functions

You can app functions as function arguments:
```
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

hypot := func(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

compute(hypot)
```

A function can be also a closure:
```
func fibonacci() func() int {
	prev := 0
	curr := 0
	return func() int {
		sum := prev + curr
		prev = curr
		curr++
		return sum
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
```

## Methods and interfaces

### Methods

You can define methods on types via a receiver argument:
```
type Person struct {
    Name string
}

func (p Person) greet(other string) string {
    return fmt.Sprintf("Hello %s, I'm %s", other, p.Name)
}

p := Person{"Alice"}
p.greet("Bob"}
```

N.B., you can declare a method only in the same package where the type is defined!

A method with a type value (as in the exemple above) gets a copy of the value.
If you want to modify the provided value, use a pointer receiver:
```
func (p *Person) changeName() {
    p.Name = "Changed"
}
```

### Interfaces

You don't need to expicetely declare implementations of interfaces:
```
type I interface {
	M()
}

type T struct {
	S string
}

// By this method, the type implements the interface
func (t T) M() {
	fmt.Println(t.S)
}
```

Under the hood, interface is a tuple of the value and the implementation type.
Keeping the type allows to understand which method must be called when we call a method on the interface type.

You can call a method on an interface implemented via a type with `nil` value.
However, you cannot call a method on an interface without an implementation.

Q: What happens when twoo interfaces share the same name?

### Zero interface

Zero interface can be of any type, because all types implement at least zero methods:
```
var i interface{}
i = 42
i = "hello"
```

It allows to implement functions that handle arguments of any type.

### Type assertions

You can get from an interface a concrete value as well as a flag whether it is possible:

```
t, ok = i.(T)
```

### Type switch

You can also use a type switch to get concrete values from an interface.
It has syntax similar to type assertions, but you write `type` in the parenthesis instead of the type name:
```
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

## Build-in interfaces

### Error

Go has in-build interface for errors:
```
type error interface {
    Error() string
}
```

Functions often return `error` instances as a second return value.
You can check them whether they are equal to `nil`, and do error processing if they are not.

You can create your own implementation of `error` by implementing the `Error()` method for your error types.

### Stringer

The `Stringer` interface is heavily used by the `fmt` package:
```
type Stringer interface {
    String() string
}
```

Make your types implementing this interface to be accepted by the functions in `fmt`.

### Reader

For reading a stream of data, Go uses the `io.Reader` interface.
It has a method `func (T) Read(b []byte) (n int, err error)` that populates the provided array with read bytes.
`n` represents here the number of bytes populated, and `err` is populated with `io.EOF` when the stream ends.

## Generics

You can provide generics to a function the following way:
```
func Index[T comparable](s []T, x T) int
```

The keyword `comparable` means here that it is possible to use the `==` and `!=` operators on the type T.

It is also possible to use generics in types:
```
type List[T any] struct {
    next *List[T]
    val  T

```

If you want to implement a method for such a type, you don't need to specify the type parameter because the type already has it:
```
func (l *List[T]) Add(val T) *List[T] {
	newNode := &List[T]{l, val}
	return newNode
}
```

## Concurrency

### Goroutines

A goroutine is a lightweigth thread that can be called using the keyword `go` before calling the function must be executed in the goroutine:
```
go f(arg1, agr2)
```

The evaluation of `f`, `arg1` and `arg2` happens in the current goroutine, while the execution of `f` in the new one
(evaluation of `f` means that Go determines which function `f` refers to, whithout executing it).

Goroutines share the same memory address space, even when they run in the same threat.
So you must be carefull with the shared variables.

### Channels

Channels are an abstraction that allows to safely share values among goroutines.

You must create a channel before use:
```
ch:= make(chan int)
```

Then you can send and recive values to/from the channel:
```
ch <- v // Sending a value
v := <-ch // Receiving a value
```

#### Buffers

You can also specify how many values a channel can keep before the values are consumed by goroutines:
```
ch := make(chan int, 2)
```

Such a channel is called "buffered".
If a buffured channel is full and you send a value there, you get a deadlock error.
The same happens if you try to receive a value from the channel when it is empty.

If you don't specify a buffer (e.g. `ch := make(chan int)`), its size is zero.
It means that if there is no goroutine consiming the values, you will get a deadlock error on sending a value:
```
ch := make(chan int)
ch <- 1 // Here we will get a deadlock error
```

So unbuffered channels require synchronous communication!

A bugger's size can be abtained using the build-in function `cap`:
```
cap(ch)
```

#### Closing

A sender can communicate there will no more values sending to a channel using the keyword `close`.
A receiver can check if there will be more values from the channel by checking a flag-field provided as a second parameter in the channel reading:
```
v, ok <- ch
```

You can read values from a channel in a loop until the channel is closed:
```
for v := range ch {
    ...
}
```

#### Select

If a gorouting reads from more than one channel, you should use the `select` statement:
```
select {
    case a := <- ch1:
        ...
    case b := <- ch2:
        ...
}
```

You can also process situations when no case is ready under the `default` case.

If you don't use `select`, the gorouting will be blocked if there is no value avaliable in a channel.
Then it will not be reading even from the channels that have values.

#### Mutex

Mutex allows mutual exclusion of value access among goroutines:
```
mu sync.Mutex
```

It has two methods: `Lock` and `Unlock`.
Only a singe goroutine can access a code between `Lock` and `Unlock` (or till the end of the function if you use `defer mu.Unlock`):
```
mu.Lock()
safeSharedCounter++
mu.Unlock()
```







