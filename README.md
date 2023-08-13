# Business Rule <a href="https://gitpod.io/#https://github.com/gouniverse/business-rule" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

[![Tests Status](https://github.com/gouniverse/rule/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/gouniverse/rule/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouniverse/rule)](https://goreportcard.com/report/github.com/gouniverse/rule)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gouniverse/rule)](https://pkg.go.dev/github.com/gouniverse/rule)

A business rule defines or constrains some aspect of business. Given a specified context (data) a business rule always resolves to either true or false.

Formal specification: // Given {context} When {condition(s)} Then {pass} Or {fail}

## Usage ##

1) Direct usage

```php
rule = Rule{}
rule.SetContext("Hello World")
rule.SetCondition(func (ctx any){
    value = ctx.(string)
    return value == "Hello world"
});

if (rule.Fails()) {
    // Execute fail logic
}

if (rule.Passes()) {
    // Execute pass logic
}
```

2. Separate Rule Type

Direct usage may not be very practical if we want to reuse on multiple places and avoid duplication of business logic.

Which is why it is better to create our own rule type


```go

// 1. Specify the business rule class
type AllowAccessRule {
    Rule
}

func NewAllowAccessRule(user models.User) AllowAccessRule {
	rule := AllowAccessRule{}
    rule.SetContext(user)
	rule.SetCondition(rule.allowAccessCondition)
	return rule
}

var _ RuleInterface = (*AllowAccessRule)(nil) // verify it extends the RuleInterface interface


func (rule *AllowAccessRule) allowAccessCondition(context any) bool {
	user := context.(models.User)
	return user.IsEmailConfirmed() && user.IsActive()
}
```

Now lets use it where we need it

```go
if NewAllowAccessRule(user).Fails() {
    panic('You are not allowed access to this part of the website');
}
````
