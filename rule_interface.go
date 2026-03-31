package rule

type RuleInterface interface {
	SetContext(ctx any) RuleInterface
	GetContext() any
	SetCondition(conditionCallback func(ctx any) bool) RuleInterface
	Passes() bool
	Fails() bool
}
