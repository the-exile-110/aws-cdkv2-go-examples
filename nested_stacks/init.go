package nested_stacks

import (
	"app/env"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type NestedStack struct {
	Scope constructs.Construct
	Id    *string
	Env   *awscdk.NestedStackProps
}

func CreateNestedStack(scope constructs.Construct, id string) *NestedStack {
	return &NestedStack{
		Scope: scope,
		Id:    &id,
		Env:   &env.NestedStackProps.NestedStackProps,
	}
}
