package config

import "fmt"

type DoSomethingOption struct {
	a string
	b int
	c bool
}

type OptionFunc func(*DoSomethingOption)

func WithB(b int) OptionFunc {
	return func(o *DoSomethingOption) {
		o.b = b
	}
}

func WithC(c bool) OptionFunc {
	return func(o *DoSomethingOption) {
		o.c = c
	}
}

const defaultValueB = 100

func NewDoSomethingOption(a string, opts ...OptionFunc) *DoSomethingOption {
	o := &DoSomethingOption{a: a, b: defaultValueB}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

var nds = NewDoSomethingOption("happy", WithB(3), WithC(false))

// interface variate
type doSomethingOption struct {
	a string
	b int
	c bool
}

type IOption interface {
	apply(*doSomethingOption)
}

type funcOption struct {
	f func(*doSomethingOption)
}

func (fo *funcOption) apply(o *doSomethingOption) {
	fo.f(o)
}

func newFuncOption(f func(option *doSomethingOption)) IOption {
	return &funcOption{
		f: f,
	}
}

func IWithB(b int) IOption {
	return newFuncOption(func(o *doSomethingOption) {
		o.b = b
	})
}

func IWithC(c bool) IOption {
	return newFuncOption(func(o *doSomethingOption) {
		o.c = c
	})
}

func DoSomething(a string, opts ...IOption) {
	o := &doSomethingOption{a: a}
	for _, opt := range opts {
		opt.apply(o)
	}
	fmt.Printf("o:%#v\n", o)
}
