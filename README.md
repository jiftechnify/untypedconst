# untypedconst

A Go linter checks that [untyped constant expressions](https://go.dev/blog/constants) are used as values of [defined type](https://go.dev/ref/spec#Type_definitions) (a.k.a. named type).

## Usage

```bash
go install github.com/jiftechnify/untypedconst/cmd/untypedconst@latest
untypedconst ./...
```

## Motivation

When you want to define enums in Go, you should go through following steps:

1. Define an "defined type" for the enum.
2. Define enum variants as values of the type you defined in 1.

Or, you should write a code like this:

```go
type TrafficLight int

const (
    None TrafficLight = iota // 0
    Green                    // 1
    Yellow                   // 2
    Red                      // 3
)
```

Then you use it like:

```go
func next(tl TrafficLight) TrafficLight {
    switch tl {
    case Green:
        return Yellow
    case Yellow:
        return Red
    case Red:
        return Green
    case None:
        return None
    default:
        panic("boom!")
    }
}

func good() {
    tl := next(Green) // tl == Yellow
}
```

So far so good, until you write a code like this by accident:

```go
func bad() {
    tl := next(4) // it compiles, unfortunately 
    // => boom!
}
```

In the code above, `4` is an **"untyped constant"**, and when you "use" it (assign it to a variable, pass it to a function, etc.), it is assigned the most suited type from the context, i.e. `TrafficLight`. So types match and **it compiles despite the actual value is out of the range of `TrafficLight`!** It's a pity that Go compiler doesn't come with a feature that prevents us from breaking codes in such a way.

**untypedconst** (indirectly) detects such a mistake by checking whether you don't use untyped constants in places that require values of some defined type:

```go
func bad() {
    tl := next(4)
    // => passing untyped constant to parameter of defined type "TrafficLight"
}
```

## License

MIT
