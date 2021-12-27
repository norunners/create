# create

[![Go Reference](https://pkg.go.dev/badge/github.com/norunners/create.svg)](https://pkg.go.dev/github.com/norunners/create)

Package `create` provides a generic option pattern for creating new values of any type.

# Install
```
go get github.com/norunners/create
```
*Requires Go 1.8 or higher.*

# Examples
The `Greeting` type will be created throughout the examples.
```go
type Greeting string
```

### New with options
Create a `Greeting` with `earth` as the noun option.
```go
greeting, err := create.New[Greeting](WithNoun("earth"))
```
> Hello earth!

### New without options
Create a `Greeting` without options so the defaults are used.
```go
greeting, err := create.New[Greeting, *GreetingBuilder]()
```
> Hello world!

### Builder
Defining `GreetingBuilder` as a struct allows fields to be added over time.
```go
type GreetingBuilder struct {
	noun string
}
```

### Builder.Default
The `Default` method provides sensible default values.
```go
func (*GreetingBuilder) Default() *GreetingBuilder {
	return &GreetingBuilder{
		noun: "world",
	}
}
```

### Options
Defining `GreetingOption` allows functional options on the `GreetingBuilder` type.  
Option `WithNoun` assigns a value to the `noun` field, which is not exported.
```go
type GreetingOption func(*GreetingBuilder)

func WithNoun(noun string) GreetingOption {
	return func(b *GreetingBuilder) {
		b.noun = noun
	}
}
```

### Builder.Build
The `Build` method validates the `noun` field and creates a new `Greeting`.
```go
func (b *GreetingBuilder) Build() (Greeting, error) {
	if b.noun == "" {
		return "", fmt.Errorf("empty noun")
	}
	return Greeting(fmt.Sprintf("Hello %s!", b.noun)), nil
}
```

### Instantiation
This instantiates `NewGreeting` from `create.New`.  
All parameterized types are required for instantiation, e.g. no type inference.
```go
var NewGreeting = create.New[Greeting, *GreetingBuilder, GreetingOption]
greeting, err := NewGreeting(...)
```
*This can be useful as a package scoped variable, e.g. `greeting.New`.*

### Satisfy the Builder interface
This ensures `GreetingBuilder` satisfies `create.Builder`.
```go
var _ create.Builder[Greeting, *GreetingBuilder] = (*GreetingBuilder)(nil)
```

# Benefits
1. A single future-proof function to create values.
2. Provide sensible defaults for any type.
3. Override defaults with options.
4. Validate values before creation.
5. Zero dependencies.

# Why?
This is a bit of an experimental exercise of generics in Go
but could also be seen as a standardized way to use the option pattern.

# References
* https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
* https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
* https://www.sohamkamani.com/golang/options-pattern/
* https://golang.cafe/blog/golang-functional-options-pattern.html

## License
* [MIT License](LICENSE)
