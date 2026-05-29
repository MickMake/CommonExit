# CommonExit

`CommonExit` is a tiny Go module that defines one project convention:

```go
for range CommonExit.Block {
    // staged work
    if err != nil {
        break
    }
}
return err
```

It exists for codebases that prefer a **common exit point at the end of a function**, while still wanting simple staged bailout inside the function body.

It is deliberately small. It contains one exported value: `Block`.

## Install

```bash
go get github.com/MickMake/CommonExit
```

## Why this exists

Go commonly encourages early returns:

```go
func Run() error {
    if err := validate(); err != nil {
        return err
    }

    if err := prepare(); err != nil {
        return err
    }

    return execute()
}
```

That is idiomatic and often excellent.

However, some codebases prefer a single visible return path, especially when a function has:

- common cleanup,
- common logging,
- shared error decoration,
- staged setup,
- a desire to make refactoring boundaries visually obvious.

Without `CommonExit.Block`, that often becomes nested conditionals:

```go
func Run() error {
    var err error

    err = validate()
    if err == nil {
        err = prepare()
        if err == nil {
            err = execute()
        }
    }

    return err
}
```

That keeps one exit, but the code starts walking sideways. Enough of that and the function begins to look like it was folded by someone trying to hide a map from a dragon.

`CommonExit.Block` keeps the staged flow flat:

```go
func Run() error {
    var err error

    for range CommonExit.Block {
        err = validate()
        if err != nil {
            break
        }

        err = prepare()
        if err != nil {
            break
        }

        err = execute()
        if err != nil {
            break
        }
    }

    return err
}
```

## What `Block` is

`Block` is a one-element zero-size array:

```go
var Block = [...]struct{}{{}}
```

That means this loop runs exactly once:

```go
for range CommonExit.Block {
    // runs once
}
```

The loop is not for repetition. It is a single-pass block that allows `break` to jump to one common exit point.

## Before and after scenarios

### Scenario 1: avoiding deeply nested conditionals

Before:

```go
func Save() error {
    var err error

    err = validate()
    if err == nil {
        err = openFile()
        if err == nil {
            err = writeFile()
        }
    }

    return err
}
```

After:

```go
func Save() error {
    var err error

    for range CommonExit.Block {
        err = validate()
        if err != nil {
            break
        }

        err = openFile()
        if err != nil {
            break
        }

        err = writeFile()
        if err != nil {
            break
        }
    }

    return err
}
```

### Scenario 2: one final logging point

```go
func Run() error {
    var err error

    for range CommonExit.Block {
        err = loadConfig()
        if err != nil {
            break
        }

        err = runJob()
        if err != nil {
            break
        }
    }

    if err != nil {
        log.Printf("run failed: %v", err)
    }

    return err
}
```

### Scenario 3: a refactoring signal

A common-exit block can act as a human marker:

```go
func Start() error {
    var err error

    for range CommonExit.Block {
        err = loadConfig()
        if err != nil {
            break
        }

        err = connectDatabase()
        if err != nil {
            break
        }

        err = startServer()
        if err != nil {
            break
        }
    }

    return err
}
```

If that block keeps growing, it is a clue that the block wants to become its own function:

```go
func Start() error {
    return startRuntime()
}
```

## Rules of use

Use `CommonExit.Block` when:

- you deliberately want one common exit point,
- the function has staged operations,
- each stage may fail,
- the block may later be extracted into a function.

Avoid it when:

- a simple early return is clearer,
- the function is tiny,
- you actually need a loop,
- you are writing code for a team that strongly prefers standard Go early returns.

Good:

```go
func DoThing() error {
    var err error

    for range CommonExit.Block {
        err = stepOne()
        if err != nil {
            break
        }

        err = stepTwo()
        if err != nil {
            break
        }
    }

    return err
}
```

Avoid mixing common-exit style with surprise returns inside the block:

```go
func DoThing() error {
    var err error

    for range CommonExit.Block {
        if bad {
            return errors.New("bad") // avoid this in a common-exit block
        }

        err = stepTwo()
        if err != nil {
            break
        }
    }

    return err
}
```

## Is this idiomatic Go?

Not especially.

It is a deliberate convention for people who value common-exit readability over strict Go idiom in selected orchestration functions.

Use it sparingly, document the convention, and do not use it as a substitute for good function boundaries. It is a signpost, not a religion. Religions have committees; this has one array.

## Testing

```bash
go test ./...
```
