package modules

import "github.com/aws/aws-cdk-go/awscdk/v2"

type Module struct {
	NestedStack awscdk.NestedStack
}

func NewModule(nestedStack awscdk.NestedStack) *Module {
	return &Module{
		NestedStack: nestedStack,
	}
}
