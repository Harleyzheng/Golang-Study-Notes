# Go

## Table of Contents
    * [Why Go?](#why-go)
        * [Advantages:](#advantages)
        * [Charactristics of Go](#charactristics-of-go)
       * [workspace &amp; packages](#workspace--packages)
       * [go tool](#go-tool)
       * [Variables](#variables)
       * [Type](#type)
       * [Pointers](#pointers)
       * [Variable Scope](#variable-scope)
       * [Deallocating Memory](#deallocating-memory)
       * [Garbage Collection](#garbage-collection)
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


## Why Go?

## Advantages:

1. **runs fast.**

  Machine language. Assembly language. High level language.

  Compilation: Translation occurs once before execution. High level language must be translated into machine language.

  Interpretation: translate on the fly. 

  Go is a good compromise between compiled vs interpreted. 

2. **garbage collection.**

  efficient, automatic garbage collection.

3. **simpler object.** 

  Weakly object oriented. 

  Go does not use class. Go uses structs with methods. 

  No inheritance, constructors, generics.

4. **efficient concurrency.** 

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



### Composite data types

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

  make() to create a map

```go
  var idMap map[string][int] // key type, value type
  idMap = make(map[string]int)
```

Map literal

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
  type struct Person {
     name string
     address string
     phone string
  }

  var hayley Person
  hayley.name = "hayley z"
  x = hayley.address

  p1 := new(Person) // initialize fields to zero
  p1 := Person(name: "Hayley", address: "a st.", phone: "123") // struct literal
```









