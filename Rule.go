package rule

type Rule struct {
	// Context (data) for the rule
	context any

	// Condition function
	condition func(context any) bool
}

var _ RuleInterface = (*Rule)(nil) // verify it extends the RuleInterface interface

func NewRule() *Rule {
	return &Rule{}
}

// Sets the context/context for the rule
func (rule *Rule) SetContext(context any) {
	rule.context = context
}

func (rule *Rule) GetContext() any {
	return rule.context
}

// Sets the condition for the rule
func (rule *Rule) SetCondition(condition func(context any) bool) {
	rule.condition = condition
}

// Returns true if the rule evaluates to false
func (rule *Rule) Fails() bool {
	return !rule.Evaluate()
}

// Returns true if the rule evaluates to true
func (rule *Rule) Passes() bool {
	return rule.Evaluate()
}

// Evaluates the condition and returns the result
func (rule *Rule) Evaluate() bool {
	return rule.condition(rule.context)
}
