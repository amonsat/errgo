# errgo

Simple go lib for checking errors.

## Instalation

```go
    go get github.com/amonsat/errgo
```

## Usage

```go
package main

import "github.com/amonsat/errgo"

func main() {
    check := errgo.NewErrorCheck(func(err error) { fmt.Println("error:", err) })
    check(errFunc())
}

func errFunc() error {
    return fmt.Errorf("some error")
}
```
