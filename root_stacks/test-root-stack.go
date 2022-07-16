package root_stacks

import (
	"app/nested_stacks"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func (s RootStack) TestRootStack() *awscdk.Stack {
	rootStack := awscdk.NewStack(s.Scope, s.Id, s.Env)

	vpcStack := nested_stacks.CreateNestedStack(rootStack, "VpcStack").VpcStack()

	fmt.Println("vpc:", vpcStack.VpcId())

	return &rootStack
}
