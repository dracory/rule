package businessrule

import (
	"testing"
)

func TestBusinessRule(t *testing.T) {
	rule := BusinessRule{
		Context: "hello",
		Condition: func(context interface{}) bool {
			s := context.(string)
			if s == "hello" {
				return true
			}
			return false
		},
	}
	//rule.SetContext("world")

	// rule.SetCondition(func(context interface{}) bool {
	// 	s := context.(string)
	// 	if s == "hello" {
	// 		return true
	// 	}
	// 	return false
	// })

	// log.Println(rule.Condition)
	// log.Println(rule.Context)

	fails, err := rule.Fails()

	if err != nil {
		t.Fatalf("On Fails. Unexpected errors thrown '%s'", err.Error())
	}

	if fails {
		t.Fatalf("On Fails. Rule should not fail")
	}

	hasPassed, err := rule.Passes()

	if err != nil {
		t.Fatalf("On Passes. Unexpected errors thrown %s", err.Error())
	}

	if !hasPassed {
		t.Fatalf("On Passes. Rule should pass")
	}
}

// func TestBusinessRuleInheritance(t *testing.T) {
// 	type isNegativeRule struct {
// 		BusinessRule
// 		// Context int64
// 	}

// 	// func (r isNegativeRule) Condition() bool {
// 	// 	i := r.Context

// 	// 	if i < 0 {
// 	// 		return true
// 	// 	}

// 	// 	return false
// 	// }

// 	negativeRule := isNegativeRule{}
// 	negativeRule.SetCondition(func(i interface{}) bool {
// 		intVar := i.(int)

// 		if intVar < 0 {
// 			return true
// 		}

// 		return false
// 	})
// 	negativeRule.SetContext(5)

// 	log.Println(negativeRule.Context)
// 	if fails, err := negativeRule.Fails(); fails {
// 		if err != nil {
// 			t.Fatalf("No errors should be returned. but error '%s'", err.Error())
// 		}
// 		t.Fatalf("Rule should not fail")
// 	}
// 	// if !negativeRule.Passes() {
// 	// 	t.Fatalf("Rule should not pass")
// 	// }

// 	// negativeRule.SetContext(-5)

// 	// log.Println(negativeRule.Context)

// 	// if negativeRule.Fails() {
// 	// 	t.Fatalf("Rule should not fail")
// 	// }
// 	// if !negativeRule.Passes() {
// 	// 	t.Fatalf("Rule should not pass")
// 	// }
// }
