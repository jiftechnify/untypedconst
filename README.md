# untypedconst

A Go linter checks that [untyped constant expressions](https://go.dev/blog/constants) are used as values of [defined type](https://go.dev/ref/spec#Type_definitions) (a.k.a. named type).

## Usage

```bash
# as a standalone linter
go install github.com/jiftechnify/untypedconst/cmd/untypedconst@latest
untypedconst ./...

# with `go vet`
go install github.com/jiftechnify/untypedconst/cmd/untypedconst@latest
go vet -vettool=$(which untypedconst) ./...
```

<details>
<summary>Run as a golangci-lint's custom linter</summary>

### Run as a `golangci-lint`'s custom linter 

You can also run **untypedconst** as a custom linter of [golangci-lint](https://golangci-lint.run/) with the following steps.
Please refer to [the official document](https://golangci-lint.run/docs/plugins/module-plugins/) for details.

1. Create `.custom-gcl.yml` under the root directory of your project:

    ```yaml
    version: v2.5.0
    plugins:
      - module: "github.com/jiftechnify/untypedconst"
        import: "github.com/jiftechnify/untypedconst/cmd/gciplugin"
        version: "latest" # or specify a specific version e.g. "v0.2.0"
    ```

2. Run `golangci-lint custom` to build a custom `golangci-lint` binary with **untypedconst** included as a plugin.
    - By default, the output path of the binary is `./custom-gcl`. You can change it by specifying `destination` field in `.custom-gcl.yml`. 

3. Add lines of configuration to your `.golangci.yml` to enable **untypedconst**:

    ```yaml
    # for golangci-lint v2.x
    version: "2"

    linters:
      enable:
        - untypedconst

      settings:
        custom:
          untypedconst:
            type: "module"
            description: "Detects suspicious usage of untyped constants"
    ```

    ```yaml
    # for golangci-lint v1.x
    linters:
      enable:
        - untypedconst

    linters-settings:
      custom:
        untypedconst:
          type: "module"
          description: "Detects suspicious usage of untyped constants"
    ```

4. Run the custom `golangci-lint` executable against your codebase:

    ```bash
    ./custom-gcl run ./...
    ```
</details>

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
