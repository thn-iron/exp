// Package exp implements a binary expression tree which can be used to evaluate
// arbitrary binary expressions. You can use this package to build your own
// expressions however a few expressions are provided out of the box.
package exp

// The Exp interface represents a tree node. There are several implementations
// of the interface in this package, but one may define custom Exp's as long as
// they implement the Eval function.
type Exp interface {
	Eval(Params) bool
}

// Params defines the interface needed by Exp in order to be able to validate
// conditions. An example implementation of this interface would be
// https://golang.org/pkg/net/url/#Values.
//
// A simple implementation of Params can be described as a map of strings.
//
// 	type Params map[string]string
//
// 	func (p Params) Get(s string) string {
// 		return p[s]
// 	}
type Params interface {
	Get(string) string
}

// Map is a simple implementation of Params using a map of strings.
type Map map[string]string

// Get returns the value pointed to by key.
func (m Map) Get(key string) string {
	return m[key]
}

// And

type expAnd struct{ elems []Exp }

func (a expAnd) Eval(p Params) bool {
	for _, elem := range a.elems {
		if !elem.Eval(p) {
			return false
		}
	}
	return true
}

// And evaluates to true if all t's are true.
func And(t ...Exp) Exp {
	return expAnd{t}
}

// Or

type expOr struct{ elems []Exp }

func (o expOr) Eval(p Params) bool {
	for _, elem := range o.elems {
		if elem.Eval(p) {
			return true
		}
	}
	return false
}

// Or evaluates to true if any t's are true.
func Or(t ...Exp) Exp {
	return expOr{t}
}

// Not

type expNot struct{ elem Exp }

func (n expNot) Eval(p Params) bool {
	return !n.elem.Eval(p)
}

// Not evaluates to the opposite of t.
func Not(t Exp) Exp {
	return expNot{t}
}