package modules

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	eks "github.com/aws/aws-cdk-go/awscdk/v2/awseks"
	"github.com/aws/jsii-runtime-go"
)

func EksCluster(stack awscdk.Stack, clusterName string) eks.Cluster {

	cluster := eks.NewCluster(stack, jsii.String(clusterName), &eks.ClusterProps{
		Version: eks.KubernetesVersion_V1_22(),
	})

	return cluster
}
