# tPassword

A small Go library for generating random passwords, built on `crypto/rand`.

## Installation

```bash
go get github.com/tPortal-Dev/tPassword
```

## Usage

### Simple generation

By default, `Generate` creates a password using lowercase letters and digits:

```go
package main

import (
    "fmt"
    "log"
	
    "github.com/tPortal-Dev/tPassword"
)

func main() {
    password, err := tPassword.Generate(16)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(password)
}
```

### Options

The default behavior can be customized by passing `Option`s to `Generate`:

| Option               | Effet                                                  |
|----------------------|--------------------------------------------------------|
| `WithUppercase()`    | Adds uppercase letters                                 |
| `WithSymbols()`      | Adds symbols (`.?!&#()$%+-=_-@`)                       |
| `WithoutLowercase()` | Removes lowercase letters                              |
| `WithoutDigits()`    | Removes digits                                         |
| `WithoutAmbiguous()` | Removes ambiguous characters (`O`, `0`, `1`, `l`, `I`) |

```go
password, err := tPassword.Generate(20,
	tPassword.WithUppercase(),
	tPassword.WithSymbols(),
	tPassword.WithoutAmbiguous(),
)
```

> [!WARNING]
> At least one character set must remain enabled, otherwise `Generate` returns `ErrEmptyCharset`.

### Custom charset

To define your own allowed characters:

```go
password, err := tPassword.GenerateWithCustomCharset(12, "abc123!?")
```

### Multiple generation

Generate multiple passwords at once using the same options or the same custom charset:

```go
// Using options
passwords, err := tPassword.GenerateMulti(5, 16, tPassword.WithUppercase())

// Using a custom charset
passwords, err := tPassword.GenerateMultiWithCustomCharset(5, 16, "abc123")
```

### Errors

The library exposes two sentinel errors that can be checked with `errors.Is`:

```go
var (
    tPassword.ErrEmptyCharset  // no character set selected
    tPassword.ErrInvalidLength // length <= 0
)
```