package main

import (
	"app/root_stacks"
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

func main() {
	app := awscdk.NewApp(nil)

	root_stacks.CreateRootStack(app, "RootStack").TestRootStack()

	app.Synth(nil)
}
