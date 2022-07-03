package modules

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	s3 "github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/jsii-runtime-go"
)

func Vpc(stack awscdk.NestedStack, vpcName, cidr string, useVpcFlowLogToS3 bool, useVpcFlowLogToCloudWatch bool) ec2.Vpc {

	vpc := ec2.NewVpc(stack, jsii.String(vpcName), &ec2.VpcProps{
		VpcName:                jsii.String(vpcName),
		Cidr:                   jsii.String(cidr),
		DefaultInstanceTenancy: "",
		EnableDnsHostnames:     nil,
		EnableDnsSupport:       nil,
		FlowLogs:               nil,
		GatewayEndpoints:       nil,
		MaxAzs:                 nil,
		NatGatewayProvider:     nil,
		NatGateways:            nil,
		NatGatewaySubnets:      nil,
		SubnetConfiguration:    nil,
		VpnConnections:         nil,
		VpnGateway:             nil,
		VpnGatewayAsn:          nil,
		VpnRoutePropagation:    nil,
	})

	if useVpcFlowLogToS3 {
		bucket := s3.NewBucket(stack, jsii.String("MyBlockedBucket"), &s3.BucketProps{
			BlockPublicAccess: s3.BlockPublicAccess_BLOCK_ALL(),
		})

		vpc.AddFlowLog(jsii.String("FlowLogS3"), &ec2.FlowLogOptions{
			Destination: ec2.FlowLogDestination_ToS3(bucket, jsii.String("vpc-flow-logs")),
		})
	}

	if useVpcFlowLogToCloudWatch {

		vpc.AddFlowLog(jsii.String("FlowLogCloudWatch"), &ec2.FlowLogOptions{
			TrafficType: ec2.FlowLogTrafficType_REJECT,
		})
	}

	return vpc
}
