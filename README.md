# log4g

[![license card](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg?label=license)](https://github.com/acmestack/log4g/blob/main/LICENSE)
[![go version](https://img.shields.io/github/go-mod/go-version/acmestack/log4g)](#)
[![go report](https://goreportcard.com/badge/github.com/acmestack/log4g)](https://goreportcard.com/report/github.com/acmestack/log4g)
[![codecov report](https://codecov.io/gh/acmestack/log4g/branch/main/graph/badge.svg)](https://codecov.io/gh/acmestack/log4g)
[![workflow](https://github.com/acmestack/log4g/actions/workflows/go.yml/badge.svg?event=push)](#)
[![lasted release](https://img.shields.io/github/v/release/acmestack/log4g?label=lasted)](https://github.com/acmestack/log4g/releases)

## How To Use?

```
go get -u github.com/acmestack/log4g
```

## Import

```
import "github.com/acmestack/log4g/log"
```

## Log Example

```
log.SetLevel(log.TraceLevel)
defer log.Reset()

log.Trace("a", "=", "1")
log.Tracef("a=%d", 1)

log.Trace(func() []interface{} {
    return log.T("a", "=", "1")
})

log.Tracef("a=%d", func() []interface{} {
    return log.T(1)
})

...
```

```
ctx := context.WithValue(context.TODO(), traceIDKey, "0689")

log.SetLevel(log.TraceLevel)
log.SetOutput(myOutput)
defer log.Reset()

logger := log.Ctx(ctx)
logger.Trace("level:", "trace")
logger.Tracef("level:%s", "trace")

...

logger = log.Tag("__in")
logger.Ctx(ctx).Trace("level:", "trace")
logger.Ctx(ctx).Tracef("level:%s", "trace")

...
```

## Color Example

```
fmt.Println(color.BgBlack.Sprint("ok"))
fmt.Println(color.BgRed.Sprint("ok"))
fmt.Println(color.BgGreen.Sprint("ok"))
fmt.Println(color.BgYellow.Sprint("ok"))

attributes := []color.Attribute{
    color.Bold,
    color.Italic,
    color.Underline,
    color.ReverseVideo,
    color.CrossedOut,
    color.Red,
    color.BgGreen,
}

fmt.Println(color.NewText(attributes...).Sprint("ok"))
```

## Assert Example

```
assert.True(t, true)
assert.Nil(t, nil)
assert.Equal(t, 0, "0")
assert.NotEqual(t, "0", 0)
assert.Same(t, 0, "0")
assert.NotSame(t, "0", "0")
assert.Panic(g, func() {}, "an error")
assert.Matches(g, "there's no error", "an error")
assert.Error(g, errors.New("there's no error"), "an error")
assert.TypeOf(g, new(int), (*int)(nil))
assert.Implements(g, errors.New("error"), (*error)(nil))
```
