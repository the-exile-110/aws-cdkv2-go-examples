package models

import "github.com/aws/aws-cdk-go/awscdk/v2"

type AppStackProps struct {
	awscdk.StackProps
}

type AppNestedStackProps struct {
	awscdk.NestedStackProps
}
