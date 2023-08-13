# Business Rule <a href="https://gitpod.io/#https://github.com/gouniverse/business-rule" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

[![Tests Status](https://github.com/gouniverse/rule/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/gouniverse/rule/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouniverse/rule)](https://goreportcard.com/report/github.com/gouniverse/rule)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gouniverse/rule)](https://pkg.go.dev/github.com/gouniverse/rule)

A business rule defines or constrains some aspect of business. Given a specified context (data) a business rule always resolves to either true or false.

## Formal specification: 

<b>Given</b> {context} <b>When</b> {condition(s)} <b>Then</b> {pass} <b>Or</b> {fail}

## Specification

The syntax Given {context} When {condition(s)} Then {pass} Or {fail} is a common way to express business rules. It is used to define a set of conditions that must be met in order for a business process to be successful.

The Given clause specifies the context of the rule. This could be a specific event, such as a customer making a purchase, or a general condition, such as the availability of a product.

The When clause specifies the conditions that must be met. These conditions can be anything from the customer having a valid credit card to the product being in stock.

The Then clause specifies the outcome of the rule. If all of the conditions are met, the business process will pass. If any of the conditions are not met, the business process will fail.

Here is an example of a business rule expressed in this syntax:

```
Given a customer makes a purchase
When the customer has a valid credit card
Then the purchase will pass
```

This rule states that a purchase will only be successful if the customer has a valid credit card. If the customer does not have a valid credit card, the purchase will fail.

Business rules expressed in this syntax can be used to automate business processes and ensure that they are consistent and efficient. They can also be used to prevent errors and fraud.


## Install ##

```go
go get github.com/gouniverse/rule
```

## Usage ##

1) Direct usage

```php
rule = rule.Rule{}
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

2. Your Custom Rule Type

Direct usage may not be very practical if we want to reuse on multiple places and avoid duplication of business logic.

Which is why it is recommended to create your own custom rule type


```go

// 1. Specify the business rule class
type AllowAccessRule {
    rule.Rule
}

func NewAllowAccessRule(user models.User) AllowAccessRule {
	rule := AllowAccessRule{}
    rule.SetContext(user)
	rule.SetCondition(rule.allowAccessCondition)
	return rule
}

var _ rule.RuleInterface = (*AllowAccessRule)(nil) // (optional) verify it extends the RuleInterface interface


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
