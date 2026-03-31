package rule

import "testing"

func TestRule(t *testing.T) {
	rule := New().
		SetContext("hello").
		SetCondition(func(ctx any) bool {
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
	*Rule
}

func NewIsNegativeRule() *isNegativeRule {
	rule := &isNegativeRule{Rule: New()}
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

	isNegative := NewIsNegativeRule().SetContext(5)

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

func TestRuleValidate(t *testing.T) {
	rule := New().
		SetContext("hello").
		SetCondition(func(ctx any) bool {
			s := ctx.(string)
			return s == "hello"
		})

	passed, messages := rule.Validate()
	if !passed {
		t.Fatalf("Validate should return passed=true")
	}
	if len(messages) != 0 {
		t.Fatalf("Validate should return empty messages when no pass messages set")
	}

	rule.AddPassMessage("Context is hello")
	passed, messages = rule.Validate()
	if !passed {
		t.Fatalf("Validate should return passed=true")
	}
	if len(messages) != 1 || messages[0] != "Context is hello" {
		t.Fatalf("Validate should return pass messages")
	}
}

func TestRuleValidateFail(t *testing.T) {
	rule := New().
		SetContext("world").
		SetCondition(func(ctx any) bool {
			s := ctx.(string)
			return s == "hello"
		})

	passed, messages := rule.Validate()
	if passed {
		t.Fatalf("Validate should return passed=false")
	}
	if len(messages) != 0 {
		t.Fatalf("Validate should return empty messages when no fail messages set")
	}

	rule.AddFailMessage("Context is not hello")
	passed, messages = rule.Validate()
	if passed {
		t.Fatalf("Validate should return passed=false")
	}
	if len(messages) != 1 || messages[0] != "Context is not hello" {
		t.Fatalf("Validate should return fail messages")
	}
}

func TestRuleNilCondition(t *testing.T) {
	rule := New().SetContext("test")
	// No condition set

	if rule.Passes() {
		t.Fatalf("Rule with nil condition should not pass")
	}
	if !rule.Fails() {
		t.Fatalf("Rule with nil condition should fail")
	}
}

func TestRuleMessages(t *testing.T) {
	rule := New()

	// Test fail messages
	rule.AddFailMessage("Error 1")
	rule.AddFailMessage("Error 2")

	if len(rule.FailMessages()) != 2 {
		t.Fatalf("Should have 2 fail messages")
	}
	if rule.FailMessageFirst() != "Error 1" {
		t.Fatalf("First fail message should be 'Error 1'")
	}
	if rule.FailMessageLast() != "Error 2" {
		t.Fatalf("Last fail message should be 'Error 2'")
	}

	// Test pass messages
	rule.AddPassMessage("Success 1")
	rule.AddPassMessage("Success 2")

	if len(rule.PassMessages()) != 2 {
		t.Fatalf("Should have 2 pass messages")
	}
	if rule.PassMessageFirst() != "Success 1" {
		t.Fatalf("First pass message should be 'Success 1'")
	}
	if rule.PassMessageLast() != "Success 2" {
		t.Fatalf("Last pass message should be 'Success 2'")
	}
}

func TestRuleEmptyMessages(t *testing.T) {
	rule := New()

	if rule.FailMessageFirst() != "" {
		t.Fatalf("Empty fail messages should return empty string for First")
	}
	if rule.FailMessageLast() != "" {
		t.Fatalf("Empty fail messages should return empty string for Last")
	}
	if rule.PassMessageFirst() != "" {
		t.Fatalf("Empty pass messages should return empty string for First")
	}
	if rule.PassMessageLast() != "" {
		t.Fatalf("Empty pass messages should return empty string for Last")
	}
}
