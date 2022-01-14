package main

import (
	"log"
)

// convert types take an int and return a string value.
type condition func() bool

type BusinessRule struct {
	Context interface{}
}

func (r BusinessRule) Condition(fn condition) bool {
	return fn()
}

// Evaluate evaulates the condition and returns the result
func (r BusinessRule) Evaluate() bool {
	return true
}

// Passes - returns true if the rule evaluates to true
func (r BusinessRule) Passes() bool {
	return r.Evaluate()
}

// Fails returns true if the rule evaluates to false
func (r BusinessRule) Fails() bool {
	return !r.Evaluate()
}

type NegativeRule struct {
	BusinessRule
	Context interface{}
}

func (r NegativeRule) Condition() bool {
	i := r.Context.(int64)

	if i < 0 {
		return true
	}

	return false
}

func main() {
	rule := BusinessRule{
		Context: "hello",
	}
	rule.Condition(func() bool {
		s := rule.Context.(string)
		if s == "hello" {
			return true
		}
		return false
	})

	negativeRule := NegativeRule{
		Context: 5,
	}

	log.Println(negativeRule.Fails())
	log.Println(rule.Passes())
}
