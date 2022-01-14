package main

import (
	"log"
)

type BusinessRule struct {
	Context interface{}
}

func (r BusinessRule) Condition() bool {
	return true
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

func main() {
	rule := BusinessRule{}
	log.Println(rule.Passes())
}
