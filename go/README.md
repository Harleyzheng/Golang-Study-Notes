# Go

Table of Contents
=================

   * [Go](#go)
   * [Table of Contents](#table-of-contents)
      * [Why Go?](#why-go)
         * [runs fast](#runs-fast)
         * [garbage collection](#garbage-collection)
         * [simpler object](#simpler-object)
         * [efficient concurrency](#efficient-concurrency)
      * [Charactristics of Go](#charactristics-of-go)
         * [workspace &amp; packages](#workspace--packages)
         * [go tool](#go-tool)
         * [Variables](#variables)
         * [Type](#type)
         * [Pointers](#pointers)
         * [Variable Scope](#variable-scope)
         * [Deallocating Memory](#deallocating-memory)
         * [Garbage Collection](#garbage-collection-1)
         * [Comments](#comments)
         * [Printing format strings](#printing-format-strings)
         * [Type Conversions](#type-conversions)
         * [Encoding](#encoding)
         * [String](#string)
         * [Unicode Package.](#unicode-package)
         * [Strings Package](#strings-package)
         * [Strconv Package](#strconv-package)
         * [Constants](#constants)
         * [iota -](#iota--)
         * [Control flow](#control-flow)
      * [Composite data types](#composite-data-types)
         * [Slices](#slices)
         * [Variable Slices](#variable-slices)
         * [Maps](#maps)
         * [Structs](#structs)
   * [Communication Protocols &amp; Formats](#communication-protocols--formats)
         * [Packages for RFCs](#packages-for-rfcs)
   * [Function](#function)
   * [Function Types](#function-types)
   * [Classeses](#classeses)
   * [Encapsulation](#encapsulation)

## Why Go? 

### runs fast

  Machine language. Assembly language. High level language.

  Compilation: Translation occurs once before execution. High level language must be translated into machine language.

  Interpretation: translate on the fly. 

  Go is a good compromise between compiled vs interpreted. 

### garbage collection

  efficient, automatic garbage collection.

### simpler object

  Weakly object oriented. 

  Go does not use class. Go uses structs with methods. 

  No inheritance, constructors, generics.

### efficient concurrency

  Goroutines - thread

  Channels - communication between tasks

  Select



## Charactristics of Go

### workspace & packages

  Directory hierarchy. src, pkg, bin.. (can be customized)

  GOPATH: default workspace

  Packages. Group of related source files.

​    'package' defines the file it's in

​    'import' includes other packages

​    building main package generates an executed program.



### go tool

​    'import' used to access other packages. it searches directories specified by GOROOT and GOPATH.

​    'go build', 'go doc', 'go foramt', 'go get', 'go list', 'go run', 'go test'



### Variables

  declaration

```go
var x int = 10

var a,b,c int

x := 100 // x not declared yet
```

 

### Type

  define alias for a type

  Uninitialized vars have a zero value. 0, "", etc



### Pointers

  is an address to data in memory

  & returns the address 

  \* dereference. returns data at the address

```go
var x *int = &y 
```

  

  new() creates a variable and returns a pointer to the variable



### Variable Scope

  blocks - universe block, package block, file block, if/else

  lexcial scoping. Bi >= Bj, if Bj is defined inside Bi. 

  Definition. Variable accessible from block Bj if 

​    \1. Variable is declared in block Bi, and

​    \2. Bi >= Bj



### Deallocating Memory

  heap - is persistent.

  stack - local variables in function. deallocated after function completes

  hard to determine when a variable is no longer used. i.e. a function returning a pointer.



### Garbage Collection

  garbade collection runs in the background. (trade off with performance)

  Compiler determines stack vs heap.



### Comments

  single line: //

  block /* */ /*



### Printing format strings

  import from fmt package. 

  fmt.Printf("hi %s", "name");

  fmt.Printf("my number %d", 12);



### Type Conversions

  T():

​    var x int32 = 1

​    var y int16 = 2

​    x = int32(y)



### Encoding

  UTF-8. Unicode. 

  Code point is a term for a Unicode character. 2^32 code points. 

  In Go, a code point is a Rune. 



### String

  String. Each byte is a Rune. Immutable. 



### Unicode Package. 

``` go
IsDigit(r rune)
IsSpace(r rune)
IsLetter(r rune)
IsLower(r rune)
IsPunct(r rune)
ToUpper(r rune) // conversion
ToLower(r rune) // conversion
```



### Strings Package

``` go
Compare(a, b)
Contains(a, b)
HasPrefix(s, prefix)
Index(s, substr)

Replace()
ToLower(s) 
ToUpper(s) 
TrimSpace(s) // return a new string.
```



### Strconv Package

```    go
 Atoi(s) 
 Itoa(s) 
 FormatFloat() 
 ParseFloat()
```



### Constants

  Expression whose value is known at compile time. Type is inferred from righthand side.



### iota - 

​    used when representing a property which has distinct possible values. Like enum.

​    Constants must be different but actual value is not important.

```go
type Grades int
const (
  A Grades = iota     
  B
  C
)
```

​    

### Control flow

  if,

  for loop, for i < 10 {}, for

  switch (auto break at each case. no fall through.)

  tagless switch. first true case is executed.

  break, continue. same as c++.

  Scan reads user input and is blocking. 



## Composite data types

Arrays

  fixed length. 

  elements indexed using [].

  elements initialized to zero value.

```go
var x [5]int = [5]{1, 2, 3, 4, 5}
```



  ... infers size of array

```go
var x [5]int = [...]{1, 2, 3, 4, 5}
```

  

  iterate through arrays. for loop.

```go
for i, v range x {

}
```



### Slices

  A window on an underlying array

  size can change

  Pointer indicates the start of the slice

  Length len()

  Capacity is max number of elements. cap() 

``` go
arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}

s1 := arr[1:3] (includes 1, excludes 3)
s2 := arr[2:5]

len(s1) // 2
cap(s1) // 6
```

  writing to a slice changes underlying array

  Slice Literals - Creates the underlying array and creates a slice. Slice points to the start of the array, length is capacity.

```   go
sli := []int{1, 2, 3}
```



### Variable Slices

make(). Create a slice (and array)

make() with 2 args: type, length/capacity

it initializes to zero values

``` go
sli = make([]int, 10)
```

make() with 3 args: type, length, capacity

append() adds elements to the end of a slice, 

it inserts into underlying array

increases size of array if necessary

``` go
sli = make([]int, 0, 3)
sli = append(sli, 100) 
```



### Maps 

Hash Tables concepts same as other languages

Advantage: can use arbitrary key. i.e. slices, arrays

**make() to create a map**

```go
  var idMap map[string][int] // key type, value type
  idMap = make(map[string]int)
```

**Map literal**

``` go
  idMap := map[string]int {
    "joe": 123
  }
```

``` go
idMap["hayley"] // Accessing Maps. Returns zero if key is not present.
idMap["hayley"] = 4 // Adding a k,v pair
delete(idMap, "hayley") // Deleting a k,v pair
id, p := idMap["joe"]  // two-value assignment. id is value, p is presence boolean of key
len(idMap) // returns number of elements

// iterating through a map use a for loop
for key, val := range idMap {
}
```



### Structs

Aggregate data type.

``` go
  type Person struct {
     name string
     address string
     phone string
  }

  var hayley Person
  hayley.name = "hayley z"
  x = hayley.address

  p1 := new(Person) // initialize fields to zero
	p1 := Person{name: "Hayley", address: "a st.", phone: "123"} // struct literal
```



# Communication Protocols & Formats

### Packages for RFCs

Functions with encode and decode protocol format

For example "net/http" "net" 

**JSON Marshalling  & Unmarshalling**

``` go
p1 := Person(name: "hayley", address: "home")
jsonObject, err := json.Marshal(p1) // jsonObject is a []byte
var p2 Person
err := json.Unmarshal(jsonObject, &p2) // 
```

**Files**. File access are linear access. 

**Basic Operations**

	1. Open - Get handle for access
 	2. Read - read bytes into []byte
 	3. Write - write []byte into file
 	4. Close - Release handle
 	5. Seek - move read/write head

**"io/ioutil"**

``` go
dat, e := ioutil.ReadFile("test.txt")  // dat is []byte with contens of entire file
// Large files may cause a memory problem

err := ioutil.WriteFile("test.txt", "Hello", 0777) // Create, dump everything to file. 0777 is permission bytes for everyone
```

**"os"**

``` go
os.Open() // opens a file, returns a file descriptor
os.Close() // close a file
os.Read() // reads from a file into []byte, and fills the []byte. Control the amount read.
os.Write() // write

// Opening and reading
f, err := os.Open("dt.txt")
byteArray := make([]byte, 10)
numBytesRead, err := f.Read(byteArray) // returns number of bytes read, maybe less than []byte
f.Close()

// Create/Write

f, err := os.Create("outfile.txt")

byteArray := []byte{1, 2, 3}
nb, err := f.Write(byteArray)
nb, err := f.WriteString("hi Hayley")
```



# Function

main() is called immediately. 

Parameters are defined in function declaration. Arguements are passed in function calls.

**Multiple return values**

``` go
func foo(x int) (int, int) {
  return x, x + 1;
}
a, b := foo(3)
```

**Call by Value**

Passed arguments are copied parameters.

- Advantage: Data Encapsulation

- Disadvantage: Large objects may take a long time to copy.

**Call by Reference**

``` go
func foo(x *int) {
	*x = *x + 1
}
x := 2
foo(&x)
```

- Advantage: Don't need to copy arguments. 
- Disadvantage:  Data Encapsulation

**Passing Arrays**

Call by value - Array arguments are copied

Pass array pointers - (**NOT the neat way in Go)**

``` go
func foo(x *[3]int) {
	(*x)[0] = (*x)[0] + 1
}
x := [3]int{1,2,3}
foo(&x)
```

Instead, pass Slices instead. Passing a slice copies the pointer.

``` go
func foo(sli int) int {
	sli[0] = sli[0] + 1
}
x := [3]int{1,2,3}
foo(x)
```



# Function Types

**First-Class Values**

Functions are first-class. Can be treated like other types.

Can be created dynamically. Can be passed as arguments and returned as values.

Can be stored in data structures.

**Declare a variable as a function**

```go
var funcVar func(int) int 
func incFn(x int) int {
	return x + 1
}
func main() {
	funcVar = incFn
}
```

**Functions as Arguments**

```go
func applyIt(aFunc func (int) int, val int) {
	return aFunc(val)
}
```

**Anonymout Functions**

```go
applyIt(func (x int) int {return x + 1}, 2)
```

**Returning Functions**

```go
getFunc() func (x int) int {
	return func (x int) int {return x + 1}
}
```

**Closure**

When passing a function, the function and its environment are also passed. Just like Javascript.

**Variable Argument Number**

```go
func getMax(vals ...int) int {  // vals is treated as a slice

}
```

**Variadic Slice Argument**

```go
func main() {
	getMax(slice...)
}
```

**Deferred Function Calls** - Typically used for cleanup activities

**Arguments of a deferred call are evaluated right away**

```

```



# Classeses

**Receiver Type**

```go
type MyInt int

// Method (Double) has a receiver type **(mi MyInt)** that it's associated with. 
func (mi MyInt) Double () int {
	return int(mi*2)
}

func main() {
  v := MyInt(3)
  v.Double() // v is passed by value as an argument to Double here
}
```

**Implict Method Argument**

Receiver Type is commonly a struct. 

```go
type Point struct {
  x float64
  y float64
}

func (p Point) DistToOrig() {
  t := math.Pow(p.x, 2) + main.Pow(p.y, 2)
  return math.Sqrt(t)
}

func main() {
  p1 := Point{3, 4}
  fmt.Println(p1.DistToOrig()) // p1 implicitly pass to the function
}
```

# Encapsulation

**Controlling Package Access** 

**public functions**

```go
// file 1
package data
var x int=1
func PrintX() {fmt.Println(x)}

// file 2
package main
import "data"
func main() {
	data.PrintX()
}
```

**Structs** with public functions

``` go
package data

type Point struct {
  x float64  // hidden
  y float64  // hidden
}

func (p *Point) InitMe (xn, yn float64) {
  p.x = xn // no need to dereference p
  p.y = yn
}

```

**Pointer Receivers**

Because receivers are passed in by value, the actual value cannot by changed. Another problem is receivers could be large. We can pass a pointer to address both problems. 

**When using pointer receivers, no need to dereference inside the method, and no need to reference when calling the method.**

Good practice. Either all methods use pointer receivers, or none of the methods use pointer receivers.

```go
func main() {
  p := Point(x, y)
  p.InitMe(2,4) // no need to reference p
}
```

