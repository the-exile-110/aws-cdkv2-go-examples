package modules

import (
	"app/models"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	s3 "github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/jsii-runtime-go"
)

func (s ModuleService) Vpc(props models.Vpc) ec2.Vpc {

	vpc := ec2.NewVpc(s.NestedStack, jsii.String(*props.VpcName), &props.VpcProps)

	if props.UseVpcFlowLogToS3 {
		bucket := s3.NewBucket(s.NestedStack, jsii.String("MyBlockedBucket"), &s3.BucketProps{
			BlockPublicAccess: s3.BlockPublicAccess_BLOCK_ALL(),
		})

		vpc.AddFlowLog(jsii.String("FlowLogS3"), &ec2.FlowLogOptions{
			Destination: ec2.FlowLogDestination_ToS3(bucket, jsii.String("vpc-flow-logs")),
		})
	}

	if props.UseVpcFlowLogToCloudWatch {

		vpc.AddFlowLog(jsii.String("FlowLogCloudWatch"), &ec2.FlowLogOptions{
			TrafficType: ec2.FlowLogTrafficType_REJECT,
		})
	}

	return vpc
}
