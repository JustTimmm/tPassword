# tPassword
A small Go library for generating random passwords, built on `crypto/rand`.

## Compatibility
- Go 1.26+

## Installation

```bash
go get github.com/tPortal-Dev/tPassword
```

<!--
## Package

| Package | Import path                        | Description            |
|---------|------------------------------------|------------------------|
|v1       | `github.com/tPortal-Dev/tPassword` | Only tPassword version |
-->

## Usage

### Generate

```go
import "github.com/tPortal-Dev/tPassword"

// default configuration
password, err := tPassword.Generate(16)

// custom options
password2, err := tPassword.Generate(16,
    tPassword.WithUppercase(),
    tPassword.WithSymbols(),
    tPassword.WithoutAmbiguous(),
)

// multiple passwords
passwords, err := tPassword.GenerateMulti(3, 16)
```

### Generate with custom charset

```go
import "github.com/tPortal-Dev/tPassword"

password, err := tPassword.GenerateWithCustomCharset(16, "abcd1234?!-_")

// multiple passwords
passwords, err := tPassword.GenerateMultiWithCustomCharset(3, 16, "abcd1234?!-_")
```

### Default behavior

By default, generated passwords contain:

- lowercase letters (`a-z`)
- digits (`0-9`)

### Options

The default options can be customized by passing `Option`s to `Generate / GenerateMulti`:

| Option               | Effect                                                 |
|----------------------|--------------------------------------------------------|
| `WithUppercase()`    | Adds uppercase letters                                 |
| `WithSymbols()`      | Adds symbols (`.?!&#()$%+=_-@`)                        |
| `WithoutLowercase()` | Removes lowercase letters                              |
| `WithoutDigits()`    | Removes digits                                         |
| `WithoutAmbiguous()` | Removes ambiguous characters (`O`, `0`, `1`, `l`, `I`) |

> [!WARNING]
> At least one character set must remain enabled.
> Otherwise, `Generate` and `GenerateMulti` return `ErrEmptyCharset`.

### Errors

The library exposes two sentinel errors that can be checked with `errors.Is`:

```go
var (
    tPassword.ErrEmptyCharset  // no character set selected
    tPassword.ErrInvalidLength // length <= 0
)
```

## License
[MIT](LICENSE)