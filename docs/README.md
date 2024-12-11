# Golang Tutorial

## Installation

Visit this [link](https://go.dev/doc/install) to install `go` on your machine.

> [!IMPORTANT]
> This Markdown is made by my personal researches, so if you want to learn from better source, I strongly recommand you to [Take a tour on Golang](https://go.dev/tour/)

### What is GOPATH?

In Golang, `GOPATH` is an environment variable that specifies the **root directory of your Go workspace**. The Go workspace is where Go code is organized and managed. It includes source code, compiled binaries, and package objects.

#### Structure of `GOPATH`

A typical GOPATH directory contains three main subdirectories:

1. `src`: This directory holds the source code of your Go projects and external packages.
2. `pkg`: This directory stores compiled package files, which Go uses to speed up builds.
3. `bin`: This directory stores compiled binary executables.

For example, if GOPATH is set to `/home/user/go`, the structure looks like:

```
/home/user/go/
├── bin/
├── pkg/
└── src/
```

> [!NOTE]
> Since the introduction of `Go modules` *(starting with Go 1.11)*, **the use of GOPATH has become less critical**. With modules, Go projects can exist outside the GOPATH directory, and dependencies are managed in a `go.mod` file. This change enables Go developers to organize projects more flexibly without relying on the strict workspace layout.
> In module mode, the `GOPATH` is mainly used for caching downloaded dependencies in the `GOPATH/pkg/mod directory`.

**Key Points**

- Before modules, all Go projects had to reside within the GOPATH/src.
- With modules enabled, you can place your projects anywhere, but GOPATH still acts as a default location for dependency caching and storage.
- Use the go env command to check your current GOPATH:

```
go env GOPATH
```

### Packages vs Modules

In Golang, **packages** and **modules** are essential concepts for organizing and managing code. Understanding the distinction and their roles helps in structuring Go projects effectively.

---

#### **1. Package**

###### Definition

A **package** in Go is a collection of Go source files in the same directory that are compiled together. Packages enable code reuse and organization. Every Go file belongs to a package.

#### Key Features

- **Declared at the Top of the File**: Every `.go` file starts with a `package` declaration.

  ```go
  package main
  ```

- **Single Directory = One Package**: All `.go` files in a directory belong to the same package.

#### Types of Packages

1. **Main Package**:
   - Special package for defining the entry point of a Go application.
   - It contains a `func main()` which is executed when the program runs.
   - Example:

     ```go
     package main
     
     import "fmt"
     
     func main() {
         fmt.Println("Hello, Go!")
     }
     ```

2. **Library Packages**:
   - Used to organize reusable code.
   - These packages cannot have a `main()` function.
   - Example:

     ```
     ├── mathutils/
         ├── mathutils.go
         ├── advanced.go
     ```

     `mathutils.go`:

     ```go
     package mathutils
     
     func Add(a, b int) int {
         return a + b
     }
     ```

#### Importing Packages

Go provides many built-in packages (e.g., `fmt`, `os`) and allows importing custom packages:

```go
import (
    "fmt"
    "myproject/mathutils"
)
```

---

### **2. Module**

#### Definition

A **module** is a collection of related Go packages. It is the unit of versioning and distribution in Go. Modules are defined and managed using the `go.mod` file.

#### Key Features

- **Introduced in Go 1.11**: Modules replaced the older `GOPATH`-based dependency management system.
- **Multi-Package Support**: A module can contain multiple packages organized in subdirectories.
- **Versioned**: Modules have versions (e.g., `v1.2.3`) for dependency management.

#### Creating a Module

1. Initialize a module with:

   ```bash
   go mod init example.com/myproject
   ```

2. This creates a `go.mod` file:

   ```go
   module example.com/myproject
   
   go 1.20
   ```

#### Module Structure

```
myproject/
├── go.mod
├── go.sum
├── main.go
└── mathutils/
    ├── mathutils.go
    └── advanced.go
```

- **`go.mod`**: Tracks the module name and dependencies.
- **`go.sum`**: Records checksums for module versions to ensure integrity.
- **Packages**: Organized as directories within the module.

---

### **Relationship Between Packages and Modules**

- A **module** can contain multiple **packages**.
- For example:
  - Module: `example.com/myproject`
  - Packages:
    - `example.com/myproject/mathutils`
    - `example.com/myproject/utils`

---

### **Comparison: Package vs Module**

| Feature               | Package                               | Module                                |
|-----------------------|---------------------------------------|---------------------------------------|
| Definition            | Collection of source files in one directory | Collection of related packages        |
| Purpose               | Code reuse and organization          | Dependency management and distribution|
| Scope                 | Single directory                     | Entire project                        |
| Declaration           | `package <name>` at the top of a file| `go.mod` file defines the module      |
| Examples              | `fmt`, `os`                   | `example.com/myproject`               |
| Dependency Management | Not versioned                        | Versioned                             |

---

### **Example**

1. **Module Definition** (`go.mod`):

   ```go
   module example.com/myproject
   
   go 1.20
   ```

2. **Main Package**:

   ```go
   package main
   
   import (
       "fmt"
       "example.com/myproject/mathutils"
   )
   
   func main() {
       fmt.Println(mathutils.Add(2, 3))
   }
   ```

3. **Custom Package** (`mathutils/mathutils.go`):

   ```go
   package mathutils
   
   func Add(a, b int) int {
       return a + b
   }
   ```

---

###### Summary

- A **package** organizes code into reusable units within a project.
- A **module** groups related packages together for versioned distribution.
- Modules allow Go projects to manage dependencies, enabling a flexible and scalable way to work with Go code.

### Packages

Programs start running in package `main`. By convention, the package name is the same as the last element of the import path.
For instance `math/rand` package comprises files that begin with the statement `package rand`.

### Imports

There is 2 way for importing:

1. Factoring import, using parenthesize

```go
import (
"fmt"
"math"
)
```

2. Regular import statements

```go
import "fmt"
import "math"
```

### Exports

In Go, a name is exported if `it begins only with capital letter`. For instance, `Pi` is exported from math package but `pi` is not exported because it doesn't started with capital letter.

When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.

### Go Declaration syntax

In Go, declare a variable is look like this:

```go
var variable_name typeOfVariable; 
var pointer_name *typeOfVariable; 
```

**Examples:**

```go
var number int;
var name string;
var arr [3]int;
var p *int;
```

### Functions

A function can take zero or more arguments, with pre-defined type. Notice that types come after the variable names.

> [!NOTE]
> The `main` function takes no argument.

###### Declare a function in Go

```go
// Declare both as int
func sum(num1 int, num2 int) int { ... }

// Define both num1 and num2 as int
func multiple(num1, num2 int){ ... }

// Get a function as a argument and return function
func f (func(int, int) int, int) func (int, int) int { /* ... */ }

// Declare a function in a single line
sum := func (a,b int) int { return a+b };
```

#### Multiple returns

A function can return any number of results.
The swap function returns two string

```go
func swap(x, y string) (string, string) {
 return y, x
}

func main() {
 a, b := swap("hello", "world")
 fmt.Println(a, b)
}
```

#### Named return values

Go's return values may be named. If so, they are treated as variables defined at the top of the function.
These names should be used to document the meaning of the return values.
A **return statement without arguments** returns the named return values. This is known as a `naked` return.

The var statement declares a list of variables; as in function argument lists, the type is last.

### Variables

> [!NOTE]
> The var statement declares a list of variables; as in function argument lists, the type is last.
> A var statement can be at package or function level.

### Variables with initializers

A var declaration can include initializers, one per variable.


### Pointers

Note that **values are stored in memory**.
There is 2 type of variables:

- **Values is variables**:

Stored until the variable not used anymore in code

```go
name := "Amir"
```

- **Just-in-time values**:

Only used once, not stored after that.

```go
fmt.Println("Amir")
```

> [!IMPORTANT]
> **Garbage collector**:
> There is an ability in Golang names **Garbage Collector** that are responsible for remove and free up memory from un-used values. When you define a variable and not use it in code, you will get an exception in while compile, because of the `garbage collector`.

Values are stored in **computer memory** and offers various `Storage slots` with individual addresses.

```go
var myName := "Amir" // 0x0000018050
var myAge := 20 // 0x0000025011
```

**Variables know the address of the value and get the value from there when needed.**

#### What is Pointer

Pointer is a variable that only store the address of the value in memory instead of store the value itself.

Pros of pointers:

- Can passed to a function as a parameter and avoid unnecessary copies from value, less usage of memory.

- Function can directly work `on/with` pointers. (e.g Change them in memory)
It means that by passing a pointer as a parameter to a function, the changes in memory will happen and it is unnecessary to return the pointer.

**Examples:**

```go
var name string = "Amir"
var pName *string = &name
fmt.Println(name) // Amir
fmt.Println(pName) // 0xc00005e040
fmt.Println(*pName) // Amir

name += " Mahdi"

fmt.Println(name) // Amir Mahdi
fmt.Println(pName) // 0xc00005e040
fmt.Println(*pName) // Amir Mahdi

*pName = "Ali"
fmt.Println(name) // Ali
fmt.Println(pName) // Ali
```

```go
number := 20
pNumber := &number

fmt.Println(number) // 20
fmt.Println(pNumber) // 0xc00508e120
fmt.Println(*pNumber) // 20

*pNumber = 100
fmt.Println(number) // 100
fmt.Println(*pNumber) // 100
```

##### Using pointers in function

```go
func sayHello(message *string) string {
    *message = "Hello, " + *message
}

func main(){
    message = "Amir"
    sayHello(&message)
    fmt.Println(message)
}
```
