package rule

type RuleInterface interface {
	SetContext(ctx any) RuleInterface
	GetContext() any
	SetCondition(conditionCallback func(ctx any) bool) RuleInterface
	Passes() bool
	Fails() bool
	Evaluate() bool
	Validate() (passed bool, messages []string)
	AddFailMessage(message string)
	AddPassMessage(message string)
	FailMessages() []string
	PassMessages() []string
	FailMessageFirst() string
	FailMessageLast() string
	PassMessageFirst() string
	PassMessageLast() string
}
