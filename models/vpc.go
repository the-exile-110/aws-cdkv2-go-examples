package models

import ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"

type Vpc struct {
	ec2.VpcProps
	UseVpcFlowLog             bool
	UseVpcFlowLogToS3         bool
	UseVpcFlowLogToCloudWatch bool
}
