package nested_stacks

import (
	"app/env"
	"app/models"
	"app/modules"
	"fmt"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/jsii-runtime-go"
)

func (s NestedStack) VpcStack() awsec2.Vpc {
	nestedStack := awscdk.NewNestedStack(s.Scope, s.Id, s.Env)

	vpcName := fmt.Sprintf("%s-vpc", env.ProjectName)
	vpcCidr := "10.0.0.0/16"

	vpcProps := models.Vpc{
		VpcProps: awsec2.VpcProps{
			VpcName:                jsii.String(vpcName),
			Cidr:                   jsii.String(vpcCidr),
			MaxAzs:                 jsii.Number(3),
			DefaultInstanceTenancy: "",
			EnableDnsHostnames:     nil,
			EnableDnsSupport:       nil,
			FlowLogs:               nil,
			GatewayEndpoints:       nil,
			NatGatewayProvider:     nil,
			NatGateways:            jsii.Number(0),
			NatGatewaySubnets:      nil,
			SubnetConfiguration:    nil,
			VpnConnections:         nil,
			VpnGateway:             nil,
			VpnGatewayAsn:          nil,
			VpnRoutePropagation:    nil,
		},
		UseVpcFlowLog:             false,
		UseVpcFlowLogToCloudWatch: false,
		UseVpcFlowLogToS3:         false,
	}
	vpc := modules.CreateModule(nestedStack).Vpc(vpcProps)

	return vpc
}
