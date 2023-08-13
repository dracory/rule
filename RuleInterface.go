package rule

type RuleInterface interface {
	SetContext(ctx any)
	GetContext() any
	SetCondition(conditionCallback func(ctx any) bool)
	Passes() bool
	Fails() bool
}
