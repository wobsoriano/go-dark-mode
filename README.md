# go-dark-mode

Control the macOS dark mode in Go. Requires macOS 10.10 or later.

## Install

```bash
$ go get github.com/wobsoriano/go-dark-mode
```

## Usage

```go
package main

import (
	"fmt"

	darkmode "github.com/wobsoriano/go-dark-mode"
)

func main() {
	darkmode.Enable()

  darkmode.Disable()

  darkmode.Toggle()

  v := darkMode.IsEnabled()
  fmt.Printf("is dark mode: %s", v)
}
```

## License

MIT
