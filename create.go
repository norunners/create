// Package create provides a generic option pattern for creating new values of any type.
package create

// Builder models a way to build values of any type.
// Type parameter T is the type to build, e.g. *MyStruct.
// Type parameter B is the type of the Builder, e.g. *MyStructBuilder.
type Builder[T, B any] interface {
	// Default creates a new Builder value of type B.
	// This is where sensible defaults are defined for type B.
	// The zero value of B must be receivable by Default, e.g. nil pointer receiver.
	Default() B
	// Build creates a new value of type T.
	// An error can be returned on failures, e.g. validation.
	Build() (T, error)
}

// New creates a new value of any type using the option pattern.
// Type parameter T is the type to create, e.g. *MyStruct.
// Type parameter B is the type of the Builder, e.g. *MyStructBuilder.
// Type parameter O is the type of option function, e.g. MyStructOption or func(*MyStructBuilder).
// Options are applied to Builder type B instead of type T.
// The logic of options must be limited to trivial assignments.
// Error handling must be performed by the Builder.Build method when needed.
func New[T any, B Builder[T, B], O ~func(B)](options ...O) (T, error) {
	b := (*new(B)).Default()
	for _, option := range options {
		option(b)
	}
	return b.Build()
}
