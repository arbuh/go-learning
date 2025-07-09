# Go

## Glossary

Initializer - a value we asign to a variable while declaring it
Naked return - a return statement without argumetns, which returns named return values
Name - everything can be declared, e.g. `var myInt int` or `func helper()`
Named return value - values in a function that can be returned by a naked return and have names providing meaning behind the values

## Packages

`main` - a package with the entry point

This is how we can import other packages:

```
import (
	"fmt"
	"math/rand" <- use `rand.*` to call items in the package
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





