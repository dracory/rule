package rule

type Rule struct {
	// Context (data) for the rule
	context any

	// Condition function
	condition func(context any) bool

	// Optional extras

	// helper array to hold fail messages
	failMessages []string

	// helper array to hold pass messages
	passMessages []string
}

var _ RuleInterface = (*Rule)(nil) // verify it extends the RuleInterface interface

// NewRule creates a new rule
// Deprecated use New as it is shorter
func NewRule() *Rule {
	return &Rule{}
}

// New creates a new rule
func New() *Rule {
	return &Rule{}
}

// Sets the context/context for the rule
func (rule *Rule) SetContext(context any) RuleInterface {
	rule.context = context
	return rule
}

func (rule *Rule) GetContext() any {
	return rule.context
}

// Sets the condition for the rule
func (rule *Rule) SetCondition(condition func(context any) bool) RuleInterface {
	rule.condition = condition
	return rule
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

// AddFailMessage adds a fail message
func (rule *Rule) AddFailMessage(message string) {
	rule.failMessages = append(rule.failMessages, message)
}

// AddPassMessage adds a pass message
func (rule *Rule) AddPassMessage(message string) {
	rule.passMessages = append(rule.passMessages, message)
}

// FailMessages retrieves the fail messages
func (rule *Rule) FailMessages() []string {
	return rule.failMessages
}

// PassMessages retrieves the pass messages
func (rule *Rule) PassMessages() []string {
	return rule.passMessages
}

// FailMessageFirst retrieves the first fail message
func (rule *Rule) FailMessageFirst() string {
	if len(rule.failMessages) < 1 {
		return ""
	}
	return rule.failMessages[0]
}

// FailMessageLast retrieves the last fail message
func (rule *Rule) FailMessageLast() string {
	if len(rule.failMessages) < 1 {
		return ""
	}
	return rule.failMessages[len(rule.failMessages)-1]
}

// PassMessageFirst retrieves the first pass message
func (rule *Rule) PassMessageFirst() string {
	if len(rule.passMessages) < 1 {
		return ""
	}
	return rule.passMessages[0]
}

// PassMessageLast retrieves the last pass message
func (rule *Rule) PassMessageLast() string {
	if len(rule.passMessages) < 1 {
		return ""
	}
	return rule.passMessages[len(rule.passMessages)-1]
}
