# Business Rule <a href="https://gitpod.io/#https://github.com/gouniverse/business-rule" style="float:right:"><img src="https://gitpod.io/button/open-in-gitpod.svg" alt="Open in Gitpod" loading="lazy"></a>

[![Tests Status](https://github.com/gouniverse/rule/actions/workflows/tests.yml/badge.svg?branch=main)](https://github.com/gouniverse/rule/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/gouniverse/rule)](https://goreportcard.com/report/github.com/gouniverse/rule)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/gouniverse/rule)](https://pkg.go.dev/github.com/gouniverse/rule)

A business rule defines or constrains some aspect of business. Given a specified context (data) a business rule always resolves to either true or false.

Business rules are intended to assert business structure or to control or influence the behavior of the business. Business rules describe the operations, definitions and constraints that apply to an organization. Business rules can apply to people, processes, corporate behavior and computing systems in an organization, and are put in place to help the organization achieve its goals.


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

### 1) Direct usage

```php
rule = rule.Rule{}

// Your data. The given context for this rule
rule.SetContext("Hello World")

// Your condition. When will the condition be met
rule.SetCondition(func (ctx any){
    value = ctx.(string)
    return value == "Hello world"
});

if (rule.Fails()) {
    // Then Execute fail logic
}

if (rule.Passes()) {
    // Then Execute pass logic
}
```

### 2) Your Custom Rule Type

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

## More About Business Rules

- Source: https://en.wikipedia.org/wiki/Business_rule 

Business rules can be classified into three main types:

- Structural rules define the relationships between different entities in a business, such as customers, products, and orders.
- Decision rules determine how the business should respond to certain events, such as a customer making a purchase or a product becoming out of stock.
- Process flow rules specify the steps that must be taken to complete a business process, such as the steps involved in processing a sales order.

Business rules are important for a number of reasons. They help to ensure that the business operates in a consistent and efficient manner. They also help to prevent errors and fraud. Additionally, business rules can be used to automate business processes, which can save time and money.

The benefits of business rules include:

- Consistency: Business rules ensure that the business operates in a consistent manner, regardless of who is performing the task. This is important for maintaining high-quality customer service and reducing errors.
- Efficiency: Business rules can be used to automate business processes, which can save time and money. For example, a business rule can be used to automatically approve sales orders that meet certain criteria.
- Compliance: Business rules can be used to ensure that the business complies with regulations. For example, a business rule can be used to verify that customers are of legal age to purchase alcohol.
- Flexibility: Business rules can be easily changed as the business needs change. This is important for businesses that are constantly evolving.

The challenges of business rules include:

- Complexity: Business rules can be complex, especially for large businesses with complex operations. This can make it difficult to manage and maintain business rules.
- Documentation: It is important to document business rules in a clear and concise manner so that they can be easily understood by everyone who needs to use them.
- Testing: Business rules should be tested thoroughly to ensure that they work as expected. This can be a time-consuming and expensive process.
- Change management: Business rules need to be managed effectively as the business changes. This includes updating the rules when necessary and communicating the changes to everyone who needs to know about them.

Overall, business rules are an important tool for businesses of all sizes. They can help to improve efficiency, compliance, and flexibility. However, it is important to manage business rules effectively to avoid the challenges associated with them.

