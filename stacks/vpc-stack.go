package stacks

import (
	"app/env"
	"app/modules"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/constructs-go/constructs/v10"
)

func VpcStack(scope constructs.Construct, id string) awsec2.Vpc {
	nestedStack := awscdk.NewNestedStack(scope, &id, &env.NestedStackProps.NestedStackProps)

	vpcName := "DemoVpc"
	vpcCidr := "10.0.0.0/16"
	vpc := modules.Vpc(nestedStack, vpcName, vpcCidr, false, false)

	return vpc
}
