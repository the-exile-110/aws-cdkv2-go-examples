package root_stacks

import (
	"app/env"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

type RootStack struct {
	Scope constructs.Construct
	Id    *string
	Env   *awscdk.StackProps
}

func CreateRootStack(scope constructs.Construct, id string) *RootStack {
	return &RootStack{
		Scope: scope,
		Id:    &id,
		Env:   &env.StackProps.StackProps,
	}
}
