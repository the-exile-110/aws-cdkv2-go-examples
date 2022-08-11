package modules

import (
	"app/models"
	"fmt"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	s3 "github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/jsii-runtime-go"
)

func (s Module) Vpc(props models.Vpc) ec2.Vpc {

	vpc := ec2.NewVpc(s.NestedStack, jsii.String(*props.VpcName), &props.VpcProps)

	if props.UseVpcFlowLogToS3 {
		bucketName := fmt.Sprintf("%s-flow-logs-bucket", *props.VpcName)
		bucket := s3.NewBucket(s.NestedStack, jsii.String(bucketName), &s3.BucketProps{
			BucketName:        jsii.String(bucketName),
			BlockPublicAccess: s3.BlockPublicAccess_BLOCK_ALL(),
		})

		vpc.AddFlowLog(jsii.String("FlowLogS3"), &ec2.FlowLogOptions{
			Destination: ec2.FlowLogDestination_ToS3(bucket, jsii.String(fmt.Sprintf("%s-flow-logs", *props.VpcName))),
		})
	}

	if props.UseVpcFlowLogToCloudWatch {
		vpc.AddFlowLog(jsii.String("FlowLogCloudWatch"), &ec2.FlowLogOptions{
			TrafficType: ec2.FlowLogTrafficType_REJECT,
		})
	}

	return vpc
}
