package create_test

import (
	"fmt"
	"testing"

	"github.com/norunners/create"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		options  []GreetingOption
		err      error
		expected Greeting
	}{
		{
			name:     "nil options",
			expected: "Hello world!",
		},
		{
			name:     "zero options",
			options:  []GreetingOption{},
			expected: "Hello world!",
		},
		{
			name:     "with noun",
			options:  []GreetingOption{WithNoun("planet")},
			expected: "Hello planet!",
		},
		{
			name:     "override noun",
			options:  []GreetingOption{WithNoun("planet"), WithNoun("earth")},
			expected: "Hello earth!",
		},
		{
			name:    "empty noun",
			options: []GreetingOption{WithNoun("")},
			err:     fmt.Errorf("empty noun"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			greeting, err := create.New[Greeting](test.options...)
			if ex := test.err; ex != nil {
				if acError := fmt.Sprintf("%v", err); acError != ex.Error() {
					t.Fatalf("expected error: %v but found: %v", test.err, err)
				}
			}
			if greeting != test.expected {
				t.Errorf("expected greeting: %s but found: %s", test.expected, greeting)
			}
		})
	}
}

func TestInstantiatedNew(t *testing.T) {
	tests := []struct {
		name     string
		options  []GreetingOption
		err      error
		expected Greeting
	}{
		{
			name:     "nil options",
			expected: "Hello world!",
		},
		{
			name:     "zero options",
			options:  []GreetingOption{},
			expected: "Hello world!",
		},
		{
			name:     "with noun",
			options:  []GreetingOption{WithNoun("planet")},
			expected: "Hello planet!",
		},
		{
			name:     "override noun",
			options:  []GreetingOption{WithNoun("planet"), WithNoun("earth")},
			expected: "Hello earth!",
		},
		{
			name:    "empty noun",
			options: []GreetingOption{WithNoun("")},
			err:     fmt.Errorf("empty noun"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			greeting, err := NewGreeting(test.options...)
			if ex := test.err; ex != nil {
				if acError := fmt.Sprintf("%v", err); acError != ex.Error() {
					t.Fatalf("expected error: %v but found: %v", test.err, err)
				}
			}
			if greeting != test.expected {
				t.Errorf("expected greeting: %s but found: %s", test.expected, greeting)
			}
		})
	}
}

func TestNewOptionsZero(t *testing.T) {
	greeting, err := create.New[Greeting, *GreetingBuilder]()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ex := Greeting("Hello world!"); greeting != ex {
		t.Errorf("expected greeting: %s but found: %s", ex, greeting)
	}
}

func TestNewOptionsSingle(t *testing.T) {
	greeting, err := create.New[Greeting](WithNoun("option"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ex := Greeting("Hello option!"); greeting != ex {
		t.Errorf("expected greeting: %s but found: %s", ex, greeting)
	}
}

func TestNewOptionsFunc(t *testing.T) {
	greeting, err := create.New[Greeting](func(b *GreetingBuilder) {
		b.noun = "func"
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if ex := Greeting("Hello func!"); greeting != ex {
		t.Errorf("expected greeting: %s but found: %s", ex, greeting)
	}
}

type Greeting string

// NewGreeting is an instantiation of create.New.
var NewGreeting = create.New[Greeting, *GreetingBuilder, GreetingOption]

// Ensures GreetingBuilder satisfies create.Builder.
var _ create.Builder[Greeting, *GreetingBuilder] = (*GreetingBuilder)(nil)

type GreetingBuilder struct {
	noun string
}

func (*GreetingBuilder) Default() *GreetingBuilder {
	return &GreetingBuilder{
		noun: "world",
	}
}

type GreetingOption func(*GreetingBuilder)

func WithNoun(noun string) GreetingOption {
	return func(b *GreetingBuilder) {
		b.noun = noun
	}
}

func (b *GreetingBuilder) Build() (Greeting, error) {
	if b.noun == "" {
		return "", fmt.Errorf("empty noun")
	}
	return Greeting(fmt.Sprintf("Hello %s!", b.noun)), nil
}
