package businessrule

import (
	"errors"
	"log"
)

// convert types take an int and return a string value.
type condition func(interface{}) bool

type BusinessRule struct {
	// Context (data) for the rule
	Context   interface{}
	Condition condition
}

// Sets the context for the rule
func (r BusinessRule) SetContext(context interface{}) {
	r.Context = context
}

// Sets the condition for the rule
func (r BusinessRule) SetCondition(fn condition) {
	log.Println(fn)
	r.Condition = fn
	log.Println(r.Condition)
}

// Evaluate evaulates the condition and returns the result
func (r BusinessRule) Evaluate() (bool, error) {
	if r.Condition == nil {
		return false, errors.New("condition not set")
	}
	if r.Context == nil {
		return false, errors.New("context not set")
	}
	return r.Condition(r.Context), nil
}

// Passes - returns true if the rule evaluates to true
func (r BusinessRule) Passes() (bool, error) {
	return r.Evaluate()
}

// Fails returns true if the rule evaluates to false
func (r BusinessRule) Fails() (bool, error) {
	val, err := r.Evaluate()
	return !val, err
}
