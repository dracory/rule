package rule

import "testing"

func TestRule(t *testing.T) {
	rule := Rule{}

	rule.SetContext("hello")

	rule.SetCondition(func(ctx any) bool {
		s := ctx.(string)

		return s == "hello"
	})

	if rule.Fails() {
		t.Fatalf("On Fails. Rule should not fail")
	}

	if !rule.Passes() {
		t.Fatalf("On Passes. Rule should pass")
	}
}

type isNegativeRule struct {
	Rule
}

func NewIsNegativeRule() *isNegativeRule {
	rule := &isNegativeRule{}
	rule.SetCondition(rule.condition)
	return rule
}

var _ RuleInterface = (*isNegativeRule)(nil) // verify it extends the RuleInterface interface

func (rule isNegativeRule) condition(context any) bool {
	intVar := context.(int)
	return intVar < 0
}

func TestRuleInheritance(t *testing.T) {
	// isNegative := isNegativeRule{}

	// isNegative.SetCondition(func(context any) bool {
	// 	intVar := context.(int)

	// 	return intVar < 0
	// })

	isNegative := NewIsNegativeRule()

	isNegative.SetContext(5)

	if !isNegative.Fails() {
		t.Fatalf("Rule should fail")
	}

	if isNegative.Passes() {
		t.Fatalf("Rule should not pass")
	}

	isNegative.SetContext(-5)

	if isNegative.Fails() {
		t.Fatalf("Rule should not fail")
	}
	if !isNegative.Passes() {
		t.Fatalf("Rule should not pass")
	}
}
