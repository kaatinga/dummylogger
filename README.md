# dummylogger

A dummy logger for a custom package.

## Usage

Create a custom logger in your project.

```go
// Custom logger to pass to a package
type customLogger bool

var CustomLogger = customLogger(true)

func (customLogger) Errorf(format string, v ...interface{}) {
	logger.Error().Msgf(format, v...)
}
```

Use the custom logger in your package.

```go
import "github.com/kaatinga/dummylogger"

func (c *CustomType) DoSomething() {
    dummylogger.CustomLogger.Errorf("Something went wrong %w", err)
}
```
