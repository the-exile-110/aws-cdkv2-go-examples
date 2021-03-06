package modules

import "github.com/aws/aws-cdk-go/awscdk/v2"

type ModuleService struct {
	NestedStack awscdk.NestedStack
}

func CreateModule(nestedStack awscdk.NestedStack) *ModuleService {
	return &ModuleService{
		NestedStack: nestedStack,
	}
}
