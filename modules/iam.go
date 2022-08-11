package modules

import (
	iam "github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/jsii-runtime-go"
)

func (s Module) IamRole(props iam.RoleProps) iam.Role {
	role := iam.NewRole(s.NestedStack, jsii.String(*props.RoleName), &props)

	return role
}
