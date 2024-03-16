# dummylogger

A dummy logger for a custom package.

## Usage

Use the custom logger in your package.

```go
import (
    "github.com/kaatinga/dummylogger"
)

var log = dummylogger.Get

func Init(logger dummylogger.I) {
    dummylogger.Set(logger)
}
```

Use it where you need it in the package.

```go
func main() {
    log().Errorf("This is an error message")
}
```

Create a custom logger in your project.

```go
type customLogger bool

var CustomLogger = customLogger(true)

func (customLogger) Errorf(format string, v ...interface{}) {
	logger.Error().Msgf(format, v...)
}
```

Pass the custom logger to init the package.

```go
import (
    "github.com/yourname/yourpackage"
)

func main() {
    yourpackage.Init(CustomLogger)
}
```