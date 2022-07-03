package services

import (
	"app/env"
	"app/stacks"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/constructs-go/constructs/v10"
)

func TestRootStack(scope constructs.Construct, id string) *awscdk.Stack {
	rootStack := awscdk.NewStack(scope, &id, &env.StackProps.StackProps)

	vpc := stacks.VpcStack(rootStack, "VpcStack")

	fmt.Println("vpc:", vpc.VpcId())

	return &rootStack
}
